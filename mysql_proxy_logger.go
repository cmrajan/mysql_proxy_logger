package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"flag"
	"os"
	"os/signal"
    "syscall"
)

const (
	comQuit byte = iota + 1
	comInitDB
	comQuery
	comFieldList
	comCreateDB
	comDropDB
	comRefresh
	comShutdown
	comStatistics
	comProcessInfo
	comConnect
	comProcessKill
	comDebug
	comPing
	comTime
	comDelayedInsert
	comChangeUser
	comBinlogDump
	comTableDump
	comConnectOut
	comRegiserSlave
	comStmtPrepare
	comStmtExecute
	comStmtSendLongData
	comStmtClose
	comStmtReset
	comSetOption
	comStmtFetch
)

func signalCatcher() {
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT)
        <-ch
        log.Println("CTRL-C; exiting")
        os.Exit(0)
}

	var localPort *string = flag.String("p", "3307", "localport")
    var mysqlserverIP *string = flag.String("h", "127.0.0.1", "mysqlserverIP")
    var mysqlserverPort *string = flag.String("P", "3306", "mysqlserverPort")

func main() {
	go signalCatcher()
    flag.Parse()

	fmt.Printf("Mysql Proxy Listening: localip: 127.0.0.1 localport:  %s mysqlIP: %s mysqlPort: %s \n",*localPort, *mysqlserverIP,*mysqlserverPort)
//	fmt.Printf("queries logged at: queries.log \n")
	
	ln, err := net.Listen("tcp", fmt.Sprint(":",*localPort))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go proxify(conn)
	}
}

func proxify(conn net.Conn) {
	server, err := net.Dial("tcp", fmt.Sprint(*mysqlserverIP + ":" + *mysqlserverPort))  //"localhost:3306"
	if err != nil {
		log.Println("Could not dial server")
		log.Println(err)
		conn.Close()
		return
	}
	go forward(server, conn)
	forwardWithLog(conn, server)
	server.Close()
	conn.Close()
}

func forward(src, sink net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}

		_, err = sink.Write(buffer[0:n])
		if err != nil && err != io.EOF {
			return
		}
	}
}

func forwardWithLog(src, sink net.Conn) {
//	os.OpenFile("queries.log", os.O_RDWR|os.O_APPEND, 0666)
	buffer := make([]byte, 16777219)
	for {
		n, err := src.Read(buffer)
	//	fmt.Print("dump: %s \n\n", string(buffer))
		if err != nil && err != io.EOF {
			return
		}

		if n >= 5 {
			switch buffer[4] {
			case comQuery:
			//	log.Printf("\n Query: %s \n", string(buffer[5:n]))
				fmt.Printf("\n Query: %s \n", string(buffer[5:n]))
			case comStmtPrepare:
			//	log.Printf("\n Prepare Query: %s \n", string(buffer[5:n]))
				fmt.Printf("\n Prepare Query: %s \n", string(buffer[5:n]))
			}

			switch buffer[11] {
			case comQuery:
			//	log.Printf("\n Query: %s \n", string(buffer[12:n]))
				fmt.Printf("\n Query: %s \n", string(buffer[12:n]))
			case comStmtPrepare:
			//	log.Printf("\n Prepare Query: %s \n", string(buffer[12:n]))
				fmt.Printf("\n Prepare Query: %s \n", string(buffer[12:n]))
			}
		}

		_, err = sink.Write(buffer[0:n])
		if err != nil && err != io.EOF {
			return
		}
	}

}
