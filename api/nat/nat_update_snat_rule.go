package nat

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/natModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type updateSNATRuleParameters struct {
	NatGatewayID    string `json:"nat_gateway_id"`
	PublicIpAddress string `json:"public_ip_address"`
	Description     string `json:"description"`
}

type updateSNATRuleRequest struct {
	SnatRule updateSNATRuleParameters `json:"snat_rule"`
}

type updateSNATRuleResp struct {
	SnatRule natModels.SnatRuleModel `json:"snat_rule"`
}

func UpdateSNATRule(projectID, ruleID, natGatewayID, publicIpAddress, description string) (natModels.SnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/snat_rules/%s", projectID, ruleID)
	params := updateSNATRuleParameters{
		NatGatewayID:    natGatewayID,
		PublicIpAddress: publicIpAddress,
		Description:     description,
	}
	req := updateSNATRuleRequest{SnatRule: params}
	var resp updateSNATRuleResp
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, req, &resp, nil)
	return resp.SnatRule, err
}
