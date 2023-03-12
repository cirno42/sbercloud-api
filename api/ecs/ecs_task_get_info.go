package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/ecsModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

func GetInfoAboutTask(projectID, jobID string) (*ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}
