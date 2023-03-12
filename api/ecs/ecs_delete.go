package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/ecsModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type serverToDelete struct {
	Id string `json:"id"`
}

type deleteECSRequest struct {
	Servers        []serverToDelete `json:"servers"`
	DeletePublicIp bool             `json:"delete_publicip"`
	DeleteVolume   bool             `json:"delete_volume"`
}

func DeleteEcs(projectID string, ecsIdsToDelete []string, deletePublicIp, deleteVolume bool) (ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/delete", projectID)

	serversIds := make([]serverToDelete, len(ecsIdsToDelete))
	for i := 0; i < len(ecsIdsToDelete); i++ {
		serversIds[i].Id = ecsIdsToDelete[i]
	}
	deleteRequest := deleteECSRequest{
		Servers:        serversIds,
		DeletePublicIp: deletePublicIp,
		DeleteVolume:   deleteVolume,
	}
	var job ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, deleteRequest, &job, nil)
	return job, err
}
