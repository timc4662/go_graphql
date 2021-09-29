package users

import (
	"database/sql"
	"internal/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type WrongUsernameOfPasswordError struct{}

func (err *WrongUsernameOfPasswordError) Error() string {
	return "wrong username or password"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	hashedPw, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Username, hashedPw)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("select ID from Users where Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)
	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return Id, nil
}

func (user *User) Authenticate() bool {
	stmt, err := database.Db.Prepare("select Password from users where Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(user.Username)
	var hashedPw string
	err = row.Scan(&hashedPw)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal(err)
	}
	return CheckPasswordHash(user.Password, hashedPw)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
