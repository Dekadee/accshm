package accshm

import (
	"fmt"
	"log"
	"math"
	"syscall"
	"time"
)

type LapTimeUpdate struct {
	LastTime  int32  `json:"last_time"`
	BestTime  int32  `json:"best_time"`
	Completed int32  `json:"completed"`
	Compound  string `json:"compound"`
}

type SectorTimeUpdate struct {
	SectorIndex int32 `json:"sector_index"`
	Best        int32 `json:"best"`
	LastSector  int32 `json:"last_sector"`
	Lap         int32 `json:"lap"`
}

type TrackUpdate struct {
	Flag    int32 `json:"flag"`
	Pit     int32 `json:"pit"`
	PitLane int32 `json:"pit_lane"`
}

type ACCEventPublisher struct {
	laptimeSubscriptions     map[string]func(LapTimeUpdate)
	sectortimeSubscriptions  map[string]func(SectorTimeUpdate)
	trackstatusSubscriptions map[string]func(update TrackUpdate)
	signalChan               chan interface{}
}

func NewEventPublisher() ACCEventPublisher {
	eventHandler := ACCEventPublisher{
		laptimeSubscriptions:     make(map[string]func(LapTimeUpdate), 0),
		sectortimeSubscriptions:  make(map[string]func(SectorTimeUpdate), 0),
		trackstatusSubscriptions: make(map[string]func(update TrackUpdate), 0),
		signalChan:               nil,
	}

	return eventHandler
}

func (publisher *ACCEventPublisher) Start(gTimer int) {
	if gTimer > 0 && publisher.signalChan != nil {
		publisher.signalChan = make(chan interface{})
		graphics := new(ACCGraphics)

		LTU := LapTimeUpdate{
			LastTime:  0,
			BestTime:  math.MaxInt32,
			Completed: 0,
			Compound:  "None",
		}

		STUS := make([]SectorTimeUpdate, 3)

		for i := 0; i < 3; i++ {
			STUS[i].SectorIndex = int32(i)
			STUS[i].Best = math.MaxInt32
		}

		trackUpdate := TrackUpdate{
			Flag:    0,
			Pit:     0,
			PitLane: 0,
		}

		var prevGraphicsID int32 = 0

		go func() {
			for {
				select {
				default:
					err := ReadGraphics(graphics)
					if err != nil {
						log.Println(err)
						log.Println("Failed to fetch graphics data from shared memory, sleeping 10s before retrying...")
						time.Sleep(time.Second * 10)
						continue
					}
					if prevGraphicsID != graphics.PacketId {
						prevGraphicsID = graphics.PacketId

						arrayIndex := (graphics.CurrentSectorIndex + 2) % 3
						// Handle sector time and index weirdness
						if (arrayIndex == 2 && graphics.CompletedLaps != LTU.Completed) || (arrayIndex != 2 && (graphics.CompletedLaps+1) != STUS[arrayIndex].Lap) {
							switch arrayIndex {
							case 0:
								STUS[arrayIndex].Lap = graphics.CompletedLaps + 1
								STUS[arrayIndex].LastSector = graphics.LastSectorTime
								if STUS[arrayIndex].Best > STUS[arrayIndex].LastSector && STUS[arrayIndex].LastSector != 0 {
									STUS[arrayIndex].Best = STUS[arrayIndex].LastSector
								}
								break
							case 1:
								STUS[arrayIndex].Lap = graphics.CompletedLaps + 1
								STUS[arrayIndex].LastSector = graphics.LastSectorTime - STUS[0].LastSector
								if STUS[arrayIndex].Best > STUS[arrayIndex].LastSector && STUS[arrayIndex].LastSector != 0 {
									STUS[arrayIndex].Best = STUS[arrayIndex].LastSector
								}

								break
							case 2:
								STUS[arrayIndex].Lap = graphics.CompletedLaps
								STUS[arrayIndex].LastSector = graphics.ILastTime - (STUS[0].LastSector + STUS[1].LastSector)
								if STUS[arrayIndex].Best > STUS[arrayIndex].LastSector && STUS[arrayIndex].LastSector != 0 {
									STUS[arrayIndex].Best = STUS[arrayIndex].LastSector
								}

								break
							}
							for sub := range publisher.sectortimeSubscriptions {
								publisher.sectortimeSubscriptions[sub](STUS[arrayIndex])
							}
						}

						if graphics.CompletedLaps != LTU.Completed {
							tyreCompound := syscall.UTF16ToString(graphics.CurrentTime[:])
							if LTU.Compound == "None" {
								LTU.Compound = tyreCompound
							}
							LTU.BestTime = graphics.IBestTime
							LTU.LastTime = graphics.ILastTime
							LTU.Completed = graphics.CompletedLaps
							for sub := range publisher.laptimeSubscriptions {
								publisher.laptimeSubscriptions[sub](LTU)
							}
							LTU.Compound = tyreCompound
						}

						if graphics.Flag != trackUpdate.Flag || graphics.IsInPit != trackUpdate.Pit || graphics.IsInPitLane != trackUpdate.PitLane {
							trackUpdate.Flag = graphics.Flag
							trackUpdate.Pit = graphics.IsInPit
							trackUpdate.PitLane = graphics.IsInPitLane
							for sub := range publisher.trackstatusSubscriptions {
								publisher.trackstatusSubscriptions[sub](trackUpdate)
							}
						}
					}
				case <-publisher.signalChan:
					//Exit on channel closing
					return
				}
				time.Sleep(time.Second * time.Duration(gTimer))
			}
		}()
	}
}

func (publisher *ACCEventPublisher) Stop() {
	close(publisher.signalChan)
	publisher.signalChan = nil
}

func (publisher *ACCEventPublisher) AddLaptimeSubscription(key string, handleFunc func(LapTimeUpdate)) error {
	_, ok := publisher.laptimeSubscriptions[key]
	if ok {
		return fmt.Errorf("laptime publisher already defined for %q", key)
	}
	publisher.laptimeSubscriptions[key] = handleFunc
	return nil
}

func (publisher *ACCEventPublisher) DeleteLaptimeSubscription(key string) error {
	_, ok := publisher.laptimeSubscriptions[key]
	if !ok {
		return fmt.Errorf("no laptime publisher defined for %q", key)
	} else {
		delete(publisher.laptimeSubscriptions, key)
	}
	return nil
}

func (publisher *ACCEventPublisher) AddSectortimeSubscription(key string, handleFunc func(SectorTimeUpdate)) error {
	_, ok := publisher.sectortimeSubscriptions[key]
	if ok {
		return fmt.Errorf("sector publisher already defined for %q", key)
	}
	publisher.sectortimeSubscriptions[key] = handleFunc
	return nil
}

func (publisher *ACCEventPublisher) DeleteSectortimeSubscription(key string) error {
	_, ok := publisher.sectortimeSubscriptions[key]
	if !ok {
		return fmt.Errorf("no sector publisher defined for %q", key)
	} else {
		delete(publisher.sectortimeSubscriptions, key)
	}
	return nil
}

func (publisher *ACCEventPublisher) AddTrackStatusSubscription(key string, handleFunc func(TrackUpdate)) error {
	_, ok := publisher.trackstatusSubscriptions[key]
	if ok {
		return fmt.Errorf("track status publisher already defined for %q", key)
	}
	publisher.trackstatusSubscriptions[key] = handleFunc
	return nil
}

func (publisher *ACCEventPublisher) DeleteTrackStatusSubscription(key string) error {
	_, ok := publisher.trackstatusSubscriptions[key]
	if !ok {
		return fmt.Errorf("no track status publisher defined for %q", key)
	} else {
		delete(publisher.trackstatusSubscriptions, key)
	}
	return nil
}
