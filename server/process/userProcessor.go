package process

import (
	"net"
	pb "tenthirty/common/message"
	"tenthirty/common/model"
	"tenthirty/common/util"
	"time"

	"github.com/golang/protobuf/proto"
)

// HandleLogin process player login business logic.
func HandleLogin(conn net.Conn, msg *pb.Msg) (err error) {
	// Unmarshal the login msg.
	var msgLogin pb.MsgLogin
	err = proto.Unmarshal(msg.Content, &msgLogin)
	if err != nil {
		model.LogErr(err, "HandleLogin unmarshall MsgLogin")
		return
	}

	var msgLogRes pb.MsgLoginRes
	var msgReply pb.Msg
	msgReply.Type = pb.MsgType_LoginRes

	user, err := ValidateLogin(msgLogin.User.Email, msgLogin.User.UserPwd)
	if err == nil {
		msgLogRes.Code = 100
		msgLogRes.User = user

		// Handle continuous sign in business logic.
		now := time.Now().Local()
		thisSignInTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		times, bonus, err := ValidateSignInInfo(int(user.UserID), thisSignInTime) // Return continuous sign in times and bonus coins.
		if err != nil {
			model.LogErr(err, "HandleLogin ValidateSignInInfo")
			return err
		}

		msgLogRes.SignInfo = &pb.MsgSignIn{}
		msgLogRes.SignInfo.ContinuousSignInTimes = int32(times)
		msgLogRes.SignInfo.BonusCoin = int32(bonus)
	} else {
		msgLogRes.Code = 200
		msgLogRes.Error = err.Error()
	}

	// Marshal MsgLoginRes.
	d, err := proto.Marshal(&msgLogRes)
	if err != nil {
		model.LogErr(err, "HandleLogin Marshal MsgLoginRes")
		return
	}

	msgReply.Content = d

	// Marshal reply Msg
	m, err := proto.Marshal(&msgReply)
	if err != nil {
		model.LogErr(err, "HandleLogin Marshal reply Msg")
		return
	}

	err = util.WritePkg(conn, m)
	if err != nil {
		model.LogErr(err, "HandleLogin WritePkg")
		return
	}

	return
}
