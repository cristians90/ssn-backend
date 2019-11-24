package models

type UserAvatarModel struct {
	BaseModel         `storm:"inline"`
	UserID            uint64
	BinaryImage       []byte
	BinaryContentType string
}
