package subnets

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/subnetModels"
	"sbercloud-api/internal/handlers/requestMakers"
)

func UpdateSubnet(projectID, name, description string, ipv6Enable, dhcpEnable bool,
	primaryDns, secondaryDns, vpcId, subnetId string) (*subnetModels.SubnetEntity, error) {

	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s/subnets/%s", projectID, vpcId, subnetId)
	subnetParameters := subnetCreationParameters{
		Name:         name,
		Description:  description,
		Ipv6Enable:   ipv6Enable,
		DhcpEnable:   dhcpEnable,
		PrimaryDns:   primaryDns,
		SecondaryDns: secondaryDns,
	}
	subnetQuery := subnetCreationQuery{Subnet: subnetParameters}
	var createdSubnet subnetModels.SubnetEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, &subnetQuery, &createdSubnet, nil)
	return &createdSubnet, err
}
