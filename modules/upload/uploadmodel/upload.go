package uploadmodel

import (
	"demo/common"
	"errors"
)

const EntityName = "Upload"

type ImgFor struct {
	Model    string
	Identify string
}

func NewImgFor(model string, identify string) *ImgFor {
	return &ImgFor{
		Model:    model,
		Identify: identify,
	}
}

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

//
//func (u *Upload) Mask(isAdmin bool) {
//	u.GenUID(common.DBTypeUpload, 1)
//}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
