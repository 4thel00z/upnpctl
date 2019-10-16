package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/huin/goupnp/httpu"
	"github.com/huin/goupnp/ssdp"
	"github.com/spf13/cobra"
	"strings"
	"upnpctl/utils"
)

var (
	jsonOutput     bool
	maxWaitSeconds int = 30
	numSends       int = 10
	scanCmd            = &cobra.Command{
		Use:   "scan",
		Short: "Scan the network for possible Devices and persist them",
		Long:  ``,
		Run:   Scan,
	}
)

type Device struct {
	Stn      string
	Usn      string
	Location string
}

func Scan(cmd *cobra.Command, args []string) {
	// cmd.OutOrStdout().Write()
	client, e := httpu.NewHTTPUClient()
	utils.Cry(e)

	responses, e := ssdp.SSDPRawSearch(client, ssdp.UPNPRootDevice, maxWaitSeconds, numSends)
	utils.Cry(e)

	devices := make([]Device, 0, len(responses))
	for _, resp := range responses {
		stn := resp.Header.Get("Stn")
		usn := resp.Header.Get("Usn")
		loc := resp.Header.Get("Location")
		if strings.TrimSpace(stn) == "" && strings.TrimSpace(usn) == "" && strings.TrimSpace(loc) == "" {
			continue
		}
		device := Device{
			Stn:      stn,
			Usn:      usn,
			Location: loc,
		}
		devices = append(devices, device)

	}
	if jsonOutput {
		bytes, e := json.Marshal(devices)
		utils.Cry(e)
		fmt.Println(string(bytes))
	} else {
		fmt.Println(devices)
	}
}
