package payload

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/model"
)

type FileResponse struct {
	GUID        string           `json:"guid"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	File        *api.FilePayload `json:"file"`
	Size        int64            `json:"size"`
	Extension   string           `json:"extension"`
	MimeType    string           `json:"mime_type"`
	CreatedAt   string           `json:"created_at"`
	CreatedBy   *string          `json:"created_by"`
	UpdatedAt   *string          `json:"updated_at"`
	UpdatedBy   *string          `json:"updated_by"`
}

func ToFileResponse(entity model.File) *FileResponse {
	response := FileResponse{
		GUID:      entity.GUID,
		Name:      entity.Name,
		File:      api.ToFile(entity.Path),
		Size:      entity.Size,
		Extension: entity.Extension,
		MimeType:  entity.MimeType,
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
	}

	if entity.Description.Valid {
		description := entity.Description.String
		response.Description = &description
	}

	if entity.CreatedBy.Valid {
		createdBy := entity.CreatedBy.String
		response.CreatedBy = &createdBy
	}

	if entity.UpdatedAt.Valid {
		updatedAt := entity.UpdatedAt.Time.Format(time.RFC3339)
		response.UpdatedAt = &updatedAt
	}

	if entity.UpdatedBy.Valid {
		updatedBy := entity.UpdatedBy.String
		response.UpdatedBy = &updatedBy
	}

	return &response
}

func ToFileResponses(entities []model.File) *[]*FileResponse {
	response := make([]*FileResponse, len(entities))

	for i := range entities {
		response[i] = ToFileResponse(entities[i])
	}

	return &response
}
