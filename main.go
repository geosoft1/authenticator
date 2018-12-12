// authenticator project
// Copyright (C) 2018  geosoft1  geosoft1@gmail.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package json implements encoding and decoding of JSON as defined in RFC 4627
// as high level wrapper over standard json package.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/denisbrodbeck/machineid"
)

var (
	port = flag.String("port", "8080", "port")
)

// https://gist.github.com/rucuriousyet/ab2ab3dc1a339de612e162512be39283
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func main() {
	flag.Parse()
	//https://github.com/denisbrodbeck/machineid
	id, _ := machineid.ID()
	s := struct {
		mac, id string
	}{getMacAddr(), id}
	http.HandleFunc("/authenticator", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `angular.callbacks._0({ "mac":"%s", "id":"%s" })`, s.mac, s.id)
	})
	http.ListenAndServe(":"+*port, nil)
}
