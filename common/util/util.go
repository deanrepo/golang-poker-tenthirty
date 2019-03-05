package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	pb "tenthirty/common/message"

	"github.com/golang/protobuf/proto"
)

// WritePkg writes msg package to the remote peer of the connection.
func WritePkg(conn net.Conn, data []byte) (err error) {
	// Put data length and data into buf
	var msgLen uint32
	msgLen = uint32(len(data))
	bufLen := msgLen + 4
	buf := make([]byte, bufLen)
	binary.BigEndian.PutUint32(buf[0:4], msgLen)

	n := copy(buf[4:], data)
	if n != int(msgLen) {
		err = errors.New("WritePkg do not copy all data to buf")
		return
	}

	// Send the pkg.
	n, err = conn.Write(buf)
	if n != int(bufLen) || err != nil {
		err = fmt.Errorf("Writepkg write err: %v", err)
		return
	}

	return
}

// ReadPkg reads msg package from remote peer of the connection.
func ReadPkg(conn net.Conn) (msg *pb.Msg, err error) {

	// Allocate a buf to store the msg length.
	buf := make([]byte, 4)

	_, err = conn.Read(buf[:4])
	if err != nil {
		err = fmt.Errorf("ReadPkg read Msg len err: %v", err)
		return
	}

	// Get the length of the message.
	msgLen := binary.BigEndian.Uint32(buf[:4])

	buf = make([]byte, msgLen)

	// Get the message
	n, err := conn.Read(buf[:msgLen])
	if err != nil || n != int(msgLen) {
		err = fmt.Errorf("ReadPkg read Msg err: %v", err)
		return
	}

	// Initialize msg
	msg = &pb.Msg{}
	// Unmarshall the message
	err = proto.Unmarshal(buf[:], msg)
	if err != nil {
		err = fmt.Errorf("ReadPkg unmarshal Msg err: %v", err)
		return nil, err
	}

	return
}
