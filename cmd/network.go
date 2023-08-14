package cmd

import (
	"github.com/spf13/cobra"

	"github.com/matoszz/blowfish/config"
	"github.com/matoszz/blowfish/utils"
	"github.com/stmcginnis/gofish/redfish"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Commands for viewing and interacting with network objects",
}

func init() {
	networkCmd.AddCommand(NewGetnetworkCmd())
	rootCmd.AddCommand(networkCmd)
	networkCmd.PersistentFlags().StringP("connection", "c", config.GetDefault(), "The stored connection name to use.")
}

// NewGetnetworkCmd returns a command for getting network information.
func NewGetnetworkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [NAME_OR_ID]",
		Short: "Get network information.",
		RunE:  getNetwork,
		Args:  cobra.MaximumNArgs(1),
	}

	return cmd
}

// getNetwork retrieves the network information.
func getNetwork(cmd *cobra.Command, args []string) error {
	connection, _ := cmd.Flags().GetString("connection")

	c, err := utils.GofishClient(connection)
	if err != nil {
		return utils.ErrorExit(cmd, err.Error())
	}

	defer c.Logout()

	chassiss, err := c.Service.Chassis()
	if err != nil {
		return utils.ErrorExit(cmd, "failed to retrieve system information: %v", err)
	}

	na := []*redfish.NetworkAdapter{}

	for _, chassis := range chassiss {
		chassisNetwork, err := chassis.NetworkAdapters()
		if err != nil {
			return utils.ErrorExit(cmd, "failed to get network information for chassis %q", chassis.Name)
		}

		na = append(na, chassisNetwork...)
	}

	writer := utils.NewTableWriter(
		cmd.OutOrStdout(),
		"Model", "manufacturer")

	for _, network := range na {
		if len(args) == 1 && (network.ID != args[0] && network.Name != args[0]) {
			continue
		}

		writer.AddRow(
			network.Model,
			network.Manufacturer)
	}

	if len(args) != 0 && writer.RowCount() == 0 {
		return utils.ErrorExit(cmd, "network '%s' was not found.", args[0])
	}

	writer.Render()

	return nil
}
