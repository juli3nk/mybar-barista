package main

import (
	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/colors"
	"barista.run/outputs"
	"barista.run/pango"

	"github.com/juli3nk/barista-module-wlan"
)

func outputWifi(i wlan.Info) bar.Output {
	if !i.Connecting() && !i.Connected() {
		mainModalController.SetOutput("network", makeIconOutput("󰖪"))
		return nil
	}

	mainModalController.SetOutput("network", makeIconOutput("󰖩"))
	if i.Connecting() {
		return outputs.Pango(pango.Text("mdi-wifi").Alpha(0.6), "...").
			Color(colors.Scheme("degraded"))
	}

	out := outputs.Group()

	// First segment shown in summary mode only.
	out.Append(outputs.Pango(
		pango.Text("󰖩").Alpha(0.6),
		pango.Text(truncate(i.SSID, -9)),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("network")
	})))

	// Full name, frequency, bssid in detail mode
	out.Append(outputs.Pango(
		pango.Text("󰖩").Alpha(0.6),
		pango.Text(i.SSID),
	))
	out.Append(outputs.Textf("%2.1fG", i.Frequency.Gigahertz()))
	out.Append(outputs.Pango(
		pango.Text("󰀃").Alpha(0.8),
		pango.Text(i.AccessPointMAC).Small(),
	))

	return out
}
