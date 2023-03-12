package securityGroup

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

func DeleteSecurityGroup(projectID, securityGroupID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-groups/%s", projectID, securityGroupID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeleteSecurityGroupByName(projectID, securityGroupName string) error {
	sg, err := GetInfoAboutSecurityGroupByName(projectID, securityGroupName)
	if err != nil {
		return err
	}
	err = DeleteSecurityGroup(projectID, sg.Id)
	return err
}
