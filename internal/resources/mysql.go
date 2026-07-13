package resources

import "goforge/internal/models"

type MySQL struct{}

func (MySQL) Name() string {
	return "mysql"
}
func (MySQL) Questions() []*models.Question {
	return []*models.Question{}
}
func (MySQL) BuildConfig(a models.Answers) models.Resource {
	return models.Resource{
		Name: "mysql",
		Params: map[string]any{
		},
	}
}
func (MySQL) Files(res models.Resource) []models.GenerateFile {

	files := []models.GenerateFile{
		{
			Template: "resources/database/mysql.go.tmpl",
			Output:   "internal/infrastructure/database/db.go",
		},
	}
	return files
}
