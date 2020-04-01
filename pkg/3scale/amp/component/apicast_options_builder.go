package component

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ApicastOptions struct {
	// required options
	appLabel       string
	managementAPI  string
	openSSLVerify  string
	responseCodes  string
	tenantName     string
	wildcardDomain string

	// non required options
	resourceRequirements        *v1.ResourceRequirements
	stagingResourceRequirements *v1.ResourceRequirements
	replicas                    *int32
	namespace                   *string
	environment                 *string
}

type ApicastOptionsBuilder struct {
	options ApicastOptions
}

func (a *ApicastOptionsBuilder) AppLabel(appLabel string) {
	a.options.appLabel = appLabel
}

func (a *ApicastOptionsBuilder) ManagementAPI(managementAPI string) {
	a.options.managementAPI = managementAPI
}

func (a *ApicastOptionsBuilder) OpenSSLVerify(openSSLVerify string) {
	a.options.openSSLVerify = openSSLVerify
}

func (a *ApicastOptionsBuilder) ResponseCodes(responseCodes string) {
	a.options.responseCodes = responseCodes
}

func (a *ApicastOptionsBuilder) TenantName(tenantName string) {
	a.options.tenantName = tenantName
}

func (a *ApicastOptionsBuilder) WildcardDomain(wildcardDomain string) {
	a.options.wildcardDomain = wildcardDomain
}

func (a *ApicastOptionsBuilder) ResourceRequirements(resourceRequirements v1.ResourceRequirements) {
	a.options.resourceRequirements = &resourceRequirements
}

// func (a *ApicastOptionsBuilder) ProductionResourceRequirements(resourceRequirements v1.ResourceRequirements) {
// 	a.options.productionResourceRequirements = &resourceRequirements
// }

// func (a *ApicastOptionsBuilder) StagingResourceRequirements(resourceRequirements v1.ResourceRequirements) {
// 	a.options.stagingResourceRequirements = &resourceRequirements
// }

func (a *ApicastOptionsBuilder) Replicas(replicas int32) {
	a.options.replicas = &replicas
}

func (a *ApicastOptionsBuilder) Namespace(namespace string) {
	a.options.namespace = &namespace
}

func (a *ApicastOptionsBuilder) Environment(environment string) {
	a.options.environment = &environment
}

// func (a *ApicastOptionsBuilder) StagingReplicas(replicas int32) {
// 	a.options.stagingReplicas = &replicas
// }

// func (a *ApicastOptionsBuilder) StagingNamespace(namespace string) {
// 	a.options.stagingNamespace = namespace
// }

// func (a *ApicastOptionsBuilder) ProductionReplicas(replicas int32) {
// 	a.options.productionReplicas = &replicas
// }

// func (a *ApicastOptionsBuilder) ProductionNamespace(namespace string) {
// 	a.options.productionNamespace = namespace
// }

func (a *ApicastOptionsBuilder) Build() (*ApicastOptions, error) {
	err := a.setRequiredOptions()
	if err != nil {
		return nil, err
	}

	a.setNonRequiredOptions()

	return &a.options, nil
}

func (a *ApicastOptionsBuilder) setRequiredOptions() error {
	if a.options.appLabel == "" {
		return fmt.Errorf("no AppLabel has been provided")
	}
	if a.options.managementAPI == "" {
		return fmt.Errorf("no management API has been provided")
	}
	if a.options.openSSLVerify == "" {
		return fmt.Errorf("no OpenSSLVerify option has been provided")
	}
	if a.options.responseCodes == "" {
		return fmt.Errorf("no response codes have been provided")
	}
	if a.options.tenantName == "" {
		return fmt.Errorf("no tenant name has been provided")
	}
	if a.options.wildcardDomain == "" {
		return fmt.Errorf("no wildcard domain has been provided")
	}

	return nil
}

func (a *ApicastOptionsBuilder) setNonRequiredOptions() {
	if a.options.resourceRequirements == nil {
		a.options.resourceRequirements = a.defaultStagingResourceRequirements()
	}

	// if a.options.stagingResourceRequirements == nil {
	// 	a.options.stagingResourceRequirements = a.defaultStagingResourceRequirements()
	// }

	if a.options.replicas == nil {
		var defaultReplicas int32 = 1
		a.options.replicas = &defaultReplicas
	}

	// if a.options.productionReplicas == nil {
	// 	var defaultProductionReplicas int32 = 1
	// 	a.options.productionReplicas = &defaultProductionReplicas
	// }
}

func (a *ApicastOptionsBuilder) defaultProductionResourceRequirements() *v1.ResourceRequirements {
	return &v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse("1000m"),
			v1.ResourceMemory: resource.MustParse("128Mi"),
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse("500m"),
			v1.ResourceMemory: resource.MustParse("64Mi"),
		},
	}
}

func (a *ApicastOptionsBuilder) defaultStagingResourceRequirements() *v1.ResourceRequirements {
	return &v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse("100m"),
			v1.ResourceMemory: resource.MustParse("128Mi"),
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse("50m"),
			v1.ResourceMemory: resource.MustParse("64Mi"),
		},
	}
}
