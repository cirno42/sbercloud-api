package vpcs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/vpcModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type vpcCreationParameters struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cidr        string `json:"cidr"`
}

type vpcCreationQuery struct {
	Vpc vpcCreationParameters `json:"vpc"`
}

func CreateVpc(projectID, name, description, cidr string) (*vpcModels.VpcModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs", projectID)

	vpc := vpcCreationParameters{
		Name:        name,
		Description: description,
		Cidr:        cidr,
	}
	vpcQuery := vpcCreationQuery{Vpc: vpc}
	var createdVpc vpcModels.VpcEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &vpcQuery, &createdVpc, nil)
	return &createdVpc.Vpc, err
}
