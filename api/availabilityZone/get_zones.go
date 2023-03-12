package availabilityZone

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/availabilityZoneModels"
	"sbercloud-api/internal/handlers/requestMakers"
)

type AvailabilityZoneInfoResponse struct {
	Zones []availabilityZoneModels.AvailabilityZoneInfo `json:"availabilityZoneInfo"`
}

func GetZonesList(projectID string) ([]availabilityZoneModels.AvailabilityZoneInfo, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-availability-zone", projectID)
	var zones AvailabilityZoneInfoResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &zones, nil)
	return zones.Zones, err
}
