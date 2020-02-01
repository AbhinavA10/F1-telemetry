package f1packet

import (
	"bytes"
	"encoding/binary" // encoding and decoding struct value to bytes
	"fmt"
	"reflect"
)

//F1Packet is the udp packet structure sent by F1 2012
type F1Packet struct {
	// All fields must be exported, for binary.Read to be successful
	Time                         float32
	LapTime                      float32
	LapDistance                  float32
	Distance                     float32
	X                            float32
	Y                            float32
	Z                            float32
	Speed                        float32 // read in as m/s, but converted to mph before sending to influx
	WorldSpeedX                  float32
	WorldSpeedY                  float32
	WorldSpeedZ                  float32
	XR                           float32
	Roll                         float32
	ZR                           float32
	XD                           float32
	Pitch                        float32
	ZD                           float32
	SuspensionPositionRearLeft   float32
	SuspensionPositionRearRight  float32
	SuspensionPositionFrontLeft  float32
	SuspensionPositionFrontRight float32
	SuspensionVelocityRearLeft   float32
	SuspensionVelocityRearRight  float32
	SuspensionVelocityFrontLeft  float32
	SuspensionVelocityFrontRight float32
	WheelSpeedBackLeft           float32
	WheelSpeedBackRight          float32
	WheelSpeedFrontLeft          float32
	WheelSpeedFrontRight         float32
	Throttle                     float32
	Steer                        float32
	Brake                        float32
	Clutch                       float32
	Gear                         float32
	LateralAcceleration          float32
	LongitudinalAcceleration     float32
	Lap                          float32
	EngineRevs                   float32
	// 152 bytes ends here.
	/* New Fields in Patch 12. Didn't have. */
	/*
		NewField1                  float32 // Always 1
		RacePosition               float32 // Position in race
		KersRemaining              float32 // Kers Remaining
		KersRecharge               float32 // Always 400000?
		DrsStatus                  float32 // Drs Status
		Difficulty                 float32 // 2 = Medium or Easy, 1 = Hard, 0 = Exper
		Assists                    float32 // 0 = All assists are off.  1 = some assist is on
		FuelRemaining              float32 // Not sure if laps or Litres
		SessionType                float32 // 9.5 = race, 10 = time trail / time attack, 170 = quali, practice, championsmod
		NewField10                 float32
		Sector                     float32 // Sector (0, 1, 2)
		TimeSector1                float32 // Time Intermediate
		TimeSector2                float32 // Time Intermediate
		BrakeTemperatureRearLeft   float32
		BrakeTemperatureRearRight  float32
		BrakeTemperatureFrontLeft  float32
		BrakeTemperatureFrontRight float32
		NewField18                 float32 // Always 0
		NewField19                 float32 // Always 0
		NewField20                 float32 // Always 0
		NewField21                 float32 // Always 0
		CompletedLapsInRace        float32 // Number of laps Completed (in GP only )
		TotalLapsInRace            float32 // Number of laps in GP (GP only)
		TrackLength                float32 // Track Length
		PreviousLapTime            float32 // Lap time of previous lap
		NewField26                 float32 // New, for F1 2013. Always 0
		NewField27                 float32 // New, for F1 2013. Always 0
		NewField28                 float32 // New, for F1 2013. Always 0
	*/
}

// DatagramToStruct returns a new F1Packet from bytes
func DatagramToStruct(buf []byte) *F1Packet {
	var pack F1Packet
	r := bytes.NewReader(buf) // binary.Read requires a reader
	// F1 outputs packet data in LittleEndian style
	// Bytes read from r are written using LittleEndian decoding into successive fields of the pack pointer
	// https://golang.org/pkg/encoding/binary/#Read
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		fmt.Println(err)
	}
	pack.Speed = pack.Speed * 2.23694 // convert for InfluxDB
	return &pack
}

//StructToMap converts F1Packet struct to map of string->float32
func StructToMap(packet *F1Packet) map[string]interface{} {
	m := make(map[string]interface{})
	values := reflect.ValueOf(packet).Elem() // get values of struct fields
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		val := float32(values.Field(i).Float()) // all of struct is float32
		name := types.Field(i).Name
		m[name] = val
	}
	return m
}
