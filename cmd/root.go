package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "upnpctl",
		Short: "upnpctl is a upnp utility that can discover the ",
		Long:  `TBD`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside of upnpctl", args)
		},
	}
)

const (
	ssdpUDP4Addr = "239.255.255.250:1900"
)

func initCommand() {
	scanCmdFlags := scanCmd.PersistentFlags()
	scanCmdFlags.BoolVarP(&jsonOutput, "json", "j", true, "--json")
	scanCmdFlags.IntVarP(&maxWaitSeconds, "wait", "w", 30, "--wait <num>")
	scanCmdFlags.IntVarP(&numSends, "num", "n", 10, "--num <num>")

	serveCmd.PersistentFlags().StringVarP(&addr, "addr", "a", ssdpUDP4Addr, "--addr <hostname>:<port>")
	serveCmd.PersistentFlags().StringVarP(&ifaceName, "iface", "i", "", "--iface <interfacename>")
	serveCmd.PersistentFlags().BoolVarP(&multicast, "multicast", "m", true, "--multicast <true|false>")

	rootCmd.AddCommand(scanCmd, serveCmd)
}

func Execute() {
	initCommand()
	if err := rootCmd.Execute(); err != nil {
		if rootCmd.Help() != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
