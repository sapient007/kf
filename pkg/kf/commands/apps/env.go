package apps

import (
	"fmt"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/kf/pkg/kf"
	"github.com/GoogleCloudPlatform/kf/pkg/kf/commands/config"
	"github.com/spf13/cobra"
)

// EnvironmentClient interacts with app's environment variables.
type EnvironmentClient interface {
	// List shows all the names and values of the environment variables for an
	// app.
	List(appName string, opts ...kf.ListEnvOption) (map[string]string, error)

	// Set sets the given environment variables.
	Set(appName string, values map[string]string, opts ...kf.SetEnvOption) error

	// Unset unsets the given environment variables.
	Unset(appName string, names []string, opts ...kf.UnsetEnvOption) error
}

// NewEnvCommand creates a Env command.
func NewEnvCommand(p *config.KfParams, c EnvironmentClient) *cobra.Command {
	var envCmd = &cobra.Command{
		Use:   "env APP_NAME",
		Short: "List the names and values of the environment variables for an app",
		Args:  cobra.ExactArgs(1),
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			appName := args[0]
			cmd.SilenceUsage = true

			values, err := c.List(
				appName,
				kf.WithListEnvNamespace(p.Namespace),
			)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(p.Output, 8, 4, 1, ' ', tabwriter.StripEscape)
			fmt.Fprintln(w, "NAME\tVALUE")
			for name, value := range values {
				fmt.Fprintf(w, "%s\t%s\n", name, value)
			}
			w.Flush()

			return nil
		},
	}

	return envCmd
}