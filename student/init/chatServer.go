package init

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"student/global"
	chatService "student/service/chat"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

type Room struct {	// 聊天房间
	id 			int
	children	[]client	// member 成员
}

/*
	登录用一个map  用key来判断登录没
*/

type Logged struct {
	uuid		string
	chatChan	chan string
}



var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	sum      = 0
	sumt     = make(client)
	room  = map[int]*Room{123:{id:123,children: make([]client,0)},124:{id:124,children: make([]client,0)}}
	loggedlist = make(map[string]*Logged)
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
			//flysnowRegexp := regexp.MustCompile(`(.{15})\:\s?to\:\[([\d]{3})\](.*)`)
			flysnowRegexp := regexp.MustCompile(`From\:(.{3})To\:(.{3})Room\:\[([\d]{4})\](.*)`)
			params := flysnowRegexp.FindStringSubmatch(msg)
			if len(params) <1 {
				break
			}
			//roomStr := params[3]
			//roomNum, _ := strconv.Atoi(roomStr)
			//message := params[4]
			toUser := params[2]
			if toUserLogined,ok := loggedlist[toUser];ok{
				// receive is login
				toUserLogined.chatChan <- msg
			} else {
				// receive not login save msg
				fromUser := params[1]
				chatService.SaveMessage(fromUser,toUser,msg)
			}

			//if chatRoom ,ok := room[roomNum]; ok {
			//	for _,cli := range chatRoom.children {
			//		cli <- msg
			//	}
			//} else {
			//
			//}
			//for cli := range clients {
			//	fmt.Println(msg)
			//	roms := params[2]
			//	fmt.Println(roms,msgc,reflect.TypeOf(roms))
			//	intr, _ := strconv.Atoi(roms)
			//	m := room[intr]
			//	if contains(m.children,cli) {
			//		cli <- msg
			//	}
			//}

		case cli := <-entering:
			clients[cli] = true
			//if sum == 0 {
			//	m := *room[123]
			//	room[123].children = append(m.children, cli)
			//}

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
	//ch <- "You are " + who
	//messages <- who + " has arrived"
	entering <- ch
	input := bufio.NewScanner(conn)

	var linked = false // 关联用户和管道 因为连接或得不到uuid发消息才有
	var fromUser string
	for input.Scan() {
		if linked {
			//messages <- who + ": " + input.Text()
			t := time.Now()
			tS := t.Format("2006-01-02 15:04:05")[5:]
			messages <- tS + ": " + input.Text()
		} else {
			flysnowRegexp := regexp.MustCompile(`From\:(.{36})To\:(.{36})Room\:\[([\d]{4})\](.*)`)
			params := flysnowRegexp.FindStringSubmatch(input.Text())
			if len(params) <1 {
				continue
			}
			fromUser = params[1]
			linked = true
			global.GL_LOG.Info("uuid:",fromUser,who,":连接聊天服务器")
			go func() {
				loggedlist[fromUser] = &Logged{uuid:fromUser,chatChan:ch}
				unreadMessagelist := chatService.GetMessageByToUuid(fromUser)
				if len(unreadMessagelist)>=1 {
					for _,v := range unreadMessagelist {
						ch <- v.Message
					}
				}
				defer func() {
					chatService.DeleteExpiredMessage(fromUser)
				}()
			}()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	delete(loggedlist,fromUser) // 删除已登录用户关联
	leaving <- ch


	//messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func RunChatServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:7009")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
	fmt.Println("chatServer启动成功")
}
