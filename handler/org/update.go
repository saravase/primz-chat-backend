package org

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Update godoc
// @Summary      update organization detail based on org id
// @Description  Update organization detail based on org id
// @Tags         orgs
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param        org_id   path     string  true  "Org ID"
// @Param org body OrgUpdateReq true "Organization Detail"
// @Success      200  {object}  []UpdateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/orgs/{org_id} [put]
func (h *Handler) Update(c *gin.Context) {
	var (
		req OrgUpdateReq
	)

	id := c.Param("org_id")
	if ok := handler.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	org, err := h.OrgService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get organization: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	copier.CopyWithOption(org, &req, copier.Option{IgnoreEmpty: true})

	err = h.OrgService.Update(ctx, id, org)
	if err != nil {
		log.Printf("Failed to update organization: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &UpdateResponse{
		IsUpdated: true,
	})
}
