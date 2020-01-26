// Package reader Reads UDP packet files, stored in .bin format
package reader

import (
	"fmt"
	"os"
)

// ReadData reads files from supplied path and returns the binary data
func ReadData() []byte {
	file, err := os.Open("/media/abhi_ubuntu18/600C7EC53DCD4841/f1/Packets/Packet10.bin")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close() // once ReadData() is done, close the file

	// allocate a buffer the size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	bytesRead, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("number of bytes read: ", bytesRead)
	fmt.Println("bytestream:", buffer)
	return buffer
}
