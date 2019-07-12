// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package netport

// TwoNets virtual network:
//
//	h1:net0port0 <-> r:net0port1
//	h2:net1port0 <-> r:net1port1
var TwoNets = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "10.1.0.0/31",
		Routes: []Route{
			{"10.1.0.2/31", "10.1.0.1"},
		},
		Remotes: []string{"10.1.0.2"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "10.1.0.1/31",
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "10.1.0.2/31",
		Routes: []Route{
			{"10.1.0.0/31", "10.1.0.3"},
		},
		Remotes: []string{"10.1.0.0"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "10.1.0.3/31",
	},
}

// TwoVlanNets virtual network:
//
//	h1:net0port0.1 <-> r:net0port1.1
//	h2:net1port0.2 <-> r:net1port1.2
var TwoVlanNets = NetDevs{
	{
		Vlan:    1,
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "10.1.0.0/31",
		Routes: []Route{
			{"10.1.0.2/31", "10.1.0.1"},
		},
		Remotes: []string{"10.1.0.2"},
	},
	{
		Vlan:    1,
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "10.1.0.1/31",
	},
	{
		Vlan:    2,
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "10.1.0.2/31",
		Routes: []Route{
			{"10.1.0.0/31", "10.1.0.3"},
		},
		Remotes: []string{"10.1.0.0"},
	},
	{
		Vlan:    2,
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "10.1.0.3/31",
	},
}

// Ipv6: TwoNets virtual network:
//
//	h1:net0port0 <-> r:net0port1
//	h2:net1port0 <-> r:net1port1
//port /64 testcase in script
var TwoIp6Nets = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "2001:db8:85a3::370:0001/64",
		Routes: []Route{
			//{"3001:db8:85a3::/64", "2001:db8:85a3::370:0002"},
			{"3001:db8:85a3::370:0002/64", "2001:db8:85a3::370:0002"},
		},
		Remotes: []string{"3001:db8:85a3::370:0002"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "2001:db8:85a3::370:0002/64",
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "3001:db8:85a3::370:0002/64",
		Routes: []Route{
			//{"2001:db8:85a3::/64", "3001:db8:85a3::370:0001"},
			{"2001:db8:85a3::370:0001/64", "3001:db8:85a3::370:0001"},
		},
		Remotes: []string{"2001:db8:85a3::370:0001"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "3001:db8:85a3::370:0001/64",
	},
}
