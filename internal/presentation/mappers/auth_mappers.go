package mappers

import (
	"project/internal/application/dtos"
	"project/internal/presentation/schemas"
)

func FromCreatedSchemaToDTO(schema *schemas.CreateUserSchema) *dtos.UserCreateDTO {
	return &dtos.UserCreateDTO{
		UserName:     schema.Username,
		UserPassword: schema.Password,
		UserEmail:    schema.Email,
	}
}

func FromCreatedDTOToSchema(dto *dtos.UserCreatedDTO) *schemas.UserCreatedSchema {
	return &schemas.UserCreatedSchema{
		UserID: dto.UserId,
	}
}

func FromAuthorizeSchemaToDTO(schema *schemas.AuthorizeSchema) *dtos.AuthorizeDTO {
	return &dtos.AuthorizeDTO{
		Username: schema.Username,
		Password: schema.Password,
	}
}

func FromTokensDTOToSchema(dto *dtos.TokensDTO) *schemas.TokensSchema {
	return &schemas.TokensSchema{
		AccessToken: dto.AccessToken,
	}
}
