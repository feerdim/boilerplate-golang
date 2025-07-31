package helper

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
)

func GetRangeHeaderHelper(
	ctx context.Context,
	rangeHeader string,
	attrs *storage.ObjectAttrs,
) (start, end int64, err error) {
	const invalidRangeHeader = 2

	ranges := strings.Split(rangeHeader, "=")
	if len(ranges) != invalidRangeHeader || ranges[0] != "bytes" {
		err = constant.ErrRangeHeaderInvalid
		log.WithContext(ctx).Error(err, "range header must be 2 and first part must be bytes")

		return
	}

	rangeParts := strings.Split(ranges[1], "-")
	if len(rangeParts) != invalidRangeHeader {
		err = constant.ErrRangeHeaderInvalid
		log.WithContext(ctx).Error(err, "range parts must be 2")

		return
	}

	start, err = strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		err = log.WithContext(ctx).NewError(err, constant.ErrRangeHeaderInvalid)
		return
	}

	if rangeParts[1] == "" {
		end = attrs.Size - 1
	} else {
		end, err = strconv.ParseInt(rangeParts[1], 10, 64)
		if err != nil {
			err = log.WithContext(ctx).NewError(err, constant.ErrRangeHeaderInvalid)
			return
		}
	}

	if start >= end || end >= attrs.Size {
		err = constant.ErrRequestedRangeNotSatisfiable
		log.WithContext(ctx).Error(err, "start must be less than end and end must be less than file size")

		return
	}

	return
}
