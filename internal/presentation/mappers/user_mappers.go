package mappers

import (
	"project/internal/application/dtos"
	"project/internal/presentation/schemas"
)

func FromUserInfoDTOToSchema(dto *dtos.UserInfoDTO) *schemas.UserInfoSchema {
	return &schemas.UserInfoSchema{
		UserID:    dto.UserID,
		Username:  dto.Username,
		UserEmail: dto.UserEmail,
	}
}
