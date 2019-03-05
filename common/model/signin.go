package model

import "time"

// SignIn represents user sign in information.s
type SignIn struct {
	SignInID              int       `json:"signInID"`
	UserID                int       `json:"userID"`
	SignInTime            time.Time `json:"signInTime"`
	LastSignInTime        time.Time `json:"lastSignInTime"`
	BonusCoin             int       `json:"bonusCoin"`
	ContinuousSignInTimes int       `json:"continuousSignInTimes"`
}
