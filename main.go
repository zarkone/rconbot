package main


import (
	"fmt"
	"log"
	"net"
	"bufio"
	"strings"
	"context"
)

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

	fmt.Println("sending command...")
	msg := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	cmdStr := fmt.Sprintf("rcon %s %s\n", password, command)
	msg = append(msg, []byte(cmdStr)...)
	bytes, err := conn.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	// bytes, err = fmt.Fprintf(conn, "rcon %s %s\n", password, command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bytes sent %d\n", bytes)

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
