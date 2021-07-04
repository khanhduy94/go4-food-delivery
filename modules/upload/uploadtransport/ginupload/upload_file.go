package ginupload

import (
	"demo/common"
	"demo/component"
	"demo/modules/upload/uploadbusiness"
	"demo/modules/upload/uploadmodel"
	"demo/modules/upload/uploadstorage"
	"demo/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")
		model, _ := c.GetPostForm("model")
		identify, _ := c.GetPostForm("identify")

		imgFor := uploadmodel.NewImgFor(model, identify)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		db := appCtx.GetMainDBConnection()
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		userStorage := userstorage.NewSQLStore(db)
		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbusiness.NewUploadBiz(appCtx.GetUploadProvider(), imgStore, userStorage)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename, imgFor)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
