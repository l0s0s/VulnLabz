package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "VulnLabz - Security Testing Lab",
		"version": "1.0.0",
	})
}
