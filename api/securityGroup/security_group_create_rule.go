package securityGroup

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/securityGroupModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type securityGroupRuleResponse struct {
	SecurityGroupRule securityGroupModels.SecurityGroupRuleModel `json:"security_group_rule"`
}

type securityGroupRuleCreateRequest struct {
	SecurityGroupRule securityGroupRuleCreateParameters `json:"security_group_rule"`
}

type securityGroupRuleCreateParameters struct {
	SecurityGroupId string `json:"security_group_id"`
	Description     string `json:"description"`
	Direction       string `json:"direction"`
	EtherType       string `json:"ethertype"`
	Protocol        string `json:"protocol"`
	PortRangeMin    *int   `json:"port_range_min"`
	PortRangeMax    *int   `json:"port_range_max"`
	RemoteIpPrefix  string `json:"remote_ip_prefix"`
	RemoteGroupId   string `json:"--"`
}

func CreateSecurityGroupRule(projectID, securityGroupId, description, direction,
	etherType, protocol string, portRangeMin, portRangeMax int, remoteIpPrefix, remoteGroupId string) (*securityGroupModels.SecurityGroupRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-group-rules", projectID)
	var portMin *int
	if portRangeMin != -1 {
		portMin = &portRangeMin
	}
	var portMax *int
	if portRangeMax != -1 {
		portMax = &portRangeMax
	}
	sgRuleParams := securityGroupRuleCreateParameters{
		SecurityGroupId: securityGroupId,
		Description:     description,
		Direction:       direction,
		EtherType:       etherType,
		Protocol:        protocol,
		PortRangeMin:    portMin,
		PortRangeMax:    portMax,
		RemoteIpPrefix:  remoteIpPrefix,
		RemoteGroupId:   remoteGroupId,
	}
	sgRuleRequest := securityGroupRuleCreateRequest{
		SecurityGroupRule: sgRuleParams,
	}
	var createdSGRule securityGroupRuleResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &sgRuleRequest, &createdSGRule, nil)
	return &createdSGRule.SecurityGroupRule, err
}
