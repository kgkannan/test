// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package netport

// TwoNets virtual network:
//
//	h1:net0port0 <-> r:net0port1
//	h2:net1port0 <-> r:net1port1
var FourNets = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "10.1.0.0/31",
		DummyIfs: []DummyIf{
			{"dummy0", "10.5.5.5"},
		},
		Routes: []Route{
			{"10.1.0.2/31", "10.1.0.1"},
			{"10.6.6.6", "10.1.0.1"},
		},
		Remotes: []string{"10.1.0.2", "10.6.6.6"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "10.1.0.1/31",
		Routes: []Route{
			{"10.5.5.5", "10.1.0.0"},
		},
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "10.1.0.2/31",
		DummyIfs: []DummyIf{
			{"dummy0", "10.6.6.6"},
		},
		Routes: []Route{
			{"10.1.0.0/31", "10.1.0.3"},
			{"10.5.5.5", "10.1.0.3"},
		},
		Remotes: []string{"10.1.0.0", "10.5.5.5"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "10.1.0.3/31",
		Routes: []Route{
			{"10.6.6.6", "10.1.0.2"},
		},
	},
	{
		NetPort: "net2port0",
		Netns:   "h1",
		Ifa:     "10.2.0.0/31",
		Routes: []Route{
			{"10.2.0.2/31", "10.2.0.1"},
			{"10.1.0.0/31", "10.2.0.1"},
			{"10.6.6.6", "10.2.0.1"},
		},
		Remotes: []string{"10.2.0.2"},
	},
	{
		NetPort: "net2port1",
		Netns:   "r",
		Ifa:     "10.2.0.1/31",
		Routes: []Route{
			{"10.5.5.5", "10.2.0.0"},
			{"10.1.0.0/31", "10.2.0.0"},
		},
	},
	{
		NetPort: "net3port0",
		Netns:   "h2",
		Ifa:     "10.2.0.2/31",
		Routes: []Route{
			{"10.2.0.0/31", "10.2.0.3"},
			{"10.1.0.2/31", "10.2.0.3"},
			{"10.5.5.5", "10.2.0.3"},
		},
		Remotes: []string{"10.2.0.0"},
	},
	{
		NetPort: "net3port1",
		Netns:   "r",
		Ifa:     "10.2.0.3/31",
		Routes: []Route{
			{"10.1.0.2/31", "10.2.0.2"},
			{"10.6.6.6", "10.2.0.2"},
		},
	},
}

// TwoVlanNets virtual network:
//
//	h1:net0port0.1 <-> r:net0port1.1
//	h2:net1port0.2 <-> r:net1port1.2
var FourVlanNets = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Vlan:    1,
		Ifa:     "10.1.0.0/31",
		DummyIfs: []DummyIf{
			{"dummy0", "10.5.5.5"},
		},
		Routes: []Route{
			{"10.1.0.2/31", "10.1.0.1"},
			{"10.6.6.6", "10.1.0.1"},
		},
		Remotes: []string{"10.1.0.2", "10.6.6.6"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Vlan:    1,
		Ifa:     "10.1.0.1/31",
		Routes: []Route{
			{"10.5.5.5", "10.1.0.2"},
		},
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Vlan:    2,
		Ifa:     "10.1.0.2/31",
		DummyIfs: []DummyIf{
			{"dummy0", "10.6.6.6"},
		},
		Routes: []Route{
			{"10.1.0.0/31", "10.1.0.3"},
			{"10.5.5.5", "10.1.0.3"},
		},
		Remotes: []string{"10.1.0.0", "10.5.5.5"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Vlan:    2,
		Ifa:     "10.1.0.3/31",
		Routes: []Route{
			{"10.6.6.6", "10.1.0.2"},
		},
	},
	{
		NetPort: "net2port0",
		Netns:   "h1",
		Vlan:    3,
		Ifa:     "10.2.0.0/31",
		Routes: []Route{
			{"10.2.0.2/31", "10.2.0.1"},
			{"10.6.6.6", "10.2.0.1"},
		},
		Remotes: []string{"10.2.0.2"},
	},
	{
		NetPort: "net2port1",
		Netns:   "r",
		Vlan:    3,
		Ifa:     "10.2.0.1/31",
		Routes: []Route{
			{"10.5.5.5", "10.2.0.2"},
		},
	},
	{
		NetPort: "net3port0",
		Netns:   "h2",
		Vlan:    4,
		Ifa:     "10.2.0.2/31",
		Routes: []Route{
			{"10.2.0.0/31", "10.2.0.3"},
			{"10.5.5.5", "10.2.0.3"},
		},
		Remotes: []string{"10.2.0.0"},
	},
	{
		NetPort: "net3port1",
		Netns:   "r",
		Vlan:    4,
		Ifa:     "10.2.0.3/31",
		Routes: []Route{
			{"10.6.6.6", "10.2.0.2"},
		},
	},
}

// Ipv6: TwoNets virtual network:
//
//	h1:net0port0 <-> r:net0port1
//	h2:net1port0 <-> r:net1port1
//port /64 testcase in script
//TBD: match ip4 subnet like addresses; review netport config for 4 net
var FourIp6Nets = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "2001:db8:85a3::370:0001/64",
		DummyIfs: []DummyIf{
			{"dummy0", "2222:db8:85a3::370:0001"},
		},
		Routes: []Route{
			{"2002:db8:85a3::370:0001/64", "2001:db8:85a3::370:0002"},
			{"3333:db8:85a3::370:0001", "2001:db8:85a3::370:0002"},
		},
		Remotes: []string{"2002:db8:85a3::370:0001", "3333:db8:85a3::370:0001"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "2001:db8:85a3::370:0002/64",
		Routes: []Route{
			{"2222:db8:85a3::370:0001", "2001:db8:85a3::370:0001"},
		},
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "2002:db8:85a3::370:0001/64",
		DummyIfs: []DummyIf{
			{"dummy0", "3333:db8:85a3::370:0001"},
		},
		Routes: []Route{
			{"2001:db8:85a3::370:0001/64", "2002:db8:85a3::370:0002"},
			{"2222:db8:85a3::370:0001", "2002:db8:85a3::370:0002"},
		},
		Remotes: []string{"2001:db8:85a3::370:0002", "2222:db8:85a3::370:0001"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "2002:db8:85a3::370:0002/64",
		Routes: []Route{
			{"3333:db8:85a3::370:0001", "2002:db8:85a3::370:0001"},
		},
	},
	{
		NetPort: "net2port0",
		Netns:   "h1",
		Ifa:     "3001:db8:85a3::370:0001/64",
		Routes: []Route{
			{"3002:db8:85a3::370:0001/64", "3001:db8:85a3::370:0002"},
			{"2001:db8:85a3::370:0001/64", "3001:db8:85a3::370:0002"},
			{"3333:db8:85a3::370:0001", "3001:db8:85a3::370:0002"},
		},
		Remotes: []string{"3002:db8:85a3::370:0001"},
	},
	{
		NetPort: "net2port1",
		Netns:   "r",
		Ifa:     "3001:db8:85a3::370:0002/64",
		Routes: []Route{
			{"2222:db8:85a3::370:0001", "3001:db8:85a3::370:0001"},
			{"2001:db8:85a3::370:0001/64", "3001:db8:85a3::370:0001"},
		},
	},
	{
		NetPort: "net3port0",
		Netns:   "h2",
		Ifa:     "3002:db8:85a3::370:0001/64",
		Routes: []Route{
			{"3001:db8:85a3::370:0001/64", "3002:db8:85a3::370:0002"},
			{"2002:db8:85a3::370:0001/64", "3002:db8:85a3::370:0002"},
			{"2222:db8:85a3::370:0001", "3002:db8:85a3::370:0002"},
		},
		Remotes: []string{"3001:db8:85a3::370:0001"},
	},
	{
		NetPort: "net3port1",
		Netns:   "r",
		Ifa:     "3002:db8:85a3::370:0002/64",
		Routes: []Route{
			{"3001:db8:85a3::370:0001/64", "3002:db8:85a3::370:0001"},
			{"3333:db8:85a3::370:0001", "3002:db8:85a3::370:0001"},
		},
	},
}
