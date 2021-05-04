package accshm

import (
	"bytes"
	"encoding/binary"
	"github.com/hidez8891/shm"
	"unsafe"
)

type ACCPhysics struct {
	PacketId             int32         
	Gas                  float32       
	Brake                float32       
	Fuel                 float32       
	Gear                 int32         
	RPM                  int32         
	SteerAngle           float32       
	SpeedKmh             float32       
	Velocity             [3]float32    
	AccG                 [3]float32    
	WheelSlip            [4]float32    
	WheelLoad            [4]float32    
	WheelPressure        [4]float32    
	WheelAngularSpeed    [4]float32    
	TyreWear             [4]float32    
	TyreDirtyLevel       [4]float32    
	TyreCoreTem          [4]float32    
	CamberRAD            [4]float32    
	SuspensionTravel     [4]float32    
	DRS                  float32       
	TC                   float32       
	Heading              float32       
	Pitch                float32       
	Roll                 float32       
	CgHeight             float32       
	CarDamage            [5]float32    
	NumberOfTyresOut     int32         
	PitLimiterOn         int32         
	ABS                  float32       
	KERSCharge           float32       
	KERSInput            float32       
	AutoShifterOn        int32         
	RideHeight           [2]float32    
	TurboBoost           float32       
	Ballast              float32       
	AirDensity           float32       
	AirTemp              float32       
	RoadTemp             float32       
	LocalAngularVelocity [3]float32    
	FinalFF              float32       
	PerformanceMeter     float32       
	EngineBrake          int32         
	ERSRecoup            int32         
	ERSPower             int32         
	ERSHeatCharging      int32         
	ERSIsCharging        int32         
	KERSCurrentKJ        float32       
	DRSAvailable         int32         
	DRSEnabled           int32         
	BrakeTemp            [4]float32    
	Clutch               float32       
	TyreTempI            [4]float32    
	TyreTempM            [4]float32    
	TyreTempO            [4]float32    
	IsAIControlled       int32         
	TyreContactPoint     [4][3]float32 
	TyreContactNormal    [4][3]float32 
	TyreContactHeading   [4][3]float32 
	BrakeBias            float32       
	LocalVelocity        [3]float32    
	P2PActivations       float32       
	P2PStatus            float32       
	CurrentMaxRPM        int32         
	MZ                   [4]float32    
	FX                   [4]float32    
	FY                   [4]float32    
	SlipRatio            [4]float32    
	SlipAngle            [4]float32    
	TCInAction           int32         
	ABSInAction          int32         
	SuspensionDamage     [4]float32    
	TyreTemp             [4]float32    
}

type ACCGraphics struct {
	PacketId                 int32
	Status                   int32
	SessionType              int32
	CurrentTime              [15]uint16
	LastTime                 [15]uint16
	BestTime                 [15]uint16
	Split                    [15]uint16
	CompletedLaps            int32
	Position                 int32
	ICurrentTime             int32
	ILastTime                int32
	IBestTime                int32
	SessionTimeLeft          float32
	DistanceTraveled         float32
	IsInPit                  int32
	CurrentSectorIndex       int32
	LastSectorTime           int32
	NumberOfLaps             int32
	TyreCompound             [33]uint16
	ReplayTimeMultiplier     float32
	NormalizedCarPosition    float32
	ActiveCars               int32
	CarCoordinates           [60][3]float32
	CarId                    [60]int32
	PlayerCarId              int32
	PenaltyTime              float32
	Flag                     int32
	PenaltyShortCut          int32
	IdealLineOn              int32
	IsInPitLane              int32
	SurfaceGrip              float32
	MandatoryPitDone         int32
	WindSpeed                float32
	WindDirection            float32
	IsSetupMenuVisible       int32
	MainDisplayIndex         int32
	SecondaryDisplayIndex    int32
	TC                       int32
	TCCut                    int32
	EngineMap                int32
	ABS                      int32
	FuelXLap                 int32
	RainLights               int32
	FlashingLights           int32
	LightStage               int32
	ExhaustTemperature       float32
	WiperLevel               int32
	DriverStintTotalTimeLeft int32
	DriverStintTimeLeft      int32
	RainTyres                int32
}

type ACCStatic struct {
	SMVersion                [15]uint16
	ACVersion                [15]uint16
	NumberOfSessions         int32
	NumCars                  int32
	CarModel                 [33]uint16
	Track                    [33]uint16
	PlayerName               [33]uint16
	PlayerSurName            [33]uint16
	PlayerNickname           [33]uint16
	SectorCount              int32
	MaxTorque                float32
	MaxPower                 float32
	MaxRPM                   int32
	MaxFuel                  float32
	MaxSuspensionTravel      [4]float32
	TyreRadius               float32
	MaxTurboBoost            float32
	Deprecated1              float32
	Deprecated2              float32
	PenaltiesEnabled         int32
	AidFuelRate              int32
	AidTireRate              int32
	AidMechanicalDamage      float32
	AidAllowTyreBlankets     int32
	AidStability             float32
	AidAutoClutch            int32
	AidAutoBlip              int32
	HasDRS                   int32
	HasERS                   int32
	HasKERS                  int32
	KERSMaxJ                 float32
	EngineBrakeSettingsCount int32
	ERSPowerControllerCount  int32
	TrackSplineLength        float32
	TrackConfiguration       [33]uint16
	ERSMaxJ                  float32
	IsTimedRace              int32
	HasExtraLap              int32
	CarSkin                  [33]uint16
	ReversedGridPosition     int32
	PitWindowStart           int32
	PitWindowEnd             int32
	IsOnline                 int32
}

func ReadPhysics(physics *ACCPhysics) error {
	physicsSize := (int32)(unsafe.Sizeof(*physics))

	rbuf := make([]byte, physicsSize)
	buf := &bytes.Buffer{}

	r, err := shm.Open("Local\\acpmf_physics", physicsSize)
	if err != nil {
		return err
	}
	_, err = r.Read(rbuf)
	if err != nil {
		return err
	}
	buf.Write(rbuf)
	err = binary.Read(buf, binary.LittleEndian, physics)
	if err != nil {
		return err
	}
	err = r.Close()
	if err != nil {
		return err
	}
	return nil
}

func ReadGraphics(graphics *ACCGraphics) error {
	physicsSize := (int32)(unsafe.Sizeof(*graphics))

	rbuf := make([]byte, physicsSize)
	buf := &bytes.Buffer{}

	r, err := shm.Open("Local\\acpmf_graphics", physicsSize)
	if err != nil {
		return err
	}
	_, err = r.Read(rbuf)
	if err != nil {
		return err
	}
	buf.Write(rbuf)
	err = binary.Read(buf, binary.LittleEndian, graphics)
	if err != nil {
		return err
	}
	err = r.Close()
	if err != nil {
		return err
	}
	return nil
}



func ReadStatic(static *ACCStatic) error {
	staticSize := (int32)(unsafe.Sizeof(*static))

	rbuf := make([]byte, staticSize)
	buf := &bytes.Buffer{}

	r, err := shm.Open("Local\\acpmf_static", staticSize)
	if err != nil {
		return err
	}
	_, err = r.Read(rbuf)
	if err != nil {
		return err
	}
	buf.Write(rbuf)
	err = binary.Read(buf, binary.LittleEndian, static)
	if err != nil {
		return err
	}
	err = r.Close()
	if err != nil {
		return err
	}
	return nil
}
