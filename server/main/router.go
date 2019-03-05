package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "tenthirty/common/message"
	"tenthirty/common/util"
	"tenthirty/server/process"
)

func processConn(conn net.Conn) {
	defer conn.Close()

	for {
		// Read message from client.
		msg, err := util.ReadPkg(conn)
		// Handle OpError and EOF.
		if _, ok := err.(*net.OpError); ok || err == io.EOF {
			log.Printf("client close, so the server side close, too...\n")
			return
		}

		if err != nil {
			log.Printf("processConn err: %v\n", err)
			return
		}

		// Process received message.
		switch msg.Type {

		case pb.MsgType_Login:
			// Process login.
			err := process.HandleLogin(conn, msg)
			// TODO: to implement the business logic.
			if err != nil {
				log.Printf("client login fail: %v\n", err)
				return
			}

		default:
			err = fmt.Errorf("unkonwn message type")
			log.Println(err)
			return
		}
	}
}
