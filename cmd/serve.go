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
	addr      string
	ifaceName string
	serveCmd  = &cobra.Command{
		Use:   "serve",
		Short: "Serve a UPNP honeypot that receives all the events and eventually persists them",
		Long:  ``,
		Run:   Serve,
	}
)

type Honeypot struct{}

func (Honeypot) ServeMessage(r *http.Request) {
	fmt.Println("Host:",r.Host)
	fmt.Println("Headers:",r.Header)
	fmt.Println("URL",r.URL)
}

func Serve(cmd *cobra.Command, args []string) {
	var (
		e     error
		iface *net.Interface
	)

	if ifaceName != "" {
		iface, e = net.InterfaceByName(ifaceName)
		utils.Cry(e)
	}


	honeypot := Honeypot{}
	server := httpu.Server{
		Addr:      addr,
		Multicast: multicast,
		Interface: iface,
		Handler:   honeypot,
	}

	utils.Cry(server.ListenAndServe())
}
