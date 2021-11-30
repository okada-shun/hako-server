package api

import (
	"net/http"

	"hako-server/database"
	"hako-server/hkfinance"

	"github.com/gin-gonic/gin"
)

type HakoAPI struct {
	DB *database.GormDatabase
	Tx *hkfinance.HkfinanceTx
}

func (a *HakoAPI) GetHakoInfo(ctx *gin.Context) {
	hakoInfo, err := a.Tx.GetHakoInfo()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, hakoInfo)
}

func (a *HakoAPI) GetHakoInfoFromDB(ctx *gin.Context) {
	hakoInfo, err := a.DB.GetHakoInfo()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, hakoInfo)
}
