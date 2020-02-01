package main

import (
	"F1-2012-telemetry/influxsender"
	"F1-2012-telemetry/receiver"
	_ "fmt"
)

func main() {
	const IP string = "127.0.0.1"
	const PORT string = "20777"
	packetReceiver := receiver.NewPacketReceiver(IP, PORT)
	influxSender := influxsender.NewInfluxSender()
	for {
		telemData := receiver.ReceivePacket(packetReceiver)
		influxsender.SendData(influxSender, telemData)
	}
}
