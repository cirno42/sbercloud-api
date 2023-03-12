package nat

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/natModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type snatCreateResponse struct {
	SnatRule natModels.SnatRuleModel `json:"snat_rule"`
}

type snatCreateRequest struct {
	SnatRule snatCreateParameters `json:"snat_rule"`
}

type snatCreateParameters struct {
	NatGatewayID string `json:"nat_gateway_id"`
	NetworkID    string `json:"network_id"`
	SourceType   int    `json:"source_type"`
	FloatingIPID string `json:"floating_ip_id"`
	Description  string `json:"description"`
}

func CreateSNATRule(projectID, natID, vpcID, eipID, description string, sourceType int) (natModels.SnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/snat_rules", projectID)
	params := snatCreateParameters{
		NatGatewayID: natID,
		NetworkID:    vpcID,
		SourceType:   sourceType,
		FloatingIPID: eipID,
		Description:  description,
	}
	req := snatCreateRequest{SnatRule: params}
	var createdRule snatCreateResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, &createdRule, nil)
	return createdRule.SnatRule, err

}
