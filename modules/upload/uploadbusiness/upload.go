package uploadbusiness

import (
	"bytes"
	"context"
	"demo/common"
	"demo/component/upload"
	"demo/modules/upload/uploadmodel"
	"demo/modules/user/userbiz"
	"demo/modules/user/usermodel"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider  upload.Provider
	imgStore  CreateImageStorage
	userStore userbiz.UserStorage
}

func NewUploadBiz(provider upload.Provider, imgStore CreateImageStorage, userStore userbiz.UserStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore, userStore: userStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string, imgFor *uploadmodel.ImgFor) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	if imgFor.Model == "user" {
		if user, _ := biz.userStore.FindUser(ctx, map[string]interface{}{"email": imgFor.Identify}); user != nil {
			userCreate := usermodel.UserCreate{SQLModel: common.SQLModel{Id: user.Id}, Avatar: img}
			biz.userStore.UpdateUser(ctx, &userCreate)
		}
	}

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
