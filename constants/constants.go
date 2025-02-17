package constants

import "time"

const (
	TestMode  = "test"
	DebugMode = "debug"

	JWTRefreshTokenExpireDuration = time.Hour * 72
	JWTAccessTokenExpireDuration  = time.Minute * 60
	ContextTimeoutDuration        = time.Second * 7

	Success     = "success"
	Active      = 1
	InActive    = 0
	BotToken    = "bot token"
	TgChannelID = "channel id"

	SELECT_LIST = "select_list"
	SELECT_ONE  = "select_one"
	INSERT_ITEM = "insert_item"
	UPDATE_ITEM = "update_item"
)

const (
	IncomeTransactionID = iota
	ExpenseTransactionID
	TransferTransactionID
)
