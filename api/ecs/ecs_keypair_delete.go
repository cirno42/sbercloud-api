package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

func DeleteKeypair(projectID, keypairName string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-keypairs/%s", projectID, keypairName)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}
