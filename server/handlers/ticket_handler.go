package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetTickets(c *gin.Context) {
	showId, err := primitive.ObjectIDFromHex(c.Query("showId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tickets, err := h.ticketSrv.GetAvailableTickets(context.Background(), showId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *Handler) ReserveTicket(c *gin.Context) {
	ticketID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID, err := primitive.ObjectIDFromHex(c.GetHeader("User-Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err = h.ticketSrv.ReserveTicket(ticketID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) BuyTicket(c *gin.Context) {
	ticketID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID, err := primitive.ObjectIDFromHex(c.GetHeader("User-Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err = h.ticketSrv.BuyTicket(context.Background(), ticketID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
