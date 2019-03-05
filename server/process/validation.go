package process

import (
	"database/sql"
	pb "tenthirty/common/message"
	"tenthirty/common/model"
	"tenthirty/server/mysqlop"
	"time"
)

// ValidateLogin validates user's login information from the database.
func ValidateLogin(userAcc, userPwd string) (*pb.User, error) {

	user, err := mysqlop.GetUserByAcc(userAcc)
	if err == nil {
		// Validate user password.
		if userPwd == user.UserPwd {
			return user, err
		}

		// If password doesn't match then return ErrWrongPwd error.
		err = model.ErrWrongPwd
		return nil, err
	}

	// If error is not nil then return nil and the error.
	if err == sql.ErrNoRows {
		err = model.ErrUserNotExists
	}
	return nil, err
}

// ValidateSignInInfo validates user sign in info and  return
// continuous sign in times and bonus coins, if any error occurs return -1, -1, error.
func ValidateSignInInfo(userID int, signInTime time.Time) (int, int, error) {
	signIn, err := mysqlop.GetSignInByUserID(userID)
	if err != nil {
		if err == sql.ErrNoRows {

			// If there were no records of user then create one.
			err = mysqlop.CreateSignIn(userID, signInTime, signInTime, 1, 500)
			if err != nil {
				return -1, -1, err
			}

			return 1, 500, nil
		}

		return -1, -1, err
	}

	lastSignInTime := signIn.SignInTime
	// If sign in in the same day continuously then do nothing.
	if signInTime == lastSignInTime {
		return signIn.ContinuousSignInTimes, 0, nil
	}

	// If sign in continuously then update the sign in info of the user.
	if signInTime == lastSignInTime.Add(24*time.Hour) {
		var bonusCoin int
		// Continuous times add one.
		continuousTimes := signIn.ContinuousSignInTimes + 1
		if continuousTimes <= 7 {
			bonusCoin = continuousTimes * 500
			// Update sign_in table with new sign time, last sign in time, continuous sign in times and bonus coins.
			err = mysqlop.UpdateSignIn(signIn.SignInID, signInTime, lastSignInTime, continuousTimes, bonusCoin)

			if err != nil {
				return -1, -1, err
			}
		} else {
			// If continuous times exceeds 7 times then reset the sign in times as the one.
			continuousTimes = 1
			bonusCoin = continuousTimes * 500

			// Update sign_in table with new sign in time, last sign in time, continuous sign in times and bonus coins.
			err = mysqlop.UpdateSignIn(signIn.SignInID, signInTime, lastSignInTime, continuousTimes, bonusCoin)

			if err != nil {
				return -1, -1, err
			}
		}

		return continuousTimes, bonusCoin, nil
	}

	// If not signs in continuously then updates the sign in info of the user as first time signing in.
	err = mysqlop.UpdateSignIn(signIn.SignInID, signInTime, lastSignInTime, 1, 500)

	if err != nil {
		return -1, -1, err
	}

	return 1, 500, nil
}
