package dumpModels

import (
	"github.com/cirno42/sbercloud-api/api/models/ecsModels"
	"github.com/cirno42/sbercloud-api/api/models/eipModels"
	"github.com/cirno42/sbercloud-api/api/models/evsModels"
	"github.com/cirno42/sbercloud-api/api/models/natModels"
	"github.com/cirno42/sbercloud-api/api/models/securityGroupModels"
	"github.com/cirno42/sbercloud-api/api/models/subnetModels"
	"github.com/cirno42/sbercloud-api/api/models/vpcModels"
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
