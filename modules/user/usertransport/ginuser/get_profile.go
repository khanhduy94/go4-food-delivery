package ginuser

import (
	"demo/common"
	"demo/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSONP(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}