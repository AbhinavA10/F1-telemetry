package main

import (
	"F1-2012-telemetry/influxsender"
	"F1-2012-telemetry/recorder"
	_ "fmt"
)

func main() {
	const IP string = "127.0.0.1"
	const PORT string = "20777"
	packetRecorder := recorder.NewPacketRecorder(IP, PORT)
	influxSender := influxsender.NewInfluxSender()
	influxsender.ConnectClient(influxSender)
	recorder.RecordPackets(packetRecorder)
}
