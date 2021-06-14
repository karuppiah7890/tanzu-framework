// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/log"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/tkgpackageclient"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/tkgpackagedatamodel"
)

var packageInstallOp = tkgpackagedatamodel.NewPackageOptions()

var packageInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package",
	Args:  cobra.ExactArgs(1),
	RunE:  packageInstall,
}

func init() {
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.PackageName, "package-name", "p", "", "Name of the package to be installed")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.Version, "version", "", "", "Version of the package to be installed")
	packageInstallCmd.Flags().BoolVarP(&packageInstallOp.CreateNamespace, "create-namespace", "", false, "Create namespace if the target namespace does not exist, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.Namespace, "namespace", "n", "default", "Target namespace to install the package, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.ServiceAccountName, "service-account-name", "", "", "Name of an existing service account used to install underlying package contents, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.ValuesFile, "values-file", "f", "", "The path to the configuration values file, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.KubeConfig, "kubeconfig", "", "", "The path to the kubeconfig file, optional")
	packageInstallCmd.Flags().BoolVarP(&packageInstallOp.Wait, "wait", "", true, "Wait for the package reconciliation to complete, optional")
	packageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollInterval, "poll_interval", "", 1*time.Second, "Time interval between subsequent polls of package reconciliation status, optional")
	packageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollTimeout, "poll_timeout", "", 5*time.Minute, "Timeout value for polls of package reconciliation status, optional")
	packageInstallCmd.MarkFlagRequired("package-name") //nolint
	packageInstallCmd.MarkFlagRequired("version")      //nolint
}

func packageInstall(_ *cobra.Command, args []string) error {
	packageInstallOp.InstalledPkgName = args[0]

	pkgClient, err := tkgpackageclient.NewTKGPackageClient(packageInstallOp.KubeConfig)
	if err != nil {
		return err
	}

	if err := pkgClient.InstallPackage(packageInstallOp); err != nil {
		return err
	}

	log.Infof("Added installed package '%s' in namespace '%s'\n", packageInstallOp.InstalledPkgName, packageInstallOp.Namespace)

	return nil
}
