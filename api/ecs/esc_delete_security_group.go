package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type removeSecurityGroupParameters struct {
	Name string `json:"name"`
}

type removeSecurityGroupRequest struct {
	RemoveSecurityGroup removeSecurityGroupParameters `json:"removeSecurityGroup"`
}

func RemoveSecurityGroup(projectID, ecsID, sgName string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/servers/%s/action", projectID, ecsID)
	params := removeSecurityGroupParameters{Name: sgName}
	req := removeSecurityGroupRequest{RemoveSecurityGroup: params}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, nil, nil)
	return err
}
