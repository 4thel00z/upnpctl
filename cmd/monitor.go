package cmd

import (
	"fmt"
	"github.com/huin/goupnp/httpu"
	"github.com/spf13/cobra"
	"net"
	"net/http"
	"upnpctl/utils"
)

var (
	multicast bool
	printBody bool
	addr      string
	ifaceName string
	serveCmd  = &cobra.Command{
		Use:   "monitor",
		Short: "Monitor HTTPU packages and print them on stdout",
		Long:  ``,
		Run:   Monitor,
	}
)

type honeyPot struct{}

func (honeyPot) ServeMessage(r *http.Request) {
	fmt.Println("Host:", r.Host)
	fmt.Println("Headers:", r.Header)
	fmt.Println("URL", r.URL)
	if printBody {
		fmt.Println("Body", r.Body)
	}
}

func Monitor(cmd *cobra.Command, args []string) {
	var (
		e     error
		iface *net.Interface
	)

	if ifaceName != "" {
		iface, e = net.InterfaceByName(ifaceName)
		utils.Cry(e)
	}

	honeypot := honeyPot{}
	server := httpu.Server{
		Addr:      addr,
		Multicast: multicast,
		Interface: iface,
		Handler:   honeypot,
	}

	utils.Cry(server.ListenAndServe())
}
