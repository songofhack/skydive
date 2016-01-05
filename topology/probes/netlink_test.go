/*
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package probes

import (
	"testing"

	"github.com/vishvananda/netlink"
)

type FakeOvsLink struct {
	LinkAttrs netlink.LinkAttrs
}

func (l *FakeOvsLink) Type() string {
	return "openvswitch"
}

func (l *FakeOvsLink) Attrs() *netlink.LinkAttrs {
	return &l.LinkAttrs
}

func TestOvsLinkNotHandled(t *testing.T) {
	topo := NewTopology("host-a")
	root := topo.NewNetNs("root")

	updater := NewNetLinkProbe(root)

	link := &netlink.Device{
		LinkAttrs: netlink.LinkAttrs{
			Name:         "intf1",
			HardwareAddr: []byte("1.1.1.1.1.1"),
		},
	}
	updater.addLinkToTopology(link)

	if root.GetInterface("intf1") == nil {
		t.Error("interface of type device not found in the topology")
	}

	ovsLink := &FakeOvsLink{
		LinkAttrs: netlink.LinkAttrs{
			Name:         "intf2",
			HardwareAddr: []byte("2.2.2.2.2.2"),
		},
	}
	updater.addLinkToTopology(ovsLink)

	if root.GetInterface("intf2") == nil {
		t.Error("interface of type openvswitch can also be added as a standard interface")
	}
}