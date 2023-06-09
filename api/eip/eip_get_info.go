package eip

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/ecs"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/iam"
	"github.com/cirno42/sbercloud-api/api/models/eipModels"
	"github.com/cirno42/sbercloud-api/api/models/iamModels"
	"github.com/cirno42/sbercloud-api/api/nat"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
	"sync"
)

type eipQueryingResponse struct {
	PublicIPs []eipModels.EipModel `json:"publicips"`
}

type eipGetInfoResponse struct {
	PublicIP eipModels.EipModel `json:"publicip"`
}

func GetInfoAboutEIPByAddress(projectID, ipAddress string) (*eipModels.EipModel, error) {
	ipArray, err := GetEIPsList(projectID, 1000, "")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ipArray); i++ {
		if ipArray[i].PublicIPAddress == ipAddress {
			return &ipArray[i], nil
		}
	}
	return nil, nil
}

func GetEIPInfo(projectID, publicIpId string) (*eipModels.EipModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips/%s", projectID, publicIpId)
	var publicIP eipGetInfoResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &publicIP, nil)
	return &publicIP.PublicIP, err
}

func GetEIPsList(projectID string, limit int, marker string) ([]eipModels.EipModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips?limit=%d&marker=%s", projectID, limit, marker)
	var ipsArray eipQueryingResponse
	err := requestMakers.CreateAndDoRequestInSpecifiedProject(endpoint, projectID, requestMakers.HTTP_METHOD_GET, nil, &ipsArray)
	return ipsArray.PublicIPs, err
}

func GetActiveIPsInSpecifiedProject(projectID string) ([]eipModels.ActiveEIP, error) {
	eips, err := GetEIPsList(projectID, 1000, "")
	activeIPs := make([]eipModels.ActiveEIP, 0)
	if err != nil {
		return nil, err
	}
	servers, err := ecs.GetECSList(projectID, 0, 0, "")
	for i := 0; i < len(servers); i++ {
		for _, network := range servers[i].Addresses {
			for _, net := range network {
				if net.OSEXTIPSType == "floating" {
					var eipID string
					for _, eip := range eips {
						if eip.PublicIPAddress == net.Addr {
							eipID = eip.ID
						}
					}
					ip := eipModels.ActiveEIP{
						ID:           eipID,
						Address:      net.Addr,
						InstanceID:   servers[i].ID,
						InstanceType: "ECS",
					}
					activeIPs = append(activeIPs, ip)
				}
			}
		}
	}
	snatRules, err := nat.ListSNATRules(projectID, "", "", 0)
	if err != nil {
		return nil, err
	}
	for _, rule := range snatRules {
		ip := eipModels.ActiveEIP{
			ID:                 rule.FloatingIPID,
			Address:            rule.FloatingIPAddress,
			InstanceID:         rule.ID,
			InstanceType:       "SNAT-Rule",
			ParentInstanceID:   rule.NatGatewayID,
			ParentInstanceType: "NAT",
		}
		activeIPs = append(activeIPs, ip)
	}
	dnatRules, err := nat.ListDNATRules(projectID, "", "", "", "", "", "", "", 0, 0)
	if err != nil {
		return nil, err
	}
	for _, rule := range dnatRules {
		ip := eipModels.ActiveEIP{
			ID:                 rule.FloatingIPID,
			Address:            rule.FloatingIPAddress,
			InstanceID:         rule.ID,
			InstanceType:       "DNAT-Rule",
			ParentInstanceID:   rule.NatGatewayID,
			ParentInstanceType: "NAT",
		}
		activeIPs = append(activeIPs, ip)
	}
	return activeIPs, err
}

func GetActiveIPsInAllProjects() ([]eipModels.ProjectActiveEIP, error) {
	projects, err := iam.ListProjects()
	if err != nil {
		return nil, err
	}
	activeIPs := make([]eipModels.ProjectActiveEIP, len(projects))
	var wg sync.WaitGroup
	for i, project := range projects {
		wg.Add(1)
		go func(activeIP *eipModels.ProjectActiveEIP, p iamModels.ProjectModel) {
			defer wg.Done()
			ips, _ := GetActiveIPsInSpecifiedProject(p.ID)
			*activeIP = eipModels.ProjectActiveEIP{
				ProjectID:   p.ID,
				ProjectName: p.Name,
				ActiveIP:    ips,
			}
		}(&activeIPs[i], project)
	}
	wg.Wait()
	return activeIPs, err
}
