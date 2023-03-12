package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type addSecurityGroupParameters struct {
	Name string `json:"name"`
}

type addSecurityGroupRequest struct {
	AddSecurityGroup addSecurityGroupParameters `json:"addSecurityGroup"`
}

func AddSecurityGroup(projectID, ecsID, sgName string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/servers/%s/action", projectID, ecsID)
	params := addSecurityGroupParameters{Name: sgName}
	req := addSecurityGroupRequest{AddSecurityGroup: params}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, nil, nil)
	return err
}
