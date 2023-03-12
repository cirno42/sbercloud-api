package vpcs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/subnets"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
	"time"
)

func DeleteVpc(projectID, vpcID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s", projectID, vpcID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeleteVpcRecursive(projectID, vpcID string) error {
	subnetList, err := subnets.GetSubnetsList(projectID, 0, "", vpcID)
	if err != nil {
		return err
	}
	for i := 0; i < len(subnetList); i++ {
		err := subnets.DeleteSubnetRecursive(projectID, vpcID, subnetList[i].Id)
		if err != nil {
			return err
		}
	}
	time.Sleep(1 * time.Second)
	err = DeleteVpc(projectID, vpcID)
	return err
}
