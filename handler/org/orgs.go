package org

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Orgs godoc
// @Summary      get all organization details
// @Description  Get all organization details
// @Tags         orgs
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  []model.Org
// @Failure      500  {object}  apperrors.Error
// @Router       /api/orgs [get]
func (h *Handler) Orgs(c *gin.Context) {
	ctx := c.Request.Context()
	orgs, err := h.OrgService.GetOrgs(ctx)
	if err != nil {
		log.Printf("Failed to get all organizations: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, orgs)
}
