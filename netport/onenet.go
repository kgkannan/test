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
		Ifas: []string{"10.1.0.0/31",
			"2001:db8:85a3::370:0001/64"},
		Remotes: []string{"10.1.0.1",
			"2001:db8:85a3::370:0002"},
	},
	{
		NetPort: "net0port1",
		Netns:   "h1",
		Ifa:     "10.1.0.1/31",
		Ifas: []string{"10.1.0.1/31",
			"2001:db8:85a3::370:0002/64"},
		Remotes: []string{"10.1.0.0",
			"2001:db8:85a3::370:0001"},
	},
}

// Ipv6: OneNet virtual network:
//
//	h0:net0port0 <-> h1:net0port1
var OneIp6Net = NetDevs{
	{
		NetPort: "net0port0",
		Netns:   "h0",
		Ifa:     "2001:db8:85a3::370:0001/64",
		Remotes: []string{"2001:db8:85a3::370:0002"},
	},
	{
		NetPort: "net0port1",
		Netns:   "h1",
		Ifa:     "2001:db8:85a3::370:0002/64",
		Remotes: []string{"2001:db8:85a3::370:0001"},
	},
}
