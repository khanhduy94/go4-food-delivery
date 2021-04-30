package upload

import (
	"context"
	"demo/common"
)

type Provider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
