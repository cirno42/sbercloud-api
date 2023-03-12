package subnets

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/nat"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

func DeleteSubnet(projectID, vpcID, subnetID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s/subnets/%s", projectID, vpcID, subnetID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeleteSubnetRecursive(projectID, vpcID, subnetID string) error {
	nats, err := nat.GetNatList(projectID)
	for i := 0; i < len(nats); i++ {
		if (nats[i].InternalNetworkID == subnetID) && (nats[i].RouterID == vpcID) {
			err = nat.DeleteNat(projectID, nats[i].ID)
			if err != nil {
				return err
			}
		}
	}
	err = DeleteSubnet(projectID, vpcID, subnetID)
	return err
}
