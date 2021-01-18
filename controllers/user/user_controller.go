package controllers

import (
	"context"
	"time"

	l "github.com/integr8ly/integreatly-operator/pkg/resources/logger"

	userHelper "github.com/integr8ly/integreatly-operator/pkg/resources/user"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	usersv1 "github.com/openshift/api/user/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var log = l.NewLoggerWithContext(l.Fields{l.ControllerLogContext: "subscription_controller"})

// UserReconciler reconciles a User object
type UserReconciler struct {
	// Log logr.Logger
}

// +kubebuilder:rbac:groups=user.openshift.io,resources=users,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=user.openshift.io,resources=users/status,verbs=get;update;patch

func (r *UserReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	// reqLogger := r.Log.WithValues("user", request.NamespacedName)
	log.Logger = log.WithContext(l.Fields{l.ControllerLogContext: "user_controller"})
	log.Info("Reconciling User")

	// new client to avoid caching issues
	restConfig := controllerruntime.GetConfigOrDie()
	restConfig.Timeout = time.Second * 10
	c, _ := k8sclient.New(restConfig, k8sclient.Options{})
	ctx := context.TODO()

	rhmiGroup := &usersv1.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rhmi-developers",
		},
	}

	or, err := controllerutil.CreateOrUpdate(ctx, c, rhmiGroup, func() error {
		users := &usersv1.UserList{}
		err := c.List(ctx, users)
		if err != nil {
			return err
		}

		groups := &usersv1.GroupList{}
		err = c.List(ctx, groups)
		if err != nil {
			return err
		}

		rhmiGroup.Users = mapUserNames(users, groups)

		return nil
	})
	log.Infof("Operation Result", l.Fields{"groupName": rhmiGroup.Name, "result": string(or)})

	return ctrl.Result{}, err
}

func mapUserNames(users *usersv1.UserList, groups *usersv1.GroupList) []string {
	var result = []string{}
	for _, user := range users.Items {
		// Certain users such as sre do not need to be added
		if !userHelper.UserInExclusionGroup(user, groups) {
			result = append(result, user.Name)
		}
	}

	return result
}

func (r *UserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&usersv1.User{}).
		Complete(r)
}