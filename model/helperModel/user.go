package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/hashids"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/utils"

	"github.com/jinzhu/gorm"
)

type PublicUser struct {
	UID      string `json:"userUID"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	Active   bool   `json:"active"`
}

func (u *User) Public() (pu PublicUser) {
	pu.UID = u.UID
	pu.Username = u.Username
	pu.Admin = u.Admin
	pu.Active = u.Active
	return pu
}

func RegistUser(username string, password string, admin bool) (*User, error) {
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	passwordMD5 := utils.MD5V([]byte(password))
	var user = User{
		Username: username,
		Password: passwordMD5,
		Admin:    admin,
	}
	if notExist := tx.Where("username = ?", username).First(&user).RecordNotFound(); !notExist {
		return nil, e.ErrUserExist
	}
	if err := tx.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("create user failed: %w", err)
	}
	if uid, err := hashids.EncodeUserUID(user.ID); err != nil {
		return nil, fmt.Errorf("user uid generation failed: %w", err)
	} else if err = tx.Model(&user).Update("uid", uid).Error; err != nil {
		return nil, fmt.Errorf("update user uid failed: %w", err)
	}

	tx.Commit()
	return &user, nil
}

func Login(username string, password string) (*User, error) {
	user := User{}
	passwordMD5 := utils.MD5V([]byte(password))

	query := db.LilHelperDB.Where("username = ? AND password = ?", username, passwordMD5)
	if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, e.ErrInvalidLoginParameters
	} else if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	return &user, nil
}

func GetUser(id uint) (*User, error) {
	user := User{}

	query := db.LilHelperDB.Where("id = ?", id)
	if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, e.ErrUserNotExist
	} else if err != nil {
		return nil, fmt.Errorf("User query failed: %w", err)
	} else {
		return &user, nil
	}
}
