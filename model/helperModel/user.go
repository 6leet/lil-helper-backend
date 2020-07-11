package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/utils"
)

type PublicUser struct {
	Username string `json:"username"`
}

func (u *User) Public() (pu PublicUser) {
	pu.Username = u.Username

	return pu
}

func RegistUser(username string, password string) (*User, error) {
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	password = utils.MD5V([]byte(password))
	var user = User{
		Username: username,
		Password: password,
	}
	if notExist := tx.Where("username = ?", username).First(&user).RecordNotFound(); !notExist {
		return nil, e.ErrUserExist
	}
	if err := tx.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	fmt.Println(user.ID)

	tx.Commit()
	return &user, nil
}
