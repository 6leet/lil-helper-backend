package hashids

import (
	"lil-helper-backend/pkg/e"

	"github.com/speps/go-hashids"
)

var userHID *hashids.HashID

const hidAlphabet = "0123456789ABCDEF"
const hidMinLength int = 12

const userSalt = "{user_id} user salt."

func EncodeUserUID(userID uint) (string, error) {
	return userHID.EncodeInt64([]int64{int64(userID)})
}
func DecodeUserUID(uid string) (id uint, err error) {
	ids, err := userHID.DecodeInt64WithError(uid)
	if err != nil {
		return 0, err
	} else if len(ids) != 1 {
		return 0, e.ErrHashidsInvalidLength
	}
	id = uint(ids[0])
	return id, nil
}

func init() {
	var err error
	{
		hidData := hashids.NewData()
		hidData.Alphabet = hidAlphabet
		hidData.MinLength = hidMinLength
		hidData.Salt = userSalt
		userHID, err = hashids.NewWithData(hidData)
		if err != nil {
			panic(err)
		}
	}
}
