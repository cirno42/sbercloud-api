package nat

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/internal/handlers/requestMakers"
)

func DeleteSNATRule(projectID, natID, ruleID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways/%s/snat_rules/%s", projectID, natID, ruleID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}