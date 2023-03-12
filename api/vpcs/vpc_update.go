package vpcs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/vpcModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

func UpdateVpc(projectID, id, name, description, cidr string) (*vpcModels.VpcEntity, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s", projectID, id)

	vpc := vpcCreationParameters{
		Name:        name,
		Description: description,
		Cidr:        cidr,
	}
	vpcQuery := vpcCreationQuery{Vpc: vpc}
	var createdVpc vpcModels.VpcEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, &vpcQuery, &createdVpc, nil)
	return &createdVpc, err

}
