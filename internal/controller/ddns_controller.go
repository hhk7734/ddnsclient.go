package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkingv1alpha1 "github.com/hhk7734/ddnsclient.go/api/v1alpha1"
	"github.com/hhk7734/ddnsclient.go/internal/dynamicip"
)

// DDNSReconciler reconciles a DDNS object
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
func (r *DDNSReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DDNSReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1alpha1.DDNS{}).
		Named("ddns").
		Complete(r)
}
