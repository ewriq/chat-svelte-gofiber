package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func Login(email, rpassword string) (string, bool) {
	exmail := Find(email)
	if (exmail == true) {
		password, err := Finds(email)
		fmt.Println(password, err, rpassword)
		if password == rpassword {
			token, err := FindToken(email)
			if err != nil {
				return ".", false
			} else {
				return token, true
			}
		}
	}
	return "", false
}


func Finds(email string) (string, error) {
	query := "SELECT password FROM user WHERE email = ?"
	row := db.QueryRow(query, email)
	var password string
	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Kullanıcı bulunamadı")
		}
		return "", err
	}

	return password, nil
}

func FindToken(email string) (string, error)  {
	query := "SELECT token FROM user WHERE email = ?"
	row := db.QueryRow(query, email)
	var token string
	err := row.Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Kullanıcı bulunamadı")
		}
		return "", err
	}

	return token, nil
}


func Register(username, password, email, token string) bool {
	extinguser := Find(email) 
	if extinguser == false {
		err := addUser(username, password, email, token)
		if err == nil {
			return true
		}
		fmt.Print(err)
	} else {
	 return false
	}
	return true
}

func Find(email string) bool {
	query := "SELECT email FROM user WHERE email = ?"
	var result string
	err := db.QueryRow(query, email).Scan(&result)
	if err != nil {
		fmt.Println("\x1b[31m", err)
		return false
	}

	fmt.Print("User true")
	return true
}

func addUser(username, password, email, token string) error {
	query := "INSERT INTO user (username, password, email, token) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, username, password, email, token)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("%d Added!\n", rowCount)
	return nil
}

func Users(tkn string) ([]User, error) {
	results, err := db.Query("SELECT * FROM user WHERE token=?", tkn)
	if err != nil {
		return nil, err
	}
	defer results.Close() 

	var users []User
	for results.Next() {
		var u User
		err = results.Scan(&u.Username, &u.Password, &u.Email, &u.Token)
		if err != nil {
			return nil, err 
		}

		fmt.Println(u)
		users = append(users, u)
	}
  fmt.Println(users)
	return users, nil
}