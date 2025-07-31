package payload

import "github.com/feerdim/boilerplate-golang/src/api"

type DirectoryRequest struct {
	DirectoryGUID *string `json:"directory_guid"`
	Name          string  `json:"name"`
	Description   *string `json:"description"`
}

type ReadDirectoryListRequest struct {
	api.PaginationPayload
}

type UpdateDirectoryRequest struct {
	api.GUIDPayload
	DirectoryRequest
}
