// Copyright © 2013, 2014, The Go-LXC Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"runtime"
	"strconv"
	"sync"

	"gopkg.in/lxc/go-lxc.v2"
)

var (
	lxcpath string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&lxcpath, "lxcpath", lxc.DefaultConfigPath(), "Use specified container path")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			c, err := lxc.NewContainer(strconv.Itoa(i), lxcpath)
			if err != nil {
				log.Fatalf("ERROR: %s\n", err.Error())
			}
			defer lxc.PutContainer(c)

			log.Printf("Starting the container (%d)...\n", i)
			if err := c.Start(); err != nil {
				log.Fatalf("ERROR: %s\n", err.Error())
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
