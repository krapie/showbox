package handlers

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/krapie/showbox/server/database"
)

type APIHandler struct {
	ctx context.Context
	db  *database.DB
}

func New(ctx context.Context) (*APIHandler, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &APIHandler{
		ctx: ctx,
		db:  db,
	}, nil
}

func (h *APIHandler) GetShow(c *gin.Context) {

}

func (h *APIHandler) RegisterShow(c *gin.Context) {

}
