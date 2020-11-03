package marin3r

import (
	"fmt"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) newSoftLimitAlertsReconciler() resources.AlertReconciler {
	softDailyLimits := r.RateLimitConfig.SoftDailyLimits
	if softDailyLimits == nil || len(softDailyLimits) == 0 {
		return &resources.NoopAlertReconciler{}
	}

	return &resources.AlertReconcilerImpl{
		ProductName:  "3scale",
		Installation: r.installation,
		Logger:       r.logger,
		Alerts: []resources.AlertConfiguration{
			{
				AlertName: "rate-limit-soft-limits",
				GroupName: "soft-limits.rules",
				Namespace: r.Config.GetNamespace(),
				Rules:     mapSoftLimitRules(softDailyLimits),
			},
		},
	}
}

func mapSoftLimitRules(softDailyLimits []uint32) []monitoringv1.Rule {
	result := make([]monitoringv1.Rule, len(softDailyLimits))

	for i, softDailyLimit := range softDailyLimits {
		result[i] = monitoringv1.Rule{
			Alert: fmt.Sprintf("RHOAMApiUsageSoftLimitReachedTier%d", i+1),
			Expr: intstr.FromString(
				fmt.Sprintf("(ratelimit_service_rate_limit_apicast_ratelimit_generic_key_slowpath_total_hits - ratelimit_service_rate_limit_apicast_ratelimit_generic_key_slowpath_total_hits offset 1d) > %d", softDailyLimit),
			),
			Annotations: map[string]string{
				"message": fmt.Sprintf("soft daily limit of requests reached (%d)", softDailyLimit),
			},
			Labels: map[string]string{
				"severity": "warning",
			},
		}
	}

	return result
}
