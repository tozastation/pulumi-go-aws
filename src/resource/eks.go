package resource

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func NewEKS(ctx *pulumi.Context, vpcID pulumi.IDOutput) error {
	const (
		clusterName                 = "tozastation-dev"
		clusterVersion              = "1.17"
		enableEndpointPrivateAccess = true
		enableEndpointPublicAccess  = true
	)
	var (
		clusterArgs = &eks.ClusterArgs{
			EnabledClusterLogTypes: pulumi.StringArray{
				pulumi.String("api"),
				pulumi.String("audit"),
				pulumi.String("authenticator"),
				pulumi.String("controllerManager"),
				pulumi.String("scheduler"),
			},
			EncryptionConfig: nil,
			Name:             pulumi.StringPtr(clusterName),
			RoleArn:          nil,
			Tags:             RequiredTags,
			Version:          pulumi.StringPtr(clusterVersion),
			VpcConfig: eks.ClusterVpcConfigArgs{
				ClusterSecurityGroupId: nil,
				EndpointPrivateAccess:  pulumi.BoolPtr(enableEndpointPrivateAccess),
				EndpointPublicAccess:   pulumi.BoolPtr(enableEndpointPublicAccess),
				PublicAccessCidrs:      pulumi.StringArray{},
				SecurityGroupIds:       pulumi.StringArray{},
				SubnetIds:              pulumi.StringArray{},
				VpcId:                  vpcID,
			},
		}
	)
	eksCluter, err := eks.NewCluster(ctx, clusterName, clusterArgs)
	if err != nil {
		return err
	}
	return nil
}
