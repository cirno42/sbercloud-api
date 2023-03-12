package dumpModels

import (
	"sbercloud-api/api/models/ecsModels"
	"sbercloud-api/api/models/eipModels"
	"sbercloud-api/api/models/evsModels"
	"sbercloud-api/api/models/natModels"
	"sbercloud-api/api/models/securityGroupModels"
	"sbercloud-api/api/models/subnetModels"
	"sbercloud-api/api/models/vpcModels"
)

type DumpModel struct {
	Vpcs      []vpcModels.VpcModel                     `json:"vpcs"`
	Subnets   []subnetModels.SubnetModel               `json:"subnets"`
	Eips      []eipModels.EipModel                     `json:"eips"`
	Nats      []natModels.NatModel                     `json:"nats"`
	SecGroups []securityGroupModels.SecurityGroupModel `json:"secGroups"`
	SnatRules []natModels.SnatRuleModel                `json:"snatRules"`
	DnatRules []natModels.DnatRuleModel                `json:"dnatRules"`
	ECSs      []ecsModels.ECSModel                     `json:"ECSs"`
	Disks     []evsModels.EvsModel                     `json:"disks"`
	KeyPairs  []ecsModels.Keypair                      `json:"keyPairs"`
}
