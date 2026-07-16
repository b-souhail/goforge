package resources

import (
	"goforge/internal/models"
	"goforge/internal/resources/http"
	"goforge/internal/resources/mysql"
)

var Registry = map[string]models.ResourceDefinition{
	"http":  http.HTTP{},
	"mysql": mysql.MySQL{},
}
