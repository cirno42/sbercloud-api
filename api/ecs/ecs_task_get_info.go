package ecs

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/ecsModels"
	"sbercloud-api/internal/handlers/requestMakers"
)

func GetInfoAboutTask(projectID, jobID string) (*ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}
