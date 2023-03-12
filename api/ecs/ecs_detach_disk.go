package ecs

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/internal/handlers/requestMakers"
)

func DetachDiskEcs(projectID, ecsId, volumeId string, deleteFlag int) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/detachvolume/%s?delete_flag=%d", projectID, ecsId, volumeId, deleteFlag)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}
