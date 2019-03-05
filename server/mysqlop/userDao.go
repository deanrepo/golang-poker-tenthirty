package mysqlop

import (
	"database/sql"
	"log"
	"tenthirty/common/config"
	pb "tenthirty/common/message"
	"tenthirty/common/model"
	"time"

	_ "github.com/go-sql-driver/mysql" // Initiliazes corresponding files.
)

// Db used to manipulate mysql database.
var Db *sql.DB

func init() {
	var err error

	conf, err := config.LoadConfig("../../common/config/config.json")
	driverName := conf.RelationalDB
	dsn := conf.RelationalDSN
	Db, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Printf("open database err: %v\n", err)
		panic(err)
	}

	Db.SetMaxOpenConns(200)
	Db.SetMaxIdleConns(100)
	Db.Ping()
}

// GetUserByID gets user by user ID.
func GetUserByID(userID int64) (*pb.User, error) {
	user := &pb.User{}
	err := Db.QueryRow("SELECT user_id, nickname, email, pwd, score, coin FROM user WHERE user_id = ?",
		userID).Scan(&user.UserID, &user.Nickname, &user.Email, &user.UserPwd, &user.Score, &user.Coin)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByAcc gets user by user account.
func GetUserByAcc(email string) (*pb.User, error) {
	user := &pb.User{}
	err := Db.QueryRow("SELECT user_id, nickname, email, pwd, score, coin FROM user WHERE email = ?",
		email).Scan(&user.UserID, &user.Nickname, &user.Email, &user.UserPwd, &user.Score, &user.Coin)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetSignInByUserID gets user sign in data by user ID.
func GetSignInByUserID(userID int) (*model.SignIn, error) {
	signIn := &model.SignIn{}
	err := Db.QueryRow("SELECT sign_in_id, user_id, sign_in_time, last_sign_in_time, bonus_coin, continuous_sign_in_times FROM sign_in WHERE user_id = ?",
		userID).Scan(&signIn.SignInID, &signIn.UserID, &signIn.SignInTime, &signIn.LastSignInTime, &signIn.BonusCoin, &signIn.ContinuousSignInTimes)

	if err != nil {
		return nil, err
	}

	return signIn, nil
}

// CreateSignIn creates a sign in record.
func CreateSignIn(userID int, signTime time.Time, lastSignTime time.Time, continuousTimes int, bonusCoin int) error {
	stmt, err := Db.Prepare(`INSERT INTO sign_in (user_id, sign_in_time, last_sign_in_time, continuous_sign_in_times, bonus_coin) values (?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, signTime, lastSignTime, continuousTimes, bonusCoin)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSignIn updates sign_in table with new sign time, new last sign in time, continuous sign in times and bonus coins,
// and returns error if any.
func UpdateSignIn(signInID int, signInTime time.Time, lastSignInTime time.Time, continuousTimes int, bonusCoin int) error {
	cmd := "UPDATE sign_in SET sign_in_time = ?, last_sign_in_time = ?, continuous_sign_in_times = ?, " +
		"bonus_coin = ? WHERE sign_in_id = ?"
	stmt, err := Db.Prepare(cmd)
	if err != nil {
		model.LogErr(err, "UpdateSignIn")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(signInTime, lastSignInTime, continuousTimes, bonusCoin, signInID)
	if err != nil {
		model.LogErr(err, "UpdateSignIn")
		return err
	}

	return nil
}
