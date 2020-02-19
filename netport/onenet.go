// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package netport

// OneNet virtual network:
//
//	h0:net0port0 <-> h1:net0port1
var OneNet = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h0",
		Ifa:     "10.1.0.0/31",
		Remotes: []string{"10.1.0.1"},
	},
	{
		NetPort: "net0port1",
		Netns:   "h1",
		Ifa:     "10.1.0.1/31",
		Remotes: []string{"10.1.0.0"},
	},
}

var OneNetIp6 = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h0",
		Ifa:     "fc01:1:2:3:4:5:6:1/64",
		Remotes: []string{"fc01:1:2:3:4:5:6:2"},
	},
	{
		NetPort: "net0port1",
		Netns:   "h1",
		Ifa:     "fc01:1:2:3:4:5:6:2/64",
		Remotes: []string{"fc01:1:2:3:4:5:6:1"},
	},
}
