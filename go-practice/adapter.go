// You, are developing a media player application in Golang that supports
// various types of media files, including MP3, WAV, and OGG.
// Each media file type has its own decoding library with different interfaces.
// Your goal is to implement the Adapter Pattern to create a unified
// interface for decoding and playing these media files, allowing seamless
// integration of different decoding libraries into your media player application.

/*************** Adapter Pattern Practice *****************/
package main

import (
	"fmt"
)

type Decoder interface {
	decode(data []byte)
}

// different type of media files
type MP3 struct{}
type WAV struct{}
type OGG struct{}

// implementation of the interface
func (mp3 *MP3) decode(data []byte) {
	//decode the MP3 file with the MP3 decodoing library
	fmt.Println("Decoded MP3 file")
}
func (WAV *WAV) decode(data []byte) {
	//decode the WAV file with the WAV decodoing library
	fmt.Println("Decoded WAV file")
}
func (ogg *OGG) decode(data []byte) {
	//decode the OGG file with the OGG decodoing library
	fmt.Println("Decoded OGG file")
}

// adapter
type mediaAdpter struct {
	decoder Decoder
}

// adaptee
func (ad *mediaAdpter) decodeFile(decoder Decoder, data []byte) {
	ad.decoder = decoder
	ad.decoder.decode(data)

}

func main() {

	//get the file form the client and then convert it into byte array
	data := []byte("byte data of the file")
	//then the following code executes

	//initialize the proper file type struct
	//in this case file1 is MP3, file2 is WAV, file3 is OGG for illustration purpose
	file1 := MP3{}
	file2 := WAV{}
	file3 := OGG{}

	//initialize the mediaAdapter struct
	adapter := mediaAdpter{}
	adapter.decodeFile(&file1, data)
	adapter.decodeFile(&file2, data)
	adapter.decodeFile(&file3, data)

}
