package main

import (
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"

	"github.com/juli3nk/mybar-barista/modules/wlan"
)

func outputWifi(i wlan.Info) bar.Output {
	if i.Connecting() {
		return outputs.Pango(
			pango.Icon("mdi-wifi"),
			"...",
		).Color(colors.Scheme("degraded"))
	}

	out := outputs.Group()

	// First segment shown in summary mode only.
	out.Append(outputs.Pango(
		pango.Icon("mdi-wifi").Color(colors.Scheme("color10")),
		spacer,
		pango.Text(truncate(i.SSID, -9)),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("network")
	})))

	// Full name
	out.Append(outputs.Pango(
		pango.Icon("mdi-wifi"),
		spacer,
		pango.Text(i.SSID),
	))

	// Frequency
	out.Append(outputs.Textf("%2.1fG", i.Frequency.Gigahertz()))

	// Mac address
	out.Append(outputs.Pango(
		pango.Icon("mdi-access-point"),
		spacer,
		pango.Text(i.AccessPointMAC).Small(),
	))

	return out
}
