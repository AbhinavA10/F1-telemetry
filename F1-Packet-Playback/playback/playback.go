// Package playback will read a series of UDP packet file, stored in .bin format,
// and publish to specified udp port
package playback

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"sync"
)

//PacketPlayer holds udp playback settings
type PacketPlayer struct {
	ip    string
	port  int
	mutex sync.Mutex
	//TODO: define udp socket
}

// NewPacketPlayer instantiates and returns a new PacketPlayer with supplied config
func NewPacketPlayer(IP string, PORT int) *PacketPlayer {
	player := PacketPlayer{ip: IP, port: PORT}
	return &player
}

func getNumFiles(path string) int {
	files, _ := ioutil.ReadDir(path)
	return len(files)
}

// PlayPackets reads bytes from a folder and publishes to udp port
func PlayPackets(player *PacketPlayer) {
	const BaseFilePath string = "/media/abhi_ubuntu18/600C7EC53DCD4841/f1/Packets"
	var n int
	var filepath string
	var numFiles int = getNumFiles(BaseFilePath)
	fmt.Println(player)
	//TODO: init UDP socket
	for n = 0; n < numFiles; n++ {
		player.mutex.Lock()
		filepath = BaseFilePath + "/Packet" + strconv.Itoa(n) + ".bin"
		udpData := ReadData(filepath)
		PublishPacket(udpData)
		player.mutex.Unlock()
	}
}

// PublishPacket publishes supplied data onto a udp packet
func PublishPacket(data []byte) {
	//TODO: do udp things
}
