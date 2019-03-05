package process

import (
	"errors"
	"fmt"
	"net"
	"tenthirty/common/config"
	pb "tenthirty/common/message"
	"tenthirty/common/model"
	"tenthirty/common/util"

	"github.com/golang/protobuf/proto"
)

// Login processes login business logic.
func Login(email string, userPwd string) (err error) {

	configFile, err := config.LoadConfig("../../common/config/config.json")
	if err != nil {
		return
	}

	// Get the server address.
	serverAddr := configFile.ListenAddress

	// connect the server
	conn, err := net.Dial("tcp", serverAddr)

	if err != nil {
		model.LogErr(err, "Login dial")
		return
	}

	// TODO: Need optimization! Where to close the conn?
	defer conn.Close()

	// Create login message.
	var msg pb.Msg
	var loginMsg pb.MsgLogin

	msg.Type = pb.MsgType_Login

	loginMsg.User = &pb.User{}
	loginMsg.User.Email = email
	loginMsg.User.UserPwd = userPwd

	// Serialize login message.
	ld, err := proto.Marshal(&loginMsg)
	if err != nil {
		model.LogErr(err, "Login Marshall LoginMsg")
		return
	}

	msg.Content = ld

	// Marshall msg.
	m, err := proto.Marshal(&msg)
	if err != nil {
		model.LogErr(err, "Login Marshall Msg")
		return
	}

	// Send the msg to server.
	err = util.WritePkg(conn, m)
	if err != nil {
		model.LogErr(err, "Login WritePkg")
		return
	}

	// Receive reply from server.
	msgReply, err := util.ReadPkg(conn)
	if err != nil {
		model.LogErr(err, "Login ReadPkg")
		return
	}

	if msgReply.Type != pb.MsgType_LoginRes {
		err = fmt.Errorf("login got wrong reply msg type")
		return
	}

	var msgLoginRes pb.MsgLoginRes
	err = proto.Unmarshal(msgReply.Content, &msgLoginRes)
	if err != nil {
		model.LogErr(err, "Login Unmarshal MsgLoginRes")
		return
	}

	// TODO: implement the business logic.
	if msgLoginRes.Code == 100 {
		fmt.Println("Login success!")
		fmt.Printf("Login success! Info: %v\n", msgLoginRes)
	} else {
		fmt.Println("Login fail:", msgLoginRes.Error)
		err = errors.New(msgLoginRes.Error)
	}

	return
}
