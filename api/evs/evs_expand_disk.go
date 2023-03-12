package evs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type expandDiskRequestParameters struct {
	NewSize int `json:"new_size"`
}

type expandDiskRequest struct {
	OsExtend expandDiskRequestParameters `json:"os-extend"`
}

func ExpandDisk(projectID, volumeID string, newSize int) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/volumes/%s/action", projectID, volumeID)
	params := expandDiskRequestParameters{
		NewSize: newSize,
	}
	request := expandDiskRequest{
		OsExtend: params,
	}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, nil, nil)
	return err
}
