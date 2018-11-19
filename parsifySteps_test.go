/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

import (
	"fmt"
	"strings"
	"testing"
)

var jsonString = `{
	"jsonrpc":"2.0",
	"id":0,
	"result": {
		"timestamp":"0x5bf1ec6e",
		"extraData":"0x65746865726d696e652d657535",
		"logsBloom":"0x00102000004000020001020100040040109884082000001850020800c02e010400000a080001020880901000000008000200040008400a11082850004000000001040048000800004804010a9400200003400000814000008010000400400100002020000800000000004036121001440302100002248140000200b00020200004000114006100000042000088000249000000c0012100410000008402020092028000000000680000850428020000002184050000000800100000c410000400003042420080000000190010000000000900200192080008880342a03000100020002a4500800820000080808100200518010102008000280000181801232000", "mixHash":"0x83b0860346c8a0223cf65dedc8d7258ffaae1b3d19982d352e8ac36275f10a95",
		"number":"0x66b15f"
	}
}`
var aJSONRpcVersion = "2.0"
var rqID = uint64(0)
var actualTimeStamp = uint64(0x5bf1ec6e)
var actualBlockNumber = uint64(0x66b15f)

func TestParsifyReader(t *testing.T) {
	fmt.Printf("starting\n")
	steps := New("TestMain", []Step{
		{"VerifyStringProperty", "jsonrpc", aJSONRpcVersion, nil},
		{"VerifyNumberProperty", "id", float64(rqID), nil},
		{"EnterKey", "result", nil, nil},
	}, nil, nil)
	data, e := steps.ParsifyReader(strings.NewReader(string(jsonString)))
	if e != nil {
		panic(e)
	}

	type Block struct {
		TimeStamp   uint64
		BlockNumber uint64
	}
	block := Block{}
	ps := &Steps{
		"Parsing result",
		data,
		[]Step{
			{"StoreNumber", "timestamp", nil, &block.TimeStamp},
			{"StoreNumber", "number", nil, &block.BlockNumber},
		},
		0,
		nil,
	}
	if e := ps.Parsify(); e != nil {
		panic(e)
	}

	if block.TimeStamp != actualTimeStamp {
		fmt.Printf("actual TimeStamp: %T %#[1]v\n", block.TimeStamp)
		t.Fail()
	}
	if block.BlockNumber != actualBlockNumber {
		fmt.Printf("actual BlockNumber: %T %#[1]v\n", block.BlockNumber)
		t.Fail()
	}
}
