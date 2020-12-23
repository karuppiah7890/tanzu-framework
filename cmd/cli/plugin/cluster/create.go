package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/vmware-tanzu-private/core/apis/client/v1alpha1"
	"github.com/vmware-tanzu-private/core/pkg/v1/client"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/constants"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/tkgctl"
)

// Note: We can remove all this additional options at the time when
// we decide we really want to use `--file` option to create cluster
// and we don't have to support all these flags with create cluster command
type createClusterOptions struct {
	plan                        string
	infrastructureProvider      string
	namespace                   string
	kubernetesVersion           string
	controlPlaneMachineCount    int
	workerMachineCount          int
	timeout                     time.Duration
	generateOnly                bool
	size                        string
	controlPlaneSize            string
	workerSize                  string
	unattended                  bool
	cniType                     string
	enableClusterOptions        string
	vsphereControlPlaneEndpoint string
	clusterConfigFile           string
}

var cc = &createClusterOptions{}

var createClusterCmd = &cobra.Command{
	Use:   "create CLUSTER_NAME",
	Short: "Create a cluster",
	RunE:  create,
}

func init() {
	createClusterCmd.Flags().StringVarP(&cc.clusterConfigFile, "file", "f", "", "Cluster configuration file from which to create a Cluster")

	createClusterCmd.Flags().StringVarP(&cc.plan, "plan", "p", "", "The plan to be used for creating the workload cluster")
	createClusterCmd.Flags().StringVarP(&cc.kubernetesVersion, "kubernetes-version", "k", "", fmt.Sprintf("The kubernetes version to use for the workload cluster"))
	createClusterCmd.Flags().IntVarP(&cc.controlPlaneMachineCount, "controlplane-machine-count", "c", 0, "The number of control plane machines to be added to the workload cluster (default 1 or 3 depending on dev or prod plan)")
	createClusterCmd.Flags().IntVarP(&cc.workerMachineCount, "worker-machine-count", "w", 0, "The number of worker machines to be added to the workload cluster (default 1 or 3 depending on dev or prod plan)")
	createClusterCmd.Flags().BoolVarP(&cc.generateOnly, "dry-run", "d", false, "Does not create cluster but show the deployment YAML instead")
	createClusterCmd.Flags().StringVarP(&cc.namespace, "namespace", "n", "", "The namespace where the cluster should be deployed. Assumes 'default' if not specified")
	createClusterCmd.Flags().StringVarP(&cc.vsphereControlPlaneEndpoint, "vsphere-controlplane-endpoint", "", "", "Virtual IP address or FQDN for the cluster's control plane nodes")
	createClusterCmd.Flags().DurationVarP(&cc.timeout, "timeout", "t", constants.DefaultLongRunningOperationTimeout, "Time duration to wait for an operation before timeout. Timeout duration in hours(h)/minutes(m)/seconds(s) units or as some combination of them (e.g. 2h, 30m, 2h30m10s)")
	createClusterCmd.Flags().StringVarP(&cc.cniType, "cni", "", "", "Specify the CNI provider the workload cluster should use ['antrea' (default), 'calico', 'none'].")
	createClusterCmd.Flags().StringVarP(&cc.size, "size", "", "", "Specify size for all nodes including control plane and worker nodes. It can be overridden by --controlplane-size and --worker-size options. (See [+])")
	createClusterCmd.Flags().StringVarP(&cc.controlPlaneSize, "controlplane-size", "", "", "Specify size for the control plane node. (See [+])")
	createClusterCmd.Flags().StringVarP(&cc.workerSize, "worker-size", "", "", "Specify size of the worker node. (See [+])")
	createClusterCmd.Flags().BoolVarP(&cc.unattended, "yes", "y", false, "Create workload cluster without asking for confirmation")
	createClusterCmd.Flags().StringVarP(&cc.enableClusterOptions, "enable-cluster-options", "", "", "List of comma separated cluster options to be enabled")
	createClusterCmd.Flags().StringVarP(&cc.infrastructureProvider, "infrastructure", "i", "", "The target infrastructure on which to deploy the workload cluster.")

	// Hide some of the variables not relevant to tanzu cli at the moment
	createClusterCmd.Flags().MarkHidden("plan")                          //nolint
	createClusterCmd.Flags().MarkHidden("kubernetes-version")            //nolint
	createClusterCmd.Flags().MarkHidden("controlplane-machine-count")    //nolint
	createClusterCmd.Flags().MarkHidden("worker-machine-count")          //nolint
	createClusterCmd.Flags().MarkHidden("namespace")                     //nolint
	createClusterCmd.Flags().MarkHidden("vsphere-controlplane-endpoint") //nolint
	createClusterCmd.Flags().MarkHidden("timeout")                       //nolint
	createClusterCmd.Flags().MarkHidden("cni")                           //nolint
	createClusterCmd.Flags().MarkHidden("size")                          //nolint
	createClusterCmd.Flags().MarkHidden("controlplane-size")             //nolint
	createClusterCmd.Flags().MarkHidden("worker-size")                   //nolint
	createClusterCmd.Flags().MarkHidden("yes")                           //nolint
	createClusterCmd.Flags().MarkHidden("enable-cluster-options")        //nolint
	createClusterCmd.Flags().MarkHidden("infrastructure")                //nolint // Usually not needed as they are implied from configuration of the management cluster.
}

func create(cmd *cobra.Command, args []string) error {
	server, err := client.GetCurrentServer()
	if err != nil {
		return err
	}
	clusterName := ""
	if len(args) > 0 {
		clusterName = args[0]
	}

	if server.IsGlobal() {
		return errors.New("creating cluster with a global server is not implemented yet")
	}
	return createCluster(clusterName, server)
}

func createCluster(clusterName string, server *v1alpha1.Server) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	tkgctlClient, err := tkgctl.New(tkgctl.Options{
		ConfigDir:   configDir,
		KubeConfig:  server.ManagementClusterOpts.Path,
		KubeContext: server.ManagementClusterOpts.Context,
	})
	if err != nil {
		return err
	}

	ccOptions := tkgctl.CreateClusterOptions{
		ClusterConfigFile:           cc.clusterConfigFile,
		ClusterName:                 clusterName,
		Namespace:                   cc.namespace,
		Plan:                        cc.plan,
		InfrastructureProvider:      cc.infrastructureProvider,
		KubernetesVersion:           cc.kubernetesVersion,
		ControlPlaneMachineCount:    cc.controlPlaneMachineCount,
		WorkerMachineCount:          cc.workerMachineCount,
		GenerateOnly:                cc.generateOnly,
		Size:                        cc.size,
		ControlPlaneSize:            cc.controlPlaneSize,
		WorkerSize:                  cc.workerSize,
		CniType:                     cc.cniType,
		EnableClusterOptions:        cc.enableClusterOptions,
		VsphereControlPlaneEndpoint: cc.vsphereControlPlaneEndpoint,
		SkipPrompt:                  cc.unattended,
		Timeout:                     cc.timeout,
	}

	return tkgctlClient.CreateCluster(ccOptions)
}
