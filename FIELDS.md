# Memory Page Fields

This is basically a writeup of the original source.
The order of these field is as they are in memory, which is defined by the game.

`[15]uint16` type is a 15-character long UTF-16 string, which can be read via:
```go
// Pass a copy of the slice to the syscall
syscall.UTF16ToString(graphics.CurrentTime[:])
```

## Physics

Memory page containing most of the cars state. Will always report data for the focussed cars (may not be the player car inr replays).

| Name | Type | Description|
--- | --- | ---
| PacketID | `int32` | Incrementing number identifying each time update step (0 if not updating) |
| Gas | `float32` | Gas Pedal State `[0;1]` |
| Brake | `float32` | Brake Pedal State `[0;1]` |
| Fuel | `float32` | Fuel left in tank in Liters |
| Gear | `int32` | Selected gear (0 => Reverse, 1 => Neutral, 2 => First Gear, ...) |
| RPM | `int32` | Engine RPM |
| SteerAngle | `float32` |  |
| SpeedKmh | `float32` | Speed in KPH |
| Velocity | `[3]float32` | Velocity for each cartesian axis |
| AccG | `[3]float32` | Acceleration for each cartesian axis in G |
| WheelSlip | `[4]float32` |  |
| WheelLoad | `[4]float32` |  |
| WheelPressure | `[4]float32` |  |
| WheelAngularSpeed | `[4]float32` |  |
| TyreWear | `[4]float32` | |
| TyreDirtyLevel | `[4]float32` | |
| CamberRAD | `[4]float32` | |
| SuspensionTravel | `[4]float32` | |
| SuspensionTravel | `[4]float32` | |
| DRS | `float32` | |
| TC | `float32` | |
| Heading | `float32` | |
| Pitch | `float32` | |
| Roll | `float32` | |
| CgHeight | `float32` | |
| CarDamage | `[5]float32` | |
| NumberOfTyresOut | `int32` | |
| PitLimiterOn | `int32` | |
| ABS | `float32` | |
| KERSCharge | `float32` | |
| KERSInput | `float32` | |
| AutoShifterOn | `int32` | |
| RideHeight | `[2]float32` | Front and rear ride height |
| TurboBoost | `float32` | |
| Ballast | `float32` | |
| AirDensity | `float32` | |
| AirTemp | `float32` | |
| RoadTemp | `float32` | |
| LocalAngularVelocity | `[3]float32` | |
| FinalFF | `float32` | Something to do with Force Feedback |
| PerformanceMeter | `float32` | |
| EngineBreak | `int32` | |
| ERSRecoup | `int32` | |
| ERSPower | `int32` | |
| ERSHeatCharging | `int32` | |
| ERSIsCharging | `int32` | |
| ERSCurrentKJ | `float32` | |
| DRSAvailable | `int32` | |
| DRSEnabled | `int32` | |
| BrakeTemp | `[4]float32` | |
| Clutch | `float32` | |
| TyreTempI | `[4]float32` | |
| TyreTempM | `[4]float32` | |
| TyreTempO | `[4]float32` | |
| IsAIControlled | `int32` | |
| TyreContactPoint | `[4][3]float32` | Vector for each wheel |
| TyreContactNormal | `[4][3]float32` | Vector for each wheel |
| TyreContactHeading | `[4][3]float32` | Vector for each wheel |
| BrakeBias | `float32` | |
| LocalVelocity | `[3]float32` | |
| P2PActivations | `float32` | |
| P2PStatus | `float32` | |
| P2PStatus | `int3232` | |
| MZ | `[4]float32` | |
| FX | `[4]float32` | |
| FY | `[4]float32` | |
| SlipRatio | `[4]float32` | |
| SlipAngle | `[4]float32` | |
| TCInAction | `int32` | |
| ABSInAction | `int32` | |
| SuspensionDamage | `[4]float32` | |
| TyreTemp | `[4]float32` | |

## Graphics

Memory pages containing information about the session, other cars in the session, laptimes and some more info on the car state.
Will always report data for the focussed cars (may not be the player car inr replays).

| Name | Type | Description|
--- | --- | ---
| PacketID | `int32` | Incrementing number identifying each time update step (0 if not updating) |
| Status | `int32` | |
| CurrentTime | `[15]uint16` | |
| LastTime | `[15]uint16` | |
| BestTime | `[15]uint16` | |
| Split | `[15]uint16` | |
| CompletedLaps | `int32` | |
| Position | `int32` | |
| ICurrentTime | `int32` | |
| IBestTime | `int32` | |
| SessionTimeLeft | `float32` | |
| DistanceTraveled | `float32` | |
| IsInPit | `int32` | |
| CurrentSectorIndex | `int32` | |
| LastSectorTime | `int32` | |
| NumberOfLaps | `int32` | |
| TyreCompound | `[33]uint16` | |
| ReplayTimeMultiplier | `float32` | |
| NormalizedCarPosition | `float32` | |
| ActiveCars | `int32` | |
| CarCoordinates | `[60][3]float32` | |
| CarId | `[60]int32` | |
| PlayerCarId | `int32` | |
| PenaltyTime | `float32` | |
| Flag | `int32` | |
| PenaltyShortCut | `int32` | |
| IdealLineOn | `int32` | |
| IsInPitLane | `int32` | |
| SurfaceGrip | `float32` | |
| MandatoryPitDone | `int32` | |
| WindSpeed | `float32` | |
| WindDirection | `float32` | |
| IsSetupMenuVisible | `int32` | |
| MainDisplayIndex | `int32` | |
| SecondaryDisplayIndex | `int32` | |
| TC | `int32` | |
| TCCut | `int32` | |
| EngineMap | `int32` | |
| ABS | `int32` | |
| FuelXLap | `int32` | |
| RainLights | `int32` | |
| FlashingLights | `int32` | |
| LightStage | `int32` | |
| ExhaustTemperature | `float32` | |
| WiperLevel | `int32` | |
| DriverStintTotalTimeLeft | `int32` | |
| DriverStintTimeLeft | `int32` | |
| RainTyres | `int32` | |

## Static

Memory page about static information about the car, track, session and settings.

| Name | Type | Description|
--- | --- | ---
| SMVersion | `[15]uint16` | Version of the Shared Memory page |
| ACVersion | `[15]uint16` | Version of the game |
| NumberOfSessions | `int32` | |
| NumCars | `int32` | |
| CarModel | `[33]uint16` |  |
| Track | `[33]uint16` |  |
| PlayerName | `[33]uint16` |  |
| PlayerSurName | `[33]uint16` |  |
| PlayerNickname | `[33]uint16` |  |
| SectorCount | `int32` | |
| MaxTorque | `float32` | |
| MaxPower | `float32` | |
| MaxRPM | `int32` | |
| MaxFuel | `float32` | |
| MaxSuspensionTravel | `[4]float32` | |
| TyreRadius | `float32` | |
| MaxTurboBoost | `float32` | |
| Deprecated1 | `float32` | Invalid fields, still need to be read |
| Deprecated2 | `float32` | Invalid fields, still need to be read |
| PenaltiesEnabled | `int32` | |
| AidFuelRate | `int32` | |
| AidTireRate | `int32` | |
| AidMechanicalDamage | `float32` | |
| AidAllowTyreBlankets | `int32` | |
| AidStability | `float32` | |
| AidAutoClutch | `int32` | |
| AidAutoBlip | `int32` | |
| HasDRS | `int32` | |
| HasERS | `int32` | |
| HasKERS | `int32` | |
| KERSMaxJ | `float32` | |
| EngineBrakeSettingsCount | `int32` | |
| ERSPowerControllerCount | `int32` | |
| TrackSplineLength | `float32` | |
| TrackConfiguration | `[33]uint16` |  |
| ERSMaxJ | `float32` | |
| IsTimedRace | `int32` | |
| HasExtraLap | `int32` | |
| CarSkin | `[33]uint16` |  |
| ReversedGridPosition | `int32` | |
| PitWindowStart | `int32` | |
| PitWindowEnd | `int32` | |
| IsOnline | `int32` | |
