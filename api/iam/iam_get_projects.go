package iam

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/iamModels"
	"sbercloud-api/internal/handlers/requestMakers"
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
