package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/pkg/e"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/speps/go-hashids"
)

var HID *hashids.HashID

const hidAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
const hidMinLength int = 24

const salt = "{email_token} token salt."

func EncodeToken(userID uint) (string, error) {
	return HID.EncodeInt64([]int64{int64(userID), int64(rand.Intn(10000))})
}
func DecodeToken(token string) (userID uint, err error) {
	infos, err := HID.DecodeInt64WithError(token)
	if err != nil {
		return 0, err
	} else if len(infos) != 2 {
		return 0, e.ErrHashidsInvalidLength
	}
	userID = uint(infos[0])
	return userID, nil
}

func HandleToken(token string) error {
	userID, err := DecodeToken(token)
	if err != nil {
		return err
	}
	emailtoken := Emailtoken{}
	query := db.LilHelperDB
	if err := query.Where("user_id = ?", userID).First(&emailtoken).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("haven't regist yet: %w", err)
	}
	if time.Now().After(emailtoken.Expireat) {
		UpdateToken(userID)
		// Send email with new token link
	} else if token != emailtoken.Token {
		// wrong certification
		// a click if need a new token link => a route that leads to UpdateToken
		// Send email with new token link
	} else {
		// success
		// active -> true (or certificate - > true?, should add login condition at jwt)
	}
	return nil
}

func CreateToken(userID uint) (*Emailtoken, error) {
	token, err := EncodeToken(userID)
	if err != nil {
		return nil, fmt.Errorf("token generate failed: %w", err)
	}
	emailtoken := Emailtoken{
		UserID:   userID,
		Token:    token,
		Expireat: time.Now().Add(time.Hour * 24),
	}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Create(&emailtoken).Error; err != nil {
		return nil, fmt.Errorf("token creation failed: %w", err)
	}
	tx.Commit()

	return &emailtoken, nil
}

func UpdateToken(userID uint) (*Emailtoken, error) {
	token, err := EncodeToken(userID)
	if err != nil {
		return nil, fmt.Errorf("token generate failed: %w", err)
	}
	emailtoken := Emailtoken{}
	updateEmailtoken := Emailtoken{
		Token:    token,
		Expireat: time.Now().Add(time.Hour * 24),
	}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Model(&token).Updates(updateEmailtoken).Error; err != nil {
		return nil, fmt.Errorf("token update failed: %w", err)
	}
	tx.Commit()

	return &emailtoken, nil
}
