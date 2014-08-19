package main

import "fmt"
import "encoding/hex"

func main() {

	//header

	type Header struct {
		DataLength    int
		OptDataLength int
		PacketType    int
	}

	type Package struct {
		//SyncByte 0x55
		Header    Header
		HeaderCrc []byte
		Data      []byte
		OptData   []byte
		DataCrc   []byte
	}

	var test []byte
	//test = 0xff, 0xff
	test[0] = 0xff
	test[1] = 0x01
	str := hex.EncodeToString([]byte("a"))
	fmt.Println(str)
}
