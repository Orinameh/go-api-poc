package models

import (
	"errors"
	"fmt"
	"go-api/utils"
)

// Error
var (
	ErrUserNotFound = errors.New("User not found")
)

// User struct
type User struct {
	UID       uint32 `json:"_id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int8   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewUser function
func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()
	tx, err := con.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	sql := "insert into users (nickname, email, password) values ($1, $2, $3) returning uid"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return false, err
		}

		defer stmt.Close()
		hashedPassword, err := utils.Bcrypt(user.Password)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return false, err
		}
		err = stmt.QueryRow(user.Nickname, user.Email, hashedPassword).Scan(&user.UID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	sql = "insert into wallets (public_key, usr) values ($1, $2)"

	wallet := Wallet{User: user}
	wallet.GeneratePublicKey()
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return false, err
		}
		_, err = stmt.Exec(wallet.PublicKey, wallet.User.UID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	return true, tx.Commit()
}

// GetUsers function
func GetUsers() ([]User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var users []User
	for rs.Next() {
		var user User
		err := rs.Scan(&user.UID, &user.Nickname, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUser function
func GetUser(id uint32) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where uid = $1"
	row, err := con.Query(sql, id)
	if err != nil {
		return User{}, err
	}
	defer row.Close()
	var user User
	for row.Next() {
		err := row.Scan(&user.UID, &user.Nickname, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}
	}
	if user.UID == 0 {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// GetUserByEmail function
func GetUserByEmail(email string) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where email = $1"
	row, err := con.Query(sql, email)
	if err != nil {
		return User{}, err
	}
	defer row.Close()
	var user User
	for row.Next() {
		err := row.Scan(&user.UID, &user.Nickname, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}
	}
	if user.UID == 0 {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// UpdateUser function to update a user
func UpdateUser(user User) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update users set nickname = $1, email = $2, status = $3 where uid = $4"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	rs, err := stmt.Exec(user.Nickname, user.Email, user.Status, user.UID)
	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()
}

// DeleteUser function to delete a user
func DeleteUser(uid uint32) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "delete from users where uid = $1"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	rs, err := stmt.Exec(uid)
	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()
}
