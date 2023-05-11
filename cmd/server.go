
package cmd

import (
	"blockchain-sharding-master/common"
	"blockchain-sharding-master/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A blockchain node and API server",
	Long: `Server is a full blockchain node to connect with other nodes, which will 
construct a peer-to-peer network. It runs a web server to provides REST APIs.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		common.SetConfig(config)
		server.StartServer()
	},
}

var config = common.GetConfig()

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.
	serverCmd.Flags().StringVarP(&config.IP, "ip", "i", common.DefaultIP, "the IP address of the server")
	serverCmd.Flags().IntVarP(&config.APIPort, "api-port", "a", common.DefaultAPIPort, "which port the API service listen on")
	serverCmd.Flags().IntVarP(&config.RPCPort, "rpc-port", "r", common.DefaultRPCPort, "which port the blockchain node listen on")
}
