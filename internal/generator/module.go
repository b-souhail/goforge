package generator

import (
	"fmt"
	"goforge/internal/models"
)

func ModuleFiles(module string) []models.GenerateFile {
	return []models.GenerateFile{
		{
			Template: "clean/domain/entity/entity.go.tmpl",
			Output:   fmt.Sprintf("internal/domain/entity/%s.go", module),
		},
		{
			Template: "clean/domain/repository/repository.go.tmpl",
			Output:   fmt.Sprintf("internal/domain/repository/%s_repo.go", module),
		},
		{
			Template: "clean/application/DTOs/dto.go.tmpl",
			Output:   fmt.Sprintf("internal/application/dtos/%s_dto.go", module),
		},
		{
			Template: "clean/application/usecases/usecases.go.tmpl",
			Output:   fmt.Sprintf("internal/application/usecases/%s_usecase.go", module),
		},
		{
			Template: "clean/infrastructure/repository/repo_impl.go.tmpl",
			Output:   fmt.Sprintf("internal/infrastructure/repository/%s_repo_impl.go", module),
		},
	}
}
