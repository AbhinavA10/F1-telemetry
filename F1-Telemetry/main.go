package main

import (
	"F1-2012-telemetry/recorder"
	_ "fmt"
)

func main() {
	//ServerSocket.setReuseAddress(true)  to listen to a port that another application
	// might already be listening to

	const IP string = "127.0.0.1"
	const PORT string = "20777"
	packetRecorder := recorder.NewPacketRecorder(IP, PORT)
	recorder.RecordPackets(packetRecorder)
}
