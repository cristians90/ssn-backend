package models

import "time"

type PostModel struct {
	BaseModel `storm:"inline"`
	Content   string `json:"content"`
	CreatedBy uint64 `json:"createdBy"`
}

type PostModelForApi struct {
	ID         uint64    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Content    string    `json:"content"`
	CreatedBy  uint64    `json:"createdBy"`
}

func (post *PostModel) GetPostModelForApi() PostModelForApi {
	model := PostModelForApi{
		ID:         post.ID,
		CreatedAt:  post.CreatedAt,
		ModifiedAt: post.ModifiedAt,
		Content:    post.Content,
		CreatedBy:  post.CreatedBy,
	}

	return model
}
