package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"upnpctl/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "upnpctl",
		Short: "upnpctl is a upnp utility that can discover the ",
		Long: `
                                                                             _..._                
                                                                          .-'_..._''.       .---. 
           _________   _...._         _..._   _________   _...._        .' .'      '.\      |   | 
           \        |.'      '-.    .'     '. \        |.'      '-.    / .'                 |   | 
            \        .'''''.    '. .   .-.   . \        .'''''.    '. . '               .|  |   | 
             \      |       \     \|  '   '  |  \      |       \     \| |             .' |_ |   | 
   _    _     |     |        |    ||  |   |  |   |     |        |    || |           .'     ||   | 
  | '  / |    |      \      /    . |  |   |  |   |      \      /    . . '          '--.  .-'|   | 
 .' | .' |    |     |\''-.-'   .'  |  |   |  |   |     |\''-.-'   .'   \ '.          .|  |  |   | 
 /  | /  |    |     | '-....-''    |  |   |  |   |     | '-....-''      '. '._____.-'/|  |  |   | 
|   ''.  |   .'     '.             |  |   |  |  .'     '.                 '-.______ / |  '.''---' 
'   .'|  '/'-----------'           |  |   |  |'-----------'                        '  |   /       
 '-'  '--'                         '--'   '--'                                        ''-'        

`,
		Run: func(cmd *cobra.Command, args []string) {
			utils.Cry(cmd.Help())
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
	serveCmd.PersistentFlags().BoolVarP(
		&printBody,
		"print-body",
		"p",
		false,
		"--print-body <true|false> (doesn't really make sense for ssdp discovery)",
	)

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
