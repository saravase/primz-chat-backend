package org

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"

	"github.com/saravase/primz-chat-backend/model"
)

// Create godoc
// @Summary      create new organization
// @Description  Create new organization
// @Tags         orgs
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param org body OrgCreateReq true "Organization Detail"
// @Success      200  {object}  handler.CreateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      409  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/org/ [post]
func (h *Handler) Create(c *gin.Context) {
	var (
		req    OrgCreateReq
		org    model.Org
		apperr *apperrors.Error
	)
	if ok := handler.BindData(c, &req); !ok {
		return
	}
	copier.Copy(&org, &req)

	ctx := c.Request.Context()
	org_, err := h.OrgService.GetByName(ctx, org.Name)
	if err != nil && !errors.As(err, &apperr) {
		log.Printf("Failed to get organization: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	if org_ != nil {
		e := apperrors.NewConflict("org name", org.Name)
		log.Printf("Failed to create organization: %v\n", e.Error())
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	err = h.OrgService.Create(ctx, &org)
	if err != nil {
		log.Printf("Failed to create organization: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &handler.CreateResponse{
		ID: org.OrgID,
	})
}
