package instance

import (
	instancetable "api/database/instance"
	instanceModel "api/models/instance"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	deleteBody := instanceModel.Instance{}

	if err := ctx.ShouldBindJSON(&deleteBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if deleted := instancetable.Delete(deleteBody.ID); deleted {
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

	updateBody := instanceModel.Instance{}

	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, instancetable.Update(updateBody))
}

func Create(ctx *gin.Context) {
	createBody := instanceModel.InstanceCreate{}

	if err := ctx.ShouldBindJSON(&createBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := instancetable.Create(createBody)
	ctx.JSON(http.StatusCreated, created)
}

func Read(ctx *gin.Context) {
	readBody := instanceModel.InstanceRead{}
	if err := ctx.ShouldBindJSON(&readBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, instancetable.Read(readBody.ID))
}
