package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Name string `json:"name"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} Resp
// @Router /ping [get]
func pingH(c *gin.Context) {
	resp := Resp{Name: "pong"}
	c.JSON(http.StatusOK, resp)
}
