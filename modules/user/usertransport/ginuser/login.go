package ginuser

import (
	"demo/common"
	"demo/component"
	"demo/component/hasher"
	"demo/component/tokenprovider/jwt"
	"demo/modules/user/userbiz"
	"demo/modules/user/usermodel"
	"demo/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		bussiness := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := bussiness.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSONP(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
