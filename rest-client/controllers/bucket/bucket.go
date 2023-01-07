package bucket

import (
	buckettable "api/database/bucket"
	bucketModel "api/models/bucket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	deleteBody := bucketModel.Bucket{}

	if err := ctx.ShouldBindJSON(&deleteBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if deleted := buckettable.Delete(deleteBody.ID); deleted {
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

	updateBody := bucketModel.Bucket{}

	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, buckettable.Update(updateBody))
}

func Create(ctx *gin.Context) {
	createBody := bucketModel.BucketCreate{}

	if err := ctx.ShouldBindJSON(&createBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := buckettable.Create(createBody)
	ctx.JSON(http.StatusCreated, created)
}

func Read(ctx *gin.Context) {
	readBody := bucketModel.BucketRead{}
	if err := ctx.ShouldBindJSON(&readBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, buckettable.Read(readBody.ID))
}
