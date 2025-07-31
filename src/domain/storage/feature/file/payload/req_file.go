package payload

import (
	"mime/multipart"

	"github.com/feerdim/boilerplate-golang/src/api"
)

type FileRequest struct {
	File          *multipart.FileHeader
	DirectoryGUID string  `form:"directory_guid" json:"directory_guid"`
	Name          string  `form:"name" json:"name"`
	Description   *string `form:"description" json:"description"`
	Path          string  `form:"path" json:"path"`
	Size          int64   `form:"size" json:"size"`
	Extension     string  `form:"extension" json:"extension"`
	MimeType      string  `form:"mime_type" json:"mime_type"`
}

type ReadFileListRequest struct {
	api.PaginationPayload
	DirectoryGUID string `query:"directory_guid"`
}

type UpdateFileRequest struct {
	api.GUIDPayload
	FileRequest
}

type OpenFileRequest struct {
	Path string `param:"path" validate:"required"`
}
