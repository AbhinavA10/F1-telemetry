package main

import (
	"F1-2012-telemetry/influxsender"
	"F1-2012-telemetry/receiver"
	_ "fmt"
)

func main() {
	const IP string = "127.0.0.1"
	const PORT string = "20777"
	packetRecorder := receiver.NewPacketRecorder(IP, PORT)
	influxSender := influxsender.NewInfluxSender()
	influxsender.ConnectClient(influxSender)
	receiver.RecordPackets(packetRecorder)
}
