package main

import (
	"os"

	"github.com/go-logr/zapr"
	"github.com/hhk7734/ddnsclient.go/internal/manager/controller"
	"github.com/hhk7734/ddnsclient.go/internal/pkg/logger"
	networkingv1alpha1 "github.com/hhk7734/ddnsclient.go/pkg/apis/networking.loliot.net/v1alpha1"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	_ "github.com/joho/godotenv/autoload"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(networkingv1alpha1.AddToScheme(scheme))
}

func main() {
	// env
	viper.AutomaticEnv()

	// flag
	pflag.CommandLine.AddFlagSet(logger.LogPFlags())
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	logger.SetGlobalZapLogger(logger.LogConfigFromViper())

	ctrl.SetLogger(zapr.NewLogger(zap.L()))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Metrics: server.Options{
			BindAddress: ":8080",
		},
		HealthProbeBindAddress: ":8081",
	})
	if err != nil {
		os.Exit(1)
	}

	if err = (&controller.DDNSReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
