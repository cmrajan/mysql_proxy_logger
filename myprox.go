package main

import (
	"io"
	"log"
	"net"
	"fmt"
//	"os"
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

func main() {
	fmt.Printf("Mysql Proxy Listening: 3307 \n")
//	fmt.Printf("queries logged at: queries.log \n")
	
	ln, err := net.Listen("tcp", ":3307")
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
	server, err := net.Dial("tcp", "localhost:3306")
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
