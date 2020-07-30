package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/google/tcpproxy"
	"github.com/kataras/golog"
)

type (
	// Routing define data structure
	Routing struct {
		Src string
		Dst string
	}
)

func main() {
	ip, err := externalIP()
	if err != nil {
		golog.Error(err)

		fmt.Print("Press any keys to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	routingTable := readJSONFile()

	var p tcpproxy.Proxy
	for _, port := range routingTable {
		golog.Infof("%s:%s => 0.0.0.0:%s", ip, port.Src, port.Dst)
		p.AddRoute(":"+port.Src, tcpproxy.To("0.0.0.0:"+port.Dst))
	}
	golog.Fatal(p.Run())
}

func readJSONFile() []Routing {
	// read file
	data, err := ioutil.ReadFile("routes.json")
	if err != nil {
		golog.Fatal(err)
	}

	// unmarshall it
	var routingTable []Routing
	if err := json.Unmarshal(data, &routingTable); err != nil {
		golog.Fatal("error:", err)
	}
	return routingTable
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
