package iam

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/iamModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type listProjectResponse struct {
	Projects []iamModels.ProjectModel `json:"projects"`
}

func ListProjects() ([]iamModels.ProjectModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.IamEndpoint) + "/v3/auth/projects")
	var resp listProjectResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Projects, err
}
