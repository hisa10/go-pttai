// Copyright 2018 The go-pttai Authors
// This file is part of the go-pttai library.
//
// The go-pttai library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-pttai library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-pttai library. If not, see <http://www.gnu.org/licenses/>.

package service

import "time"

// default config
var (
	DefaultConfig = Config{
		MaxPeers:          350,
		MaxImportantPeers: 100,
		MaxMemberPeers:    200,
		MaxRandomPeers:    50,
	}
)

const (
	ProtocolMaxMsgSize = 10 * 1024 * 1024 // 4MB for video-streaming

	SizeOpType   = 4 // optype uint32
	SizeCodeType = 8 // codetype uint64

	SizeChallenge = 16

	HandshakeTimeout    = 60 * time.Second
	IdentifyPeerTimeout = 10 * time.Second
)

// protocol
const (
	_ uint = iota
	Ptt1
)

var (
	ProtocolVersions = [1]uint{Ptt1}
	ProtocolName     = "ptt1"
	ProtocolLengths  = [1]uint64{4}
)

const (
	StatusMsg = 0x00

	CodeTypeJoinMsg    = 0x01
	CodeTypeJoinAckMsg = 0x02
	CodeTypeOpMsg      = 0x03
)
