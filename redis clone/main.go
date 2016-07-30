package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// the map that will hold the key value pair
var data = make(map[string]string)

func handle(conn net.Conn) {
	fmt.Println("database up and running")
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		// fmt.Println(ln) /*debug line*/
		fs := strings.Fields(ln)
		// ensuring that one of the commands is entered
		if len(fs) < 2 {
			continue
		}
		switch fs[0] {
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
		case "SET":
			if len(fs) != 3 {
				io.WriteString(conn, "We expect a value that would go with the key\n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value
		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			io.WriteString(conn, "INVALID COMMAND "+fs[0]+"\n")
		}
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
}
