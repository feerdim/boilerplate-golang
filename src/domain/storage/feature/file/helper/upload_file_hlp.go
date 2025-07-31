package helper

import (
	"context"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/payload"
	"github.com/feerdim/boilerplate-golang/src/toolkit/storage"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func UploadFileHelper(
	ctx context.Context,
	request payload.FileRequest,
) (filePath string, err error) {
	stg, err := storage.NewStorage(ctx)
	if err != nil {
		log.WithContext(ctx).Error(err, "error init google cloud storage")
		return
	}
	defer util.CloseBuffer(stg.Client)

	filePath, err = stg.UploadFile(ctx, request.File, request.Path, time.Now().Unix())
	if err != nil {
		log.WithContext(ctx).Error(err, "error upload file to google cloud storage")
		return
	}

	return
}
