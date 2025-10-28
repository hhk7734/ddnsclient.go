package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/hhk7734/ddnsclient.go/internal/dynamicip"
	networkingv1alpha1 "github.com/hhk7734/ddnsclient.go/pkg/apis/networking.loliot.net/v1alpha1"
)

// DDNSReconciler reconciles a DDNS object.
type DDNSReconciler struct {
	client.Client

	Scheme *runtime.Scheme
	IPer   dynamicip.IPer
}

// +kubebuilder:rbac:groups=networking.loliot.net,resources=ddns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.loliot.net,resources=ddns/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=networking.loliot.net,resources=ddns/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *DDNSReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DDNSReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.IPer == nil {
		r.IPer = dynamicip.NewAWSIPer()
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1alpha1.DDNS{}).
		Complete(r)
}
