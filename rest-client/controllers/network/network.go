package network

import (
	networktable "api/database/network"
	networkModel "api/models/network"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	deleteBody := networkModel.Network{}

	if err := ctx.ShouldBindJSON(&deleteBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if deleted := networktable.Delete(deleteBody.ID); deleted {
		ctx.JSON(http.StatusOK, gin.H{
			"deleted": deleted,
			"data":    deleteBody,
		})
	}

	ctx.JSON(http.StatusConflict, gin.H{
		"deleted": false,
		"data":    deleteBody,
	})

}

func Update(ctx *gin.Context) {

	updateBody := networkModel.Network{}

	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateBody.IPList == nil || updateBody.InstanceList == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IPList and InstanceList must be array!"})
		return
	}
	ctx.JSON(http.StatusOK, networktable.Update(updateBody))
}

func Create(ctx *gin.Context) {
	createBody := networkModel.NetworkCreate{}

	if err := ctx.ShouldBindJSON(&createBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createBody.IPList == nil || createBody.InstanceList == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IPList and InstanceList must be array!"})
		return
	}

	created := networktable.Create(createBody)
	ctx.JSON(http.StatusOK, created)
}

func Read(ctx *gin.Context) {
	readBody := networkModel.NetworkRead{}
	if err := ctx.ShouldBindJSON(&readBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, networktable.Read(readBody.ID))
}
