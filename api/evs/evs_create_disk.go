package evs

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/evsModels"
	"sbercloud-api/internal/handlers/requestMakers"
)

type volumeRequest struct {
	AvailabilityZone string `json:"availability_zone"`
	Name             string `json:"name"`
	VolumeType       string `json:"volume_type"`
	Count            int    `json:"count"`
	Size             int    `json:"size"`
	Multiattach      bool   `json:"multiattach"`
}

type createVolumeRequest struct {
	Volume volumeRequest `json:"volume"`
}

func CreateDisk(projectId, name, volumeType, availabilityZone string, count, size int, multiattach bool) (*evsModels.CreateBatchDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/cloudvolumes", projectId)
	req := volumeRequest{
		Name:             name,
		AvailabilityZone: availabilityZone,
		VolumeType:       volumeType,
		Count:            count,
		Size:             size,
		Multiattach:      multiattach,
	}
	request := createVolumeRequest{
		Volume: req,
	}
	var createdJob evsModels.CreateBatchDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &createdJob, nil)
	return &createdJob, err
}
