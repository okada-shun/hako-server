package api

import (
	"net/http"

	"hako-server/database"
	"hako-server/hkfinance"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	DB *database.GormDatabase
	Tx *hkfinance.HkfinanceTx
}

func (a *UserAPI) GetUserInfo(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	userInfo, err := a.Tx.GetUserInfo(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, userInfo)
}

func (a *UserAPI) GetUserInfoFromDB(ctx *gin.Context) {
	userAddress := ctx.Request.Header.Get("x-address")
	userInfo, err := a.DB.GetUserInfo(userAddress)
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, userInfo)
}
