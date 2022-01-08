// Copyright 2019 Path Network, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"net"
	"time"

	"go.uber.org/zap"
)

type Protocol int

const (
	TCP Protocol = iota
	UDP
)

type Options struct {
	Protocol           string
	ListenAddr         string
	TargetAddr4        string
	TargetAddr6        string
	Mark               int
	Verbose            int
	AllowedSubnetsPath string
	AllowedSubnets     []*net.IPNet
	Listeners          int
	Logger             *zap.Logger
	UDPCloseAfterInt   int
	UDPCloseAfter      time.Duration
}
