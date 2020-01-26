package main

import (
	"F1-PacketPlayback/playback"
)

func main() {
	const IP string = "127.0.0.1"
	const PORT int = 20777
	packetPlayer := playback.NewPacketPlayer(IP, PORT)
	playback.PlayPackets(packetPlayer)
}
