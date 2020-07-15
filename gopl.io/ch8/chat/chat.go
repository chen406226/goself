// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

type Room struct {
	id 			int
	children	[]client
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	sum      = 0
	sumt     = make(client)
	room  = map[int]*Room{123:{id:123,children: make([]client,0)},124:{id:124,children: make([]client,0)}}
)

func contains(s []client, e client) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				fmt.Println(msg)
				flysnowRegexp := regexp.MustCompile(`(.{15})\:\s?to\:\[([\d]{3})\](.*)`)
				params := flysnowRegexp.FindStringSubmatch(msg)
				for _,v:= range params{
					fmt.Println(v)
				}
				roms := params[2]
				//msgc := params[3]
				//fmt.Println(roms,msgc,reflect.TypeOf(roms))
				intr, _ := strconv.Atoi(roms)
				m := room[intr]
				//fmt.Println(room,intr)
				//fmt.Println(m)
				//for _,param :=range params {
				//	fmt.Println("sdfdfs",param)
				//}
				if contains(m.children,cli) {
					cli <- msg
				}
				//1
				//if cli == sumt {
				//	cli <- msg
				//}

			}

		case cli := <-entering:
			clients[cli] = true
			if sum == 0 {
				//1
				//sumt = cli
				m := *room[123]
				//fmt.Println(m)

				room[123].children = append(m.children, cli)
				//m.children = append(m.children, cli)
				//m.children[0] = cli

				fmt.Println(len(m.children),room[123])
			} else if sum ==1 {
				m := *room[123]
				room[123].children = append(m.children, cli)
				//m.children = append(m.children, cli)
			} else if sum ==2 {
				m := *room[124]
				room[124].children = append(m.children, cli)
				//m.children = append(m.children, cli)
			} else if sum ==3 {
				m := *room[124]
				//m.children = append(m.children, cli)
				room[124].children = append(m.children, cli)
				fmt.Println(room)
			}
			sum+=1

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	//messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		fmt.Println("23333")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
