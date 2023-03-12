package evs

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/evsModels"
	"sbercloud-api/internal/handlers/requestMakers"
)

func GetInfoAboutBatchTask(projectID, jobID string) (*evsModels.CreateBatchDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job evsModels.CreateBatchDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}

func GetInfoAboutSingleTask(projectID, jobID string) (*evsModels.CreateSingleDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job evsModels.CreateSingleDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}
