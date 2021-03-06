// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/log"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgpackageclient"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgpackagedatamodel"
)

var packageInstallOp = tkgpackagedatamodel.NewPackageOptions()

var packageInstallCmd = &cobra.Command{
	Use:   "install INSTALLED_PACKAGE_NAME",
	Short: "Install a package",
	Args:  cobra.ExactArgs(1),
	Example: `
    # Install package contour with installed package name as 'contour-pkg' with specified version and without waiting for package reconciliation to complete 	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --wait=false
	
    # Install package contour with kubeconfig flag and waiting for package reconcilaition to complete	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --kubeconfig path/to/kubeconfig`,
	RunE: packageInstall,
}

func init() {
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.PackageName, "package-name", "p", "", "Name of the package to be installed")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.Version, "version", "v", "", "Version of the package to be installed")
	packageInstallCmd.Flags().BoolVarP(&packageInstallOp.CreateNamespace, "create-namespace", "", false, "Create namespace if the target namespace does not exist, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.Namespace, "namespace", "n", "default", "Target namespace to install the package, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.ServiceAccountName, "service-account-name", "", "", "Name of an existing service account used to install underlying package contents, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.ValuesFile, "values-file", "f", "", "The path to the configuration values file, optional")
	packageInstallCmd.Flags().StringVarP(&packageInstallOp.KubeConfig, "kubeconfig", "", "", "The path to the kubeconfig file, optional")
	packageInstallCmd.Flags().BoolVarP(&packageInstallOp.Wait, "wait", "", true, "Wait for the package reconciliation to complete, optional. To disable wait, specify --wait=false")
	packageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollInterval, "poll-interval", "", tkgpackagedatamodel.DefaultPollInterval, "Time interval between subsequent polls of package reconciliation status, optional")
	packageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollTimeout, "poll-timeout", "", tkgpackagedatamodel.DefaultPollTimeout, "Timeout value for polls of package reconciliation status, optional")
	packageInstallCmd.MarkFlagRequired("package-name") //nolint
	packageInstallCmd.MarkFlagRequired("version")      //nolint
}

func packageInstall(_ *cobra.Command, args []string) error {
	packageInstallOp.PkgInstallName = args[0]

	pkgClient, err := tkgpackageclient.NewTKGPackageClient(packageInstallOp.KubeConfig)
	if err != nil {
		return err
	}

	pp := &tkgpackagedatamodel.PackageProgress{
		ProgressMsg: make(chan string, 10),
		Err:         make(chan error),
		Done:        make(chan struct{}),
		Success:     make(chan bool),
	}
	go pkgClient.InstallPackage(packageInstallOp, pp, false)

	initialMsg := fmt.Sprintf("Installing package '%s'", packageInstallOp.PackageName)
	successMsg := fmt.Sprintf("Added installed package '%s' in namespace '%s'", packageInstallOp.PkgInstallName, packageInstallOp.Namespace)
	if err := displayProgress(initialMsg, successMsg, pp); err != nil {
		return err
	}

	return nil
}

func displayProgress(initialMsg, successMsg string, pp *tkgpackagedatamodel.PackageProgress) error {
	var currMsg string

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	if err := s.Color("bgBlack", "bold", "fgWhite"); err != nil {
		return err
	}
	s.Suffix = fmt.Sprintf(" %s", initialMsg)
	s.Start()

	defer func() {
		if s.Active() {
			s.Stop()
		}
	}()
	for {
		select {
		case err := <-pp.Err:
			s.FinalMSG = fmt.Sprintf("%s\n", err.Error())
			return err
		case msg := <-pp.ProgressMsg:
			if msg != currMsg {
				log.Infof("\n")
				s.Suffix = fmt.Sprintf(" %s", msg)
				currMsg = msg
			}
		case <-pp.Done:
			for msg := range pp.ProgressMsg {
				if msg != currMsg {
					log.Infof("\n")
					s.Suffix = fmt.Sprintf(" %s", msg)
					currMsg = msg
				}
			}
			return nil

		case success := <-pp.Success:
			if success {
				log.Infof("\n %s", successMsg)
			}
			return nil
		}
	}
}
