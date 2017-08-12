package main

import (
	"log"
	"net"

	"github.com/henrylee2cn/teleport/socket"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("[CLI] dial err: %v", err)
	}
	s := socket.Wrap(conn)
	defer s.Close()
	for i := 0; i < 10; i++ {
		// write request
		var packet = socket.GetPacket(nil)
		packet.Header.Id = "1"
		packet.Header.Uri = "/a/b"
		packet.Header.Codec = "json"
		packet.Header.Gzip = 5
		packet.Body = map[string]string{"a": "A"}
		err = s.WritePacket(packet)
		if err != nil {
			log.Printf("[CLI] write request err: %v", err)
			continue
		}
		log.Printf("[CLI] write request: %v", packet)
		socket.PutPacket(packet)

		// read response
		packet = socket.GetPacket(func(_ *socket.Header) interface{} {
			return new(string)
		})
		err = s.ReadPacket(packet)
		if err != nil {
			log.Printf("[CLI] read response err: %v", err)
		} else {
			log.Printf("[CLI] read response: %v", packet)
		}
		socket.PutPacket(packet)
	}
}
