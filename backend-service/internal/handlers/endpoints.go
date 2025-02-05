package handlers

import (
	"log"
	"net/http"

	"github.com/DurkaVerder/models"
	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerGetAllPing(ctx *gin.Context) {
	pings, err := h.service.GetAllPing()
	if err != nil {
		log.Printf("Error getting all pings: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pings)
}

func (h *HandlersManager) HandlerAddPing(ctx *gin.Context) {
	var newPing models.PingResult
	if err := ctx.BindJSON(&newPing); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTablePings(newPing); err != nil {
		log.Printf("Error updating table pings: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ping added successfully"})
}
