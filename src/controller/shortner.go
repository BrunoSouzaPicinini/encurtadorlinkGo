package controller

import (
	"encurtadorLink/src/model"
	"encurtadorLink/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Slug string
	Url  string
}

func Create(ctx *gin.Context) {
	var short model.Shortener

	if err := ctx.ShouldBindBodyWithJSON(&short); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	createdShort, err := service.Create(short)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create shortener",
		})
		return
	}

	ctx.JSON(201, createdShort)
}

func RedirectBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	short, err := service.FindBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Shortener not found",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, short.OriginalUrl)
}
