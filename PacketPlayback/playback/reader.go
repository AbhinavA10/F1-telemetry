package playback

import (
	"log"
	"os"
)

// ReadData reads bytes from a single supplied path and returns the binary data
func ReadData(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close() // once ReadData() is done, close the file

	// allocate a buffer the size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return nil
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	numBytesRead, err := file.Read(buffer)
	if err != nil {
		log.Println(err)
		return nil
	}
	_ = numBytesRead
	return buffer
}
