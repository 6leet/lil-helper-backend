package e

var msgMap = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "internal server error",
	INVALID_PARAMS: "invalid parameters",
	UNAUTHORIZED:   "unauthorized",
	FORBIDDEN:      "permission denied",
	ALREADY_EXIST:  "instance already exist",

	ERR_USER_EXIST:              "user already exist",
	ERR_INVALID_USER_UID:        "invalid user uid",
	ERR_GAME_TYPE_EXIST:         "game type already exist",
	ERR_INVALID_GAME_TYPE_UID:   "invalid game type uid",
	ERR_INVALID_PAGE:            "invalid page value",
	ERR_INVALID_PAGE_LIMIT:      "invalid page limit value",
	ERR_INVALID_GAME_UID:        "invalid game uid",
	ERR_INVALID_MEMBER_UID:      "invalid member uid",
	ERR_INVALID_QUERY_TYPE:      "invalid query type",
	ERR_INVALID_TIME_STAMP:      "invalid query timestamp",
	ERR_INVALID_TRANSACTION_UID: "invalid transaction uid",
	ERR_INVALID_ROUND_TIME:      "invalid round time",
	ERR_INVALID_BET_UID:         "invalid bet uid",
	ERR_NO_SUCH_BET:             "no such bet",
	ERR_NO_SUCH_GAME:            "no such game",
	ERR_NO_SUCH_TRANSACTION:     "no such transaction",
	ERR_NO_SUCH_MEMBER:          "no such member",
	ERR_NO_SUCH_GAME_INSTANCE:   "no such game instance",
	ERR_EXCEED_BET_LIMIT:        "exceed bet limit",
	ERR_INVALID_BET_PARAM:       "invalid bet parameter",
}

func GetMsg(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}

	return msgMap[ERROR]
}
