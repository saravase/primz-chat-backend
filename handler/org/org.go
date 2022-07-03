package org

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Org godoc
// @Summary      get organization detail based on org id
// @Description  Get organization detail based on org id
// @Tags         orgs
// @Produce      json
// @Security ApiKeyAuth
// @Param        org_id   path     string  true  "Org ID"
// @Success      200  {object}  model.Org
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/org/{org_id} [get]
func (h *Handler) Org(c *gin.Context) {
	id := c.Param("org_id")
	ctx := c.Request.Context()
	org, err := h.OrgService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get organization: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, org)
}
