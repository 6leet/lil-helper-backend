package e

import "fmt"

var (
	ErrInvalidLoginParameters  = fmt.Errorf("invalid username or password")
	ErrUserExist               = fmt.Errorf("user already exist")
	ErrUserNotExist            = fmt.Errorf("no such user")
	ErrInvalidUserUID          = fmt.Errorf("invalid user uid")
	ErrGameExist               = fmt.Errorf("game already exist")
	ErrMissionNotExist         = fmt.Errorf("no such mission")
	ErrRoundPrefixExist        = fmt.Errorf("round prefix has been used")
	ErrGameTypeExist           = fmt.Errorf("game type already exist")
	ErrGameTypeNotExist        = fmt.Errorf("no such game type")
	ErrBetNotExist             = fmt.Errorf("no such bet")
	ErrMemberNotExist          = fmt.Errorf("no such member")
	ErrMemberExist             = fmt.Errorf("member already exist")
	ErrIncorrectDepositAmount  = fmt.Errorf("deposit amount must be greater than 0")
	ErrIncorrectWithdrawAmount = fmt.Errorf("withdraw amount must be greater than 0")
	ErrMemberLocked            = fmt.Errorf("member has been locked")
	ErrInsufficientFund        = fmt.Errorf("insufficient fund")
	ErrTransactionNotExist     = fmt.Errorf("no such transaction")
	ErrHashidsInvalidLength    = fmt.Errorf("hashids invalid length")
	ErrGameInstanceNotExist    = fmt.Errorf("no such game instance")
	ErrExceedBetLimit          = fmt.Errorf("exceed bet limit")
	ErrInvalidBetParam         = fmt.Errorf("invalid bet parameter")
)
