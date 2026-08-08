package mappers

import (
	"project/internal/application/dtos"
	"project/internal/presentation/schemas"
)


func FromCreatedSchemaToDTO(schema *schemas.CreateUserSchema) *dtos.UserCreateDTO {
	return &dtos.UserCreateDTO{
		UserName: schema.Username,
		UserPassword: schema.Password,
		UserEmail: schema.Email,
	}
}