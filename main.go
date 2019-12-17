package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
)

func sendRconCommand(conn net.Conn, command string, password string) (err error) {
	udpMarker := []byte{0xFF, 0xFF, 0xFF, 0xFF}

	log.Print("sending command...")
	msg := udpMarker
	cmdStr := fmt.Sprintf("rcon %s %s\n", password, command)

	msg = append(msg, []byte(cmdStr)...)
	bytes, err := conn.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("bytes sent %d\n", bytes)

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	password := "freeforall"
	command := "status"

	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", ":27960")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("connection created")

	err = sendRconCommand(conn, command, password)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(conn)

	b := strings.Builder{}

	for s.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			b.WriteString(s.Text())
			fmt.Println(s.Text())
		}
	}
	fmt.Println(b.String())
}
