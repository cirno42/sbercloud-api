package nat

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/natModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type NatCreationModel struct {
	NatGateway NatCreationParameters `json:"nat_gateway"`
}

type NatCreationParameters struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	RouterID            string `json:"router_id"`
	InternalNetworkID   string `json:"internal_network_id"`
	Spec                string `json:"spec"`
	EnterpriseProjectID string `json:"enterprise_project_id"`
}

type natCreateResponse struct {
	NatGateway natModels.NatModel `json:"nat_gateway"`
}

func CreateNAT(projectID, name, desc, routerID, internalNetworkID, spec, enterpriseProjectID string) (*natModels.NatModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways", projectID)
	params := NatCreationParameters{
		Name:                name,
		Description:         desc,
		RouterID:            routerID,
		InternalNetworkID:   internalNetworkID,
		Spec:                spec,
		EnterpriseProjectID: enterpriseProjectID,
	}
	requestBody := NatCreationModel{
		NatGateway: params,
	}
	var createdNat natCreateResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &requestBody, &createdNat, nil)
	return &createdNat.NatGateway, err
}
