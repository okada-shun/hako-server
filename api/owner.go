package api

import (
	"net/http"

	"hako-server/database"
	"hako-server/hkfinance"

	"github.com/gin-gonic/gin"
)

type OwnerAPI struct {
	DB *database.GormDatabase
	Tx *hkfinance.HkfinanceTx
}

func (a *OwnerAPI) GetOwnerInfo(ctx *gin.Context) {
	ownerInfo, err := a.Tx.GetOwnerInfo()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, ownerInfo)
}

func (a *OwnerAPI) GetOwnerInfoFromDB(ctx *gin.Context) {
	ownerInfo, err := a.DB.GetOwnerInfo()
	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}
	ctx.JSON(http.StatusOK, ownerInfo)
}
