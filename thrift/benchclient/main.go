// Copyright (c) 2015 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// benchclient is used to make requests to a specific server.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/uber/tchannel-go/testutils"
	"github.com/uber/tchannel-go/thrift"
	gen "github.com/uber/tchannel-go/thrift/gen-go/test"
)

func main() {
	ch, err := testutils.NewClient(nil)
	if err != nil {
		log.Fatalf("err")
	}

	ch.Peers().Add(os.Args[1])
	thriftClient := thrift.NewClient(ch, "bench-server", nil)
	client := gen.NewTChanSecondServiceClient(thriftClient)

	fmt.Println("bench-client started")

	rdr := bufio.NewReader(os.Stdin)
	for {
		line, err := rdr.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("stdin read failed: %v", err)
		}

		line = strings.TrimSuffix(line, "\n")
		switch line {
		case "call":
			makeCall(client)
		case "quit":
			return
		default:
			log.Fatalf("unrecognized command: %v", line)
		}
	}
}

var arg string

func makeArg() string {
	if len(arg) > 0 {
		return arg
	}

	bs := []byte{}
	// TODO(prashant) when this is 100000, get more arguments in message error.
	for i := 0; i < 10000; i++ {
		bs = append(bs, byte(i%26+'A'))
	}
	arg = string(bs)
	return arg
}

func makeCall(client gen.TChanSecondService) {
	ctx, cancel := thrift.NewContext(time.Second)
	defer cancel()

	arg := makeArg()
	started := time.Now()
	res, err := client.Echo(ctx, arg)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	if res != arg {
		log.Fatalf("Echo gave different string!")
	}
	duration := time.Since(started)
	fmt.Println(duration)
}