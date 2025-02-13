package main

import (
	"flag"
	"time"

	"github.com/barista-run/barista"
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/group/modal"
	"github.com/barista-run/barista/pango"
	"github.com/barista-run/barista/pango/icons/fontawesome"
	"github.com/barista-run/barista/pango/icons/mdi"

	"github.com/barista-run/barista/modules/battery"
	"github.com/barista-run/barista/modules/clock"
	"github.com/barista-run/barista/modules/cputemp"
	"github.com/barista-run/barista/modules/meminfo"
	"github.com/barista-run/barista/modules/meta/split"
	"github.com/barista-run/barista/modules/sysinfo"
	"github.com/juli3nk/mybar-barista/modules/kbd"
	"github.com/juli3nk/mybar-barista/modules/wlan"
	"github.com/juli3nk/mybar-barista/version"

	brightness "github.com/juli3nk/barista-module-brightness"
	volume "github.com/juli3nk/barista-module-volume"
	security "github.com/juli3nk/mybar-barista/modules/security"
)

var flgVersion bool

var spacer = pango.Text(" ").XXSmall()
var mainModalController modal.Controller

func init() {
	flag.BoolVar(&flgVersion, "version", false, "print version and exit")
	flag.Parse()

	colors.LoadFromMap(map[string]string{
		"default":  "#cccccc",
		"good":     "#4caf50",
		"warning":  "#ffd760",
		"degraded": "#ff9f40",
		"critical": "#ff5454",
		"disabled": "#777777",
		"color0":   "#2e3440",
		"color1":   "#3b4252",
		"color2":   "#434c5e",
		"color3":   "#4c566a",
		"color4":   "#d8dee9",
		"color5":   "#e5e9f0",
		"color6":   "#eceff4",
		"color7":   "#8fbcbb",
		"color8":   "#88c0d0",
		"color9":   "#81a1c1",
		"color10":  "#5e81ac",
		"color11":  "#bf616a",
		"color12":  "#d08770",
		"color13":  "#ebcb8b",
		"color14":  "#a3be8c",
		"color15":  "#b48ead",
	})
}

func main() {
	if flgVersion {
		ver := version.New()
		ver.Show()
		return
	}

	// Fonts
	mdi.Load(home(".local/Github/MaterialDesign-Webfont"))
	fontawesome.Load(home(".local/Github/Font-Awesome"))

	// Sys info
	// Load average
	loadAvg, loadAvgDetail := split.New(sysinfo.New().Output(outputLoadAvg), 1)

	// Free memory
	freeMem := meminfo.New().Output(outputFreeMem)

	// Swap memory
	swapMem := meminfo.New().Output(outputSwapMem)

	// CPU temperature
	cpuTemp := cputemp.New().RefreshInterval(2 * time.Second).Output(outputTemp)

	// Keyboard layout
	kbdlayout := kbd.New().Output(outputKeyboardLayout)

	// Brightness
	bn := brightness.New().Output(outputBrightness)

	// Volume
	// Speaker/Headphones volume (auto-detects active device)
	speaker := volume.New("@DEFAULT_AUDIO_SINK@", false).Output(outputVolumeSpeaker)

	// Microphone volume
	microphone := volume.New("@DEFAULT_AUDIO_SOURCE@", true).Output(outputVolumeMicrophone)

	// Battery
	battSummary, battDetails := split.New(battery.All().Output(outputBattery), 1)

	// Security
	securityGlobal, securityDetails := split.New(
		security.New(home(".config/local/net.json")).Output(outputSecurity), 1,
	)

	// Wifi
	wifiName, wifiDetails := split.New(wlan.Any().Output(outputWifi), 1)

	// Datetime
	localdate := clock.Local().Output(time.Second, outputLocaldate)
	localtime := clock.Local().Output(time.Second, outputLocaltime)

	// Display
	mainModal := modal.New()

	mainModal.Mode("cpu").
		SetOutput(nil).
		Detail(loadAvgDetail, cpuTemp)

	mainModal.Mode("mem").
		SetOutput(nil).
		Detail(swapMem)

	mainModal.Mode("volume").
		SetOutput(nil).
		Detail(microphone)

	mainModal.Mode("battery").
		SetOutput(nil).
		Detail(battDetails)

	mainModal.Mode("security").
		SetOutput(nil).
		Detail(securityDetails)

	mainModal.Mode("network").
		SetOutput(nil).
		Detail(wifiDetails)

	mainModal.Mode("timezones").
		SetOutput(nil).
		Detail(makeTzClock("Montreal", "America/Toronto")).
		Detail(makeTzClock("UTC", "Etc/UTC"))

	var mm bar.Module
	mm, mainModalController = mainModal.Build()

	panic(barista.Run(
		mm,
		loadAvg,
		freeMem,
		kbdlayout,
		bn,
		speaker,
		battSummary,
		securityGlobal,
		wifiName,
		localdate,
		localtime,
	))
}
