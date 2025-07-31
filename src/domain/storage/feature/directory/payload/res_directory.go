package payload

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/model"
)

type DirectoryResponse struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	CreatedBy   *string `json:"created_by"`
	UpdatedAt   *string `json:"updated_at"`
	UpdatedBy   *string `json:"updated_by"`
}

func ToDirectoryResponse(entity model.Directory) (response DirectoryResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name
	response.CreatedAt = entity.CreatedAt.Format(time.RFC3339)

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	if entity.CreatedBy.Valid {
		response.CreatedBy = &entity.CreatedBy.String
	}

	if entity.UpdatedAt.Valid {
		updatedAt := entity.UpdatedAt.Time.Format(time.RFC3339)
		response.UpdatedAt = &updatedAt
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

func ToDirectoryResponses(entities []model.Directory) (response []DirectoryResponse) {
	response = make([]DirectoryResponse, len(entities))

	for i := range entities {
		response[i] = ToDirectoryResponse(entities[i])
	}

	return
}
