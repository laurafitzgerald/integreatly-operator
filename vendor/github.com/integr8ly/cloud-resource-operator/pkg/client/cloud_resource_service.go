package client

// cloud resource service exposes functionality for creating redis, postgres resources inside a cloud
//


import (
	"context"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/integr8ly/cloud-resource-operator/pkg/providers/aws"
)

type CloudResourceSpec struct {
	Ctx context.Context
	Client client.Client
}

type CloudResourceService struct {
	Ctx context.Context
	Client client.Client
	CredentialManager aws.CredentialManager
	Logger *logrus.Entry
}

var _ CloudResourceOperatorService = (*CloudResourceService)(nil)

func NewCloudResourceService(spec *CloudResourceSpec) *CloudResourceService {
	return &CloudResourceService{
		Ctx: spec.Ctx,
		Client: spec.Client,
		CredentialManager: aws.NewCredentialMinterCredentialManager(spec.Client),
	}
}

type CloudResourceOperatorService interface {
	ReconcileStrategyMaps(timeConfig *StrategyTimeConfig, tier, namespace string) error
	GetNonOverlappingCIDR(namespace string)  (string, error)
	ReconcilePostgres(ctx context.Context, client client.Client, productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.Postgres, error)
	ReconcileRedis(ctx context.Context, client client.Client, productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.Redis, error)
	ReconcileBlobStorage(productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.BlobStorage, error)
}
