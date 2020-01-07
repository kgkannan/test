// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package netport parses testdata/netport.yaml for interface assingments along
// with utilities to build and test virtual networks configured from these
// assignments.
package netport

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/platinasystems/test"
	"gopkg.in/yaml.v2"
)

const NetPortFile = "testdata/netport.yaml"

var Goes string
var PortByNetPort, NetPortByPort map[string]string

func Init(goes string) {
	Goes = goes
	b, err := ioutil.ReadFile(NetPortFile)
	if err != nil {
		panic(err)
	}
	PortByNetPort = make(map[string]string)
	NetPortByPort = make(map[string]string)
	if err = yaml.Unmarshal(b, PortByNetPort); err != nil {
		panic(fmt.Errorf("%s: %v", NetPortFile, err))
	}
	for netport, port := range PortByNetPort {
		sysport := filepath.Join("/sys/class/net", port)
		if _, err = os.Stat(sysport); err != nil {
			panic(err)
		}
		NetPortByPort[port] = netport
	}

	IsHighVer = IsHigherVer()
}

type Route struct {
	Prefix string
	GW     string
}

type DummyIf struct {
	Ifname string
	Ifa    string
}

const (
	NETPORT_DEVTYPE_PORT = iota
	NETPORT_DEVTYPE_PORT_VLAN
	NETPORT_DEVTYPE_BRIDGE
	NETPORT_DEVTYPE_BRIDGE_PORT
)

// NetDev describes a network interface configuration.
// DevType is determined by IsBridge, Vlan, and Upper
// default PORT sets ifa for ifname derived from NetPort
// BRIDGE adds a linux bridge device with ifname and ifa
// PORT_VLAN adds a linux vlan device to ifname derived from NetPort with ifa
// BRIDGE_PORT adds a linux vlan device and sets upper to named bridge (no ifa)
// initial values filled from NetDev[], then DevType and Ifname updated
type NetDev struct {
	DevType       int
	IsBridge      bool
	BridgeIfindex int // for BRIDGE to avoid netns overlap
	BridgeMac     string
	Vlan          int    // for PORT_VLAN or BRIDGE_PORT
	NetPort       string // lookup key for NetPortFile to Ifname
	Netns         string
	Ifname        string
	Upper         string // only for BRIDGE_PORT

	Ifa      string // no ifa for BRIDGE_PORT
	DummyIfs []DummyIf
	Routes   []Route
	Remotes  []string
}

//placeholder for 1.3 check
var IsHighVer bool

// NetDevs describe all of the interfaces in the virtual network under test.
type NetDevs []NetDev

//set environment info based on platform software
func IsHigherVer() (higher bool) {
	pd := *PlatformDriver
	o, err := exec.Command("/sbin/modinfo", pd).Output()
	if err == nil && len(o) > 0 {
		const depends = "depends:"
		s := string(o)
		i := strings.Index(s, depends)
		if i > 0 {
			s = s[i+len(depends):]
			dependencies := 0
			drvdepends := [...]string{
				"onie",
				"xeth",
			}
			for _, d := range drvdepends {
				if strings.Index(s, d) >= 0 {
					dependencies += 1
				}
			}
			if dependencies == 2 {
				higher = true
			}
		}
	}
	fmt.Printf("\n IsHigherVer %v \n", higher)
	return
}

func IsIPv6(prefix string) (is_ip6 bool) {
	toip, _, _ := net.ParseCIDR(prefix)
	is_ip6 = (len(toip) == net.IPv6len && toip.To4() == nil && toip.To16() != nil)
	fmt.Printf("p:%s ipn:%s is_ip6:%v\n", prefix, toip, is_ip6)
	return
}

// netdevs list the interface configurations of the network under test
func (netdevs NetDevs) Test(t *testing.T, tests ...test.Tester) {
	test.SkipIfDryRun(t)
	assert := test.Assert{t}
	cleanup := test.Cleanup{t}
	for i := range netdevs {
		nd := &netdevs[i]
		if nd.IsBridge {
			nd.DevType = NETPORT_DEVTYPE_BRIDGE
		} else if nd.Vlan != 0 {
			if nd.Upper != "" {
				nd.DevType = NETPORT_DEVTYPE_BRIDGE_PORT
			} else {
				nd.DevType = NETPORT_DEVTYPE_PORT_VLAN
			}
		} else {
			nd.DevType = NETPORT_DEVTYPE_PORT
		}

		ns := nd.Netns
		_, err := os.Stat(filepath.Join("/var/run/netns", ns))
		if err != nil {
			assert.Program(Goes, "ip", "netns", "add", ns)
			defer cleanup.Program(Goes, "ip", "netns", "del", ns)
		}

		if nd.DevType == NETPORT_DEVTYPE_BRIDGE {
			link := "bridge"
			if isHighVer {
				link = "xeth-bridge"
			}
			if nd.BridgeIfindex != 0 {
				assert.Program(Goes, "ip", "-n", ns,
					"link", "add", nd.Ifname,
					"index", nd.BridgeIfindex,
					"type", link)
			} else {
				assert.Program(Goes, "ip", "-n", ns,
					"link", "add", nd.Ifname,
					"type", link)
			}

			if false && nd.BridgeMac != "" {
				assert.Program(Goes, "ip", "-n", ns,
					"link", "set", nd.Ifname,
					"address", nd.BridgeMac)
			}
			assert.Program(Goes, "ip", "-n", ns,
				"link", "set", nd.Ifname, "up")
			defer cleanup.Program(Goes, "ip", "-n", ns,
				"link", "del", nd.Ifname)
		} else {
			link := "vlan"
			if isHighVer {
				link = "xeth-vlan"
			}
			ifname := PortByNetPort[nd.NetPort]
			if nd.Vlan != 0 {
				link := ifname
				ifname += fmt.Sprint(".", nd.Vlan)
				assert.Program(Goes, "ip", "link", "set", link,
					"up")
				assert.Program(Goes, "ip", "link", "add",
					ifname, "link", link, "type", link,
					"id", nd.Vlan)
				defer cleanup.Program(Goes, "ip", "link",
					"del", ifname)
			}
			nd.Ifname = ifname
			assert.ProgramRetry(3, Goes, "ip", "link", "set",
				nd.Ifname, "up", "netns", ns)
			defer cleanup.Program(Goes, "ip", "-n", ns,
				"link", "set", nd.Ifname, "down", "netns", 1)
		}

		if nd.DevType == NETPORT_DEVTYPE_BRIDGE_PORT {
			assert.Program(Goes, "ip", "-n", ns,
				"link", "set", nd.Ifname, "master", nd.Upper)
			defer cleanup.Program(Goes, "ip", "-n", ns,
				"link", "set", nd.Ifname, "nomaster")
		} else if nd.Ifa != "" {
			assert.ProgramRetry(3, Goes, "ip", "-n", ns,
				"address", "add", nd.Ifa, "dev", nd.Ifname)
			defer cleanup.Program(Goes, "ip", "-n", ns,
				"address", "del", nd.Ifa, "dev", nd.Ifname)
			for _, route := range nd.Routes {
				prefix := route.Prefix
				gw := route.GW
				assert.Program(Goes, "ip", "-n", ns,
					"route", "add", prefix, "via", gw)
			}
		}
		if *test.VVV {
			t.Logf("nd %+v\n", nd)
		}
	}
	test.Tests(tests).Test(t)
}
