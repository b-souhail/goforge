package resources

import "goforge/internal/models"

var Registry = map[string]models.ResourceDefinition{
    "http":  HTTP{},
    "mysql": MySQL{},
}