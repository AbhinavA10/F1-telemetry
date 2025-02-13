// Package playback will read a series of UDP packet file, stored in .bin format,
// and publish to specified udp port
package playback

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"sync"
	"time"
)

//PacketPlayer holds udp playback settings
type PacketPlayer struct {
	connection *net.UDPConn
	mutex      sync.Mutex
}

// NewPacketPlayer instantiates and returns a new PacketPlayer with supplied config
func NewPacketPlayer(IP string, PORT string) *PacketPlayer {
	player := PacketPlayer{}
	udpAddr, err := net.ResolveUDPAddr("udp", IP+":"+PORT)
	if err != nil {
		fmt.Println(err)
	}
	// Make a udp client
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	player.connection = conn
	// Source of UDP publishing will be conn.LocalAddr().String()
	// We are publishing to conn.RemoteAddr().String()
	return &player
}

//getNumFiles returns the number of files in the supplied directory path
func getNumFiles(path string) int {
	files, _ := ioutil.ReadDir(path)
	return len(files)
}

// PlayPackets reads bytes from a folder and publishes to udp port
func PlayPackets(player *PacketPlayer) {
	defer player.connection.Close()
	const BaseFilePath string = "/media/abhi_ubuntu18/600C7EC53DCD4841/f1/Packets"
	var n int
	var filepath string
	var numFiles int = getNumFiles(BaseFilePath)
	for n = 0; n < numFiles; n++ {
		player.mutex.Lock()
		filepath = BaseFilePath + "/Packet" + strconv.Itoa(n) + ".bin"
		udpData := ReadData(filepath)
		PublishPacket(player, udpData)
		//17.98451965ms because 13178 packets were recorded in 3:57 minutes
		time.Sleep(17 * time.Millisecond)
		player.mutex.Unlock()
	}
}

// PublishPacket publishes supplied data onto a udp packet
func PublishPacket(player *PacketPlayer, data []byte) {
	_, err := player.connection.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
