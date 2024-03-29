package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/hashids"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/utils"
	"sort"

	"github.com/jinzhu/gorm"
)

type PublicUser struct {
	UID      string `json:"userUID"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	Active   bool   `json:"active"`
	Score    int    `json:"score"`
	Nickname string `json:"nickname"`
	Exp      int    `json:"exp"`
	Createat string `json:"createat"`
	Level    uint   `json:"level"`
	Email    string `json:"email"`
}

func (u *User) Public() (pu PublicUser) {
	pu.UID = u.UID
	pu.Username = u.Username
	pu.Admin = u.Admin
	pu.Active = u.Active
	pu.Score = u.Score
	pu.Nickname = u.Nickname
	pu.Exp = u.Exp
	pu.Createat = u.CreatedAt.String()[0:10]
	pu.Level = utils.ExpToLevel(u.Exp)
	pu.Email = u.Email
	return pu
}

func RegistUser(username string, password string, email string, nickname string, admin bool) (*User, error) {
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	passwordMD5 := utils.MD5V([]byte(password))
	var user = User{
		Username:    username,
		Password:    passwordMD5,
		Admin:       admin,
		Active:      false,
		Score:       0,
		Email:       email,
		Certificate: false,
		Nickname:    nickname,
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

func GetUsers(active bool, admin bool, all bool, ifsort bool, keyword string) ([]User, error) {
	users := []User{}

	query := db.LilHelperDB
	if !all {
		query = query.Where("active = ? AND admin = ?", active, admin)
	}
	query = query.Where("username LIKE ?", keyword)
	if err := query.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("query users failed: %w", err)
	}
	if ifsort {
		sort.SliceStable(users, func(i, j int) bool {
			return users[i].Score > users[j].Score
		})
	}
	return users, nil
}

func BanUser(id uint) error {
	query := db.LilHelperDB
	if err := query.Find(&User{}, id).Update("active", false).Error; err != nil {
		return fmt.Errorf("ban user failed: %w", err)
	}
	return nil
}

func SetUserScoreExp(id uint, missionID uint, addvar int) (*User, error) {
	user := User{}
	mission := Mission{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	fmt.Println(missionID)
	if err := tx.Find(&mission, missionID).Error; err != nil {
		return nil, fmt.Errorf("query mission failed: %w", err)
	}
	scorevar := mission.Score * addvar
	expvar := mission.Exp * addvar
	tx = tx.Find(&user, id)
	updateUser := map[string]interface{}{
		"score": user.Score + scorevar,
		"exp":   user.Exp + expvar,
	}
	fmt.Println(scorevar, mission.Score, addvar)
	if err := tx.Updates(updateUser).Error; err != nil {
		return nil, fmt.Errorf("user update failed: %w", err)
	}
	// if err := tx.Update("score", gorm.Expr("score + ?", scorevar)).Error; err != nil {
	// 	return nil, fmt.Errorf("user update failed: %w", err)
	// }
	tx.Commit()
	return &user, nil
}

func UpdateUser(id uint, password string, email string, nickname string) (*User, error) {
	user := User{}
	var passwordMD5 string
	if password == "" {
		passwordMD5 = ""
	} else {
		passwordMD5 = utils.MD5V([]byte(password))
	}
	updateUser := User{
		Password: passwordMD5,
		Nickname: nickname,
	}

	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Find(&user, id).Updates(updateUser).Error; err != nil {
		return nil, fmt.Errorf("user update failed: %w", err)
	}

	tx.Commit()
	return &user, nil
}
