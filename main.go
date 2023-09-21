// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// sample-bar demonstrates a sample i3bar built using barista.
package main

import (
	"time"

	"barista.run"
	"barista.run/bar"
	"barista.run/colors"
	"barista.run/group/modal"
	"barista.run/outputs"
	"barista.run/pango"

	"barista.run/modules/battery"
	"barista.run/modules/clock"
	"barista.run/modules/cputemp"
	"barista.run/modules/meminfo"
	"barista.run/modules/meta/split"
	"barista.run/modules/sysinfo"
	"github.com/juli3nk/barista-module-brightness"
	"github.com/juli3nk/barista-module-wlan"

	colorful "github.com/lucasb-eyer/go-colorful"
)

var spacer = pango.Text(" ").XXSmall()
var mainModalController modal.Controller

func makeIconOutput(key string) *bar.Segment {
	return outputs.Pango(spacer, pango.Text(key), spacer)
}

func main() {
	// Config
	colors.LoadBarConfig()
	bg := colors.Scheme("background")
	fg := colors.Scheme("statusline")
	if fg != nil && bg != nil {
		_, _, v := fg.Colorful().Hsv()
		if v < 0.3 {
			v = 0.3
		}
		colors.Set("bad", colorful.Hcl(40, 1.0, v).Clamped())
		colors.Set("degraded", colorful.Hcl(90, 1.0, v).Clamped())
		colors.Set("good", colorful.Hcl(120, 1.0, v).Clamped())
	}

	// Datetime
	localtime := clock.Local().Output(time.Second, outputLocaltime)

	// Brightness
	bn := brightness.New().Output(outputBrightness)

	// Battery
	battSummary, battDetail := split.New(battery.All().Output(outputBattery), 1)

	// Wifi
	wifiName, wifiDetails := split.New(wlan.Any().Output(outputWifi), 1)

	// Sys info
	// Load average
	loadAvg := sysinfo.New().Output(outputLoadAvg)

	loadAvgDetail := sysinfo.New().Output(func(s sysinfo.Info) bar.Output {
		return pango.Textf("%0.2f %0.2f", s.Loads[1], s.Loads[2]).Smaller()
	})

	// Uptime
	uptime := sysinfo.New().Output(outputUptime)

	// Free memory
	freeMem := meminfo.New().Output(outputFreeMem)

	// Swap memory
	swapMem := meminfo.New().Output(outputSwapMem)

	temp := cputemp.New().RefreshInterval(2 * time.Second).Output(outputTemp)

	// Display
	mainModal := modal.New()

	mainModal.Mode("sysinfo").
		SetOutput(makeIconOutput("󰄧")).
		Add(loadAvg).
		Detail(loadAvgDetail, uptime).
		Add(freeMem).
		Detail(swapMem, temp)

	mainModal.Mode("network").
		SetOutput(makeIconOutput("󰈀")).
		Add(wifiName).
		Detail(wifiDetails)

	mainModal.Mode("battery").
		SetOutput(nil).
		Summary(battSummary).
		Detail(battDetail)

	var mm bar.Module
	mm, mainModalController = mainModal.Build()
	panic(barista.Run(mm, bn, localtime))
}
