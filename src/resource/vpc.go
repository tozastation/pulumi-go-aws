package resource

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func NewVPC(ctx *pulumi.Context) (*ec2.Vpc, error) {
	const (
		cidrBlock = "10.10.0.0/16"
	)
	var (
		vpcName = LocalName + "-vpc"
		vpcArgs = &ec2.VpcArgs{
			AssignGeneratedIpv6CidrBlock: pulumi.BoolPtr(false),
			CidrBlock:                    pulumi.String(cidrBlock),
			EnableClassiclink:            pulumi.BoolPtr(false),
			EnableClassiclinkDnsSupport:  pulumi.BoolPtr(false),
			EnableDnsHostnames:           pulumi.BoolPtr(true),
			EnableDnsSupport:             pulumi.BoolPtr(true),
			InstanceTenancy:              pulumi.String(""),
			Tags:                         RequiredTags,
		}
	)
	vpc, err := ec2.NewVpc(
		ctx,
		vpcName,
		vpcArgs,
	)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}


