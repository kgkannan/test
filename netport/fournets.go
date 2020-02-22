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

var FourNetsIp6 = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h1",
		Ifa:     "fc01:1:2:3:4:5:6:1/64",
		DummyIfs: []DummyIf{
			{"dummy0", "fc:5:5:5:5:5:5:5"},
		},
		Routes: []Route{
			{"fc02:1:2:3:4:5:6:1/64", "fc01:1:2:3:4:5:6:2"},
			{"fc:6:6:6:6:6:6:6", "fc01:1:2:3:4:5:6:2"},
		},
		Remotes: []string{"fc02:1:2:3:4:5:6:1", "fc:6:6:6:6:6:6:6"},
	},
	{
		NetPort: "net0port1",
		Netns:   "r",
		Ifa:     "fc01:1:2:3:4:5:6:2/64",
		Routes: []Route{
			{"fc:5:5:5:5:5:5:5", "fc01:1:2:3:4:5:6:1"},
		},
	},
	{
		NetPort: "net1port0",
		Netns:   "h2",
		Ifa:     "fc02:1:2:3:4:5:6:1/64",
		DummyIfs: []DummyIf{
			{"dummy0", "fc:6:6:6:6:6:6:6"},
		},
		Routes: []Route{
			{"fc01:1:2:3:4:5:6:1/64", "fc02:1:2:3:4:5:6:2"},
			{"fc:5:5:5:5:5:5:5", "fc02:1:2:3:4:5:6:2"},
		},
		Remotes: []string{"fc01:1:2:3:4:5:6:1", "fc:5:5:5:5:5:5:5"},
	},
	{
		NetPort: "net1port1",
		Netns:   "r",
		Ifa:     "fc02:1:2:3:4:5:6:2/64",
		Routes: []Route{
			{"fc:6:6:6:6:6:6:6", "fc02:1:2:3:4:5:6:1"},
		},
	},
	{
		NetPort: "net2port0",
		Netns:   "h1",
		Ifa:     "fc03:1:2:3:4:5:6:1/64",
		Routes: []Route{
			{"fc04:1:2:3:4:5:6:1/64", "fc03:1:2:3:4:5:6:2"},
			{"fc01:1:2:3:4:5:6:1/64", "fc03:1:2:3:4:5:6:2"},
			{"fc:6:6:6:6:6:6:6", "fc03:1:2:3:4:5:6:2"},
		},
		Remotes: []string{"fc04:1:2:3:4:5:6:1"},
	},
	{
		NetPort: "net2port1",
		Netns:   "r",
		Ifa:     "fc03:1:2:3:4:5:6:2/64",
		Routes: []Route{
			{"fc:5:5:5:5:5:5:5", "fc03:1:2:3:4:5:6:1"},
			{"fc01:1:2:3:4:5:6:1/64", "fc03:1:2:3:4:5:6:1"},
		},
	},
	{
		NetPort: "net3port0",
		Netns:   "h2",
		Ifa:     "fc04:1:2:3:4:5:6:1/64",
		Routes: []Route{
			{"fc03:1:2:3:4:5:6:1/64", "fc04:1:2:3:4:5:6:2"},
			{"fc02:1:2:3:4:5:6:1/64", "fc04:1:2:3:4:5:6:2"},
			{"fc:5:5:5:5:5:5:5", "fc04:1:2:3:4:5:6:2"},
		},
		Remotes: []string{"fc03:1:2:3:4:5:6:1"},
	},
	{
		NetPort: "net3port1",
		Netns:   "r",
		Ifa:     "fc04:1:2:3:4:5:6:2/64",
		Routes: []Route{
			{"fc02:1:2:3:4:5:6:1/64", "fc04:1:2:3:4:5:6:1"},
			{"fc:6:6:6:6:6:6:6", "fc04:1:2:3:4:5:6:1"},
		},
	},
}
