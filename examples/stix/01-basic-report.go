// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/objects"
	"time"
)

func main() {
	r := objects.NewReport("2.0")

	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo1")
	r.SetPublished(time.Now())

	r.AddObject(r.CreateSTIXUUID("malware"))
	r.AddObject(r.CreateSTIXUUID("campaign"))
	r.AddObject(r.CreateSTIXUUID("sighting"))
	r.AddObject(r.CreateSTIXUUID("sighting"))
	r.AddObject(r.CreateSTIXUUID("threat-actor"))
	r.AddObject(r.CreateSTIXUUID("threat-actor"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))
	r.AddObject(r.CreateSTIXUUID("relationship"))

	for j := 0; j <= 4; j++ {
		r.AddObject(r.CreateSTIXUUID("indicator"))
	}

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
