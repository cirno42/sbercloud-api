package evs

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/internal/handlers/requestMakers"
)

func DeleteDisk(projectID, volumeID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/cloudvolumes/%s", projectID, volumeID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}
