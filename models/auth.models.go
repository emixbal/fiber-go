package models

import (
	"database/sql"
	"fiber-go/db"
	"fiber-go/helpers"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Password string `json:"password"`
}

func CheckLogin(username, passwordTxt string) (bool, error) {
	var user User
	var pwdHashDb string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username=?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&user.Id, &user.Name, &pwdHashDb,
	)
	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("query error")
		fmt.Println(err)
		return false, err
	}

	match, err := helpers.CheckPasswordHash(pwdHashDb, passwordTxt)
	if !match {
		fmt.Println("password salah")
		return false, err
	}

	return true, nil
}
