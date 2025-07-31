package helper

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/toolkit/storage"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func DeleteFileHelper(
	ctx context.Context,
	path string,
) {
	stg, err := storage.NewStorage(ctx)
	if err != nil {
		log.WithContext(ctx).Error(err, "error init google cloud storage")
		return
	}
	defer util.CloseBuffer(stg.Client)

	err = stg.DeleteFile(ctx, path)
	if err != nil {
		log.WithContext(ctx).Error(err, "error delete file from google cloud storage")
		return
	}
}
