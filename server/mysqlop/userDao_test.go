package mysqlop

import (
	"log"
	pb "tenthirty/common/message"
	"testing"
	"time"
)

func TestGetUserByID(t *testing.T) {
	want := &pb.User{
		UserID:   2,
		Nickname: "test",
		UserPwd:  "123",
		Email:    "test@qq.com",
		Score:    0,
		Coin:     0,
	}

	got, err := GetUserByID(2)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("want: %v, got: %v\n", want, got)
}

func TestGetUserByAcc(t *testing.T) {
	got, err := GetUserByAcc("test@qq.com")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("got: %v\n", got)
}

func TestCreateSignIn(t *testing.T) {
	st := time.Now().Local()
	err := CreateSignIn(2, st, st, 1, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSignInByUserID(t *testing.T) {
	s, err := GetSignInByUserID(2)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("got: %v\n", s)
}

func TestUpdateSignIn(t *testing.T) {
	st := time.Date(2005, time.January, 12, 0, 0, 0, 0, time.Now().Local().Location())
	err := UpdateSignIn(1, st, st, 3, 1500)
	if err != nil {
		t.Fatal(err)
	}
}
