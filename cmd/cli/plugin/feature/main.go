package main

import (
	"os"
	"time"

	"github.com/aunum/log"
	"github.com/spf13/cobra"

	cliv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/cli/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/command/plugin"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/component"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/clusterclient"
)

var logLevel int32
var logFile string
var activated, available, deactivated, nondiscoverable bool

var descriptor = cliv1alpha1.PluginDescriptor{
	Name:        "feature",
	Description: "operate on features and featuregates",
	Version:     "v0.0.1",
	Group:       cliv1alpha1.ManageCmdGroup, // set group
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err)
	}

	p.Cmd.PersistentFlags().Int32VarP(&logLevel, "verbose", "", 0, "Number for the log level verbosity(0-9)")
	p.Cmd.PersistentFlags().StringVar(&logFile, "log-file", "", "Log file path")

	p.AddCommands(
		FeatureListCmd,
		FeatureActivateCmd,
		FeatureDeactivateCmd,
	)

	FeatureListCmd.Flags().BoolVarP(&activated, "activated", "a", false, "List only activated Features")
	FeatureListCmd.Flags().BoolVarP(&available, "available", "v", false, "List only available Features")
	FeatureListCmd.Flags().BoolVarP(&deactivated, "deactivated", "d", false, "List only deactivated Features")
	FeatureListCmd.Flags().BoolVarP(&nondiscoverable, "include-hidden-features", "", false, "Include intentionally hidden Features in List. Warning: This is for development purposes.")

	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}

var FeatureListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Features",
	Args:  cobra.NoArgs,
	Example: `
    # List a clusters Features
    tanzu feature list --all (default)
    tanzu feature list --activated
    tanzu feature list --availabld
    tanzu feature list --deactivated`,
	RunE: featureList,
}

func featureList(cmd *cobra.Command, _ []string) error {
	clusterClientOptions := clusterclient.Options{GetClientInterval: 2 * time.Second, GetClientTimeout: 5 * time.Second}
	clusterClient, err := clusterclient.NewClient(server.ManagementClusterOpts.Path, server.ManagementClusterOpts.Context, clusterClientOptions)
	if err != nil {
		return err
	}

	//	listTKGFeaturessOptions := client.ListTKGFeaturesOptions{}
	//	var filter string
	//	switch {
	//	case activated:
	//		listTKGFeaturessOptions.activated = true
	//	case available:
	//		listTKGFeaturessOptions.available = true
	//	case deactivated:
	//		listTKGFeaturessOptions.deactivated = true
	//	case discoverable:
	//		listTKGFeaturessOptions.nondiscoverable = true
	//	}
	//	listTKGFeaturessOptions.Filter = filter
	features, err := t.tkgClient.GetFeatures("")
	if err != nil {
		return nil, err
	}
	var t component.OutputWriterSpinner
	t, err = component.NewOutputWriterWithSpinner(cmd.OutOrStdout(), outputFormat,
		"Retrieving Features...", true, "NAME", "DESCRIPTION", "IMMUTABLE", "DISCOVERABLE", "ACTIVATED", "OVERRIDDEN", "GATE")
	if err != nil {
		return err
	}

	// Determine if the state is due to a Gate which is overriding it, and which gate.
	var overridden bool
	if feature.Gate != "" {
		overridden == true
	}

	for feature := range features {
		t.AddRow(
			feature.Name,
			feature.Description,
			feature.Immutable,
			feature.Discoverable,
			feature.Activated,
			overridden,
			feature.Gate)
	}
	t.RenderWithSpinner()

	return nil
}
