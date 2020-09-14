package resource

import "github.com/pulumi/pulumi/sdk/v2/go/pulumi"

var (
	serviceName = "tozastation"
	env = "dev"
)

var RequiredTags = pulumi.StringMap{
	"service": pulumi.String(serviceName),
	"env": pulumi.String(env),
}

var LocalName = serviceName + "-" + env