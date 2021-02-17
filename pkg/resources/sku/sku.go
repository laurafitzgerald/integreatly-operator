package sku

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type SKU struct {
	name           string
	productConfigs map[v1alpha1.ProductName]ProductConfig
}

type ProductConfig struct {
	productName     v1alpha1.ProductName
	resourceConfigs map[string]ResourceConfig
}

type ResourceConfig struct {
	Replicas  int                         `json:"replicas,omitempty"`
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

type RateLimit struct {
	Unit            string  `json:"unit,omitempty"`
	RequestsPerUnit int64   `json:"requests_per_unit,omitempty"`
	AlertLimits     []int64 `json:"alert_limits,omitempty"`
}

type skuConfigReceiver struct {
	Name      string                    `json:"name,omitempty"`
	RateLimit RateLimit                 `json:"rate-limiting,omitempty"`
	Resources map[string]ResourceConfig `json:"resources,omitempty"`
}

func GetSKU(SKUId string, SKUConfig *corev1.ConfigMap) (*SKU, error) {
	received := &[]skuConfigReceiver{}
	err := json.Unmarshal([]byte(SKUConfig.Data["sku-configs"]), received)
	if err != nil {
		return nil, err
	}

	skuReceiver := skuConfigReceiver{}
	for _, sku := range *received {
		if sku.Name == SKUId {
			skuReceiver = sku
			break
		}
	}
	if skuReceiver.Name == "" {
		return nil, errors.New(fmt.Sprintf("could not find sku config with name '%s'", SKUId))
	}

	// map of products iterate over that to build the return map
	products := map[v1alpha1.ProductName][]string{
		v1alpha1.Product3Scale: {
			"backend_listener",
			"backend_worker",
			"apicast_production",
		},
		v1alpha1.ProductRHSSOUser: {
			"keycloak",
		},
		v1alpha1.ProductMarin3r: {},
	}
	retSku := &SKU{
		name:           SKUId,
		productConfigs: map[v1alpha1.ProductName]ProductConfig{},
	}
	// loop through array of ddcss (deployment deploymentConfig StatefulSets)
	for product, ddcssNames := range products {
		pc := ProductConfig{
			productName:     product,
			resourceConfigs: map[string]ResourceConfig{},
		}
		for _, ddcssName := range ddcssNames {
			pc.resourceConfigs[ddcssName] = skuReceiver.Resources[ddcssName]
		}
		retSku.productConfigs[product] = pc
	}

	return retSku, nil
}

func (s *SKU) GetProduct(productName v1alpha1.ProductName) ProductConfig {
	return s.productConfigs[productName]
}

func (pc *ProductConfig) GetResourceConfig(ddcssName string) corev1.ResourceRequirements {
	return pc.resourceConfigs[ddcssName].Resources
}

func (pc ProductConfig) GetReplicas(ddcssName string) int {
	return pc.resourceConfigs[ddcssName].Replicas
}