package e

type Err struct {
	Err  string
	Code int
}

func (e *Err) Error() string {
	return e.Err
}

func NewErr(code int) *Err {
	return &Err{Code: code, Err: GetMsg(code)}
}

func WrapErr(err error) *Err {
	return &Err{Code: ERROR, Err: err.Error()}
}

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	UNAUTHORIZED   = 401
	FORBIDDEN      = 403
	ALREADY_EXIST  = 409

	ERR_USER_EXIST              = 10001
	ERR_INVALID_USER_UID        = 10002
	ERR_GAME_TYPE_EXIST         = 10003
	ERR_INVALID_GAME_TYPE_UID   = 10004
	ERR_INVALID_PAGE            = 10005
	ERR_INVALID_PAGE_LIMIT      = 10006
	ERR_INVALID_GAME_UID        = 10007
	ERR_INVALID_MEMBER_UID      = 10008
	ERR_INVALID_QUERY_TYPE      = 10009
	ERR_INVALID_TIME_STAMP      = 10010
	ERR_INVALID_TRANSACTION_UID = 10011
	ERR_INVALID_ROUND_TIME      = 10012
	ERR_INVALID_BET_UID         = 10013
	ERR_NO_SUCH_BET             = 10014
	ERR_NO_SUCH_MISSION         = 10015
	ERR_NO_SUCH_TRANSACTION     = 10016
	ERR_NO_SUCH_SCREENSHOT      = 10017
	ERR_NO_SUCH_GAME_INSTANCE   = 10018
	ERR_EXCEED_BET_LIMIT        = 10019
	ERR_INVALID_BET_PARAM       = 10020
)
