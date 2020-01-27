package influxdata

// To add a new point to influxdb, I will need to convert the F1Packet struct to a hashmap

/*A channel is safe for concurrent access.
select and case is used to wait for different channels
can send data over channels

Read files into a queue in one 'thread'
Publish on udp port on another 'thread', by removing from front of queue
To be able to use queue, use "container/list" which is a doubly linked list
need to call
var list  *List
dataToPub := list.Front() // of type Element
list.Remove(dataToPub)

*/

import (
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func ConnectClient() {
	c, err := client.NewHTTPClient()
}
