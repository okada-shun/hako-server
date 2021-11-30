package api

import (
	"net/http"

	"hako-server/database"
	"hako-server/hkfinance"

	"github.com/gin-gonic/gin"
)

type HistoryAPI struct {
	DB *database.GormDatabase
	Tx *hkfinance.HkfinanceTx
}

func (a *HistoryAPI) GetTransferTokenFromHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	transferTokenHistory, err := a.DB.GetTransferTokenFromHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, transferTokenHistory)
}

func (a *HistoryAPI) GetTransferTokenToHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	transferTokenHistory, err := a.DB.GetTransferTokenToHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, transferTokenHistory)
}

func (a *HistoryAPI) GetTransferCreditFromHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	transferCreditHistory, err := a.DB.GetTransferCreditFromHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, transferCreditHistory)
}

func (a *HistoryAPI) GetTransferCreditToHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	transferCreditHistory, err := a.DB.GetTransferCreditToHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, transferCreditHistory)
}

func (a *HistoryAPI) GetJoinHakoHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	joinHakoHistory, err := a.DB.GetJoinHakoHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, joinHakoHistory)
}

func (a *HistoryAPI) GetLeaveHakoHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	leaveHakoHistory, err := a.DB.GetLeaveHakoHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, leaveHakoHistory)
}

func (a *HistoryAPI) GetDepositTokenHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	depositTokenHistory, err := a.DB.GetDepositTokenHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, depositTokenHistory)
}

func (a *HistoryAPI) GetWithdrawTokenHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	withdrawTokenHistory, err := a.DB.GetWithdrawTokenHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, withdrawTokenHistory)
}

func (a *HistoryAPI) GetRegisterBorrowingHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	registerBorrowingHistory, err := a.DB.GetRegisterBorrowingHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, registerBorrowingHistory)
}

func (a *HistoryAPI) GetLendCreditFromHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	lendCreditHistory, err := a.DB.GetLendCreditFromHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, lendCreditHistory)
}

func (a *HistoryAPI) GetLendCreditToHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	lendCreditHistory, err := a.DB.GetLendCreditToHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, lendCreditHistory)
}

func (a *HistoryAPI) GetCollectDebtFromCreditorHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	collectDebtFromHistory, err := a.DB.GetCollectDebtFromCreditorHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, collectDebtFromHistory)
}

func (a *HistoryAPI) GetCollectDebtFromDebtorHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	collectDebtFromHistory, err := a.DB.GetCollectDebtFromDebtorHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, collectDebtFromHistory)
}

func (a *HistoryAPI) GetReturnDebtToCreditorHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	returnDebtToHistory, err := a.DB.GetReturnDebtToCreditorHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, returnDebtToHistory)
}

func (a *HistoryAPI) GetReturnDebtToDebtorHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	returnDebtToHistory, err := a.DB.GetReturnDebtToDebtorHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, returnDebtToHistory)
}

func (a *HistoryAPI) GetCreateCreditHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	createCreditHistory, err := a.DB.GetCreateCreditHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, createCreditHistory)
}

func (a *HistoryAPI) GetReduceDebtHistory(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	reduceDebtHistory, err := a.DB.GetReduceDebtHistory(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, reduceDebtHistory)
}

func (a *HistoryAPI) GetChangeHakoOwnerHistory(ctx *gin.Context) {
	changeHakoOwnerHistory, err := a.DB.GetChangeHakoOwnerHistory()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, changeHakoOwnerHistory)
}

func (a *HistoryAPI) GetChangeUpperLimitHistory(ctx *gin.Context) {
	changeUpperLimitHistory, err := a.DB.GetChangeUpperLimitHistory()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, changeUpperLimitHistory)
}

func (a *HistoryAPI) GetGetRewardHistory(ctx *gin.Context) {
	getRewardHistory, err := a.DB.GetGetRewardHistory()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, getRewardHistory)
}
