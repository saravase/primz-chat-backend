package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Search godoc
// @Summary      get users detail based on search query
// @Description  Get users detail based on search query
// @Tags         auth
// @Produce      json
// @Security ApiKeyAuth
// @Param   org_id  query     string     false  "Organization ID"
// @Param   dept_id  query     string     false  "Department ID"
// @Param   group_id  query     string     false  "Group ID"
// @Param   name  query     string     false  "Name filter"
// @Success      200  {object}  []model.User
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/users/search [get]
func (h *Handler) Search(c *gin.Context) {
	queryMap := make(map[string]string)
	if orgID, ok := c.GetQuery("org_id"); ok {
		queryMap["org_id"] = orgID
	}
	if deptID, ok := c.GetQuery("dept_id"); ok {
		queryMap["dept_id"] = deptID
	}
	if groupID, ok := c.GetQuery("group_id"); ok {
		queryMap["group_id"] = groupID
	}
	if name, ok := c.GetQuery("name"); ok {
		queryMap["name"] = name
	}
	ctx := c.Request.Context()
	orgs, err := h.UserService.SearchUsers(ctx, queryMap)
	if err != nil {
		log.Printf("Failed to get all users: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, orgs)
}
