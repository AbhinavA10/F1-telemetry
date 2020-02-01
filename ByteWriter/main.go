package main

import (
	"F1-Byte-writer/recorder"
)

func main() {
	const IP string = "127.0.0.1"
	const PORT string = "20777"
	packetRecorder := recorder.NewPacketRecorder(IP, PORT)
	recorder.RecordPackets(packetRecorder)
}
