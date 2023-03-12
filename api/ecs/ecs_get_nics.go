package ecs

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type EcsNic struct {
	InterfaceAttachments []struct {
		PortState string `json:"port_state"`
		FixedIps  []struct {
			SubnetID  string `json:"subnet_id"`
			IPAddress string `json:"ip_address"`
		} `json:"fixed_ips"`
		NetID   string `json:"net_id"`
		PortID  string `json:"port_id"`
		MacAddr string `json:"mac_addr"`
	} `json:"interfaceAttachments"`
}

func GetEcsNics(projectID, serverID string) (EcsNic, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/os-interface", projectID, serverID)
	var nics EcsNic
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &nics, nil)
	return nics, err
}
