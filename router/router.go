package router

import (
	"net/http"

	"hako-server/api"
	"hako-server/database"
	"hako-server/hkfinance"

	"github.com/gin-gonic/gin"
)

func CreateRouter(db *database.GormDatabase, hkfinanceTx *hkfinance.HkfinanceTx) *gin.Engine {
	router := gin.Default()
	hakoHandler := &api.HakoAPI{
		DB: db,
		Tx: hkfinanceTx,
	}
	userHandler := &api.UserAPI{
		DB: db,
		Tx: hkfinanceTx,
	}
	ownerHandler := &api.OwnerAPI{
		DB: db,
		Tx: hkfinanceTx,
	}
	historyHandler := &api.HistoryAPI{
		DB: db,
		Tx: hkfinanceTx,
	}
	router.GET("/", home)
	router.GET("/hakoinfo/get", hakoHandler.GetHakoInfo)
	router.GET("/userinfo/get", userHandler.GetUserInfo)
	router.GET("/ownerinfo/get", ownerHandler.GetOwnerInfo)
	router.GET("/hakoinfo/getfromdb", hakoHandler.GetHakoInfoFromDB)
	router.GET("/userinfo/getfromdb", userHandler.GetUserInfoFromDB)
	router.GET("/ownerinfo/getfromdb", ownerHandler.GetOwnerInfoFromDB)
	router.GET("/history/transfer_token_from", historyHandler.GetTransferTokenFromHistory)
	router.GET("/history/transfer_token_to", historyHandler.GetTransferTokenToHistory)
	router.GET("/history/transfer_credit_from", historyHandler.GetTransferCreditFromHistory)
	router.GET("/history/transfer_credit_to", historyHandler.GetTransferCreditToHistory)
	router.GET("/history/join_hako", historyHandler.GetJoinHakoHistory)
	router.GET("/history/leave_hako", historyHandler.GetLeaveHakoHistory)
	router.GET("/history/deposit_token", historyHandler.GetDepositTokenHistory)
	router.GET("/history/withdraw_token", historyHandler.GetWithdrawTokenHistory)
	router.GET("/history/register_borrowing", historyHandler.GetRegisterBorrowingHistory)
	router.GET("/history/lend_credit_from", historyHandler.GetLendCreditFromHistory)
	router.GET("/history/lend_credit_to", historyHandler.GetLendCreditToHistory)
	router.GET("/history/collect_debt_from_creditor", historyHandler.GetCollectDebtFromCreditorHistory)
	router.GET("/history/collect_debt_from_debtor", historyHandler.GetCollectDebtFromDebtorHistory)
	router.GET("/history/return_debt_to_creditor", historyHandler.GetReturnDebtToCreditorHistory)
	router.GET("/history/return_debt_to_debtor", historyHandler.GetReturnDebtToDebtorHistory)
	router.GET("/history/create_credit", historyHandler.GetCreateCreditHistory)
	router.GET("/history/reduce_debt", historyHandler.GetReduceDebtHistory)
	router.GET("/history/change_hako_owner", historyHandler.GetChangeHakoOwnerHistory)
	router.GET("/history/change_upper_limit", historyHandler.GetChangeUpperLimitHistory)
	router.GET("/history/get_reward", historyHandler.GetGetRewardHistory)
	return router
}

func home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
