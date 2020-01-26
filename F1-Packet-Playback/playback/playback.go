// Package playback will read a series of UDP packet file, stored in .bin format,
// and publish to specified udp port
package playback

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
)

var ip string
var port int // lowercase, to avoid exporting
var configMutex sync.Mutex

//TODO: define UDP sockets

func init() {
	ip = "127.0.0.1" // default ip
	port = 20777     // default
}

// SetIP changes ip and port to supplied values
func SetIP(IP string, PORT int) {
	configMutex.Lock()
	defer configMutex.Unlock()
	ip = IP
	port = PORT
}

func getNumFiles(path string) int {
	files, _ := ioutil.ReadDir(path)
	return len(files)
}

// PlayPackets reads bytes from a folder and publishes to udp port
func PlayPackets() {
	const BaseFilePath string = "/media/abhi_ubuntu18/600C7EC53DCD4841/f1/Packets"
	var n int
	var filepath string
	var numFiles int = getNumFiles(BaseFilePath)

	//TODO: init UDP socket
	for n = 0; n < numFiles; n++ {
		filepath = BaseFilePath + "/Packet" + strconv.Itoa(n) + ".bin"
		udpData := ReadData(filepath)
		PublishPacket(udpData)
		fmt.Println(n)
	}
}

// PublishPacket publishes supplied data onto a udp packet
func PublishPacket(data []byte) {

}
