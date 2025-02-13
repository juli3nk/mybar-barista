package main

import (
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
	security "github.com/juli3nk/mybar-barista/modules/security"
)

func outputSecurity(i security.Info) bar.Output {
	out := outputs.Group()

	// First segment will be used in summary mode.
	// pango.Icon("mdi-security").Alpha(0.6),
	general := outputs.Pango(
		pango.Icon("mdi-security"),
		spacer,
		pango.Text(i.General),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("security")
	}))
	threshold(general,
		false,
		i.General == "Critical",
		i.General == "Degraded",
		false,
		i.General == "Good",
	)
	out.Append(general)

	// Others in detail mode.
	wifi := outputs.Pango(
		pango.Icon("mdi-wifi"),
		spacer,
		pango.Text(i.WiFi),
	)
	threshold(wifi,
		false,
		i.WiFi == "Error",
		false,
		i.WiFi == "Untrusted",
		i.WiFi == "Trusted",
	)
	out.Append(wifi)

	vpn := outputs.Pango(
		pango.Icon("mdi-vpn"),
		spacer,
		pango.Text(i.VPN),
	)
	threshold(vpn,
		false,
		i.WiFi == "Untrusted" && i.VPN == "NotConnected",
		false,
		false,
		i.VPN == "NotRequired",
	)
	out.Append(vpn)

	dns := outputs.Pango(
		pango.Icon("mdi-dns"),
		spacer,
		pango.Text(i.DNS),
	)
	threshold(dns,
		false,
		i.DNS == "Error",
		i.DNS == "Custom",
		false,
		i.DNS == "Local",
	)
	out.Append(dns)

	firewall := outputs.Pango(
		pango.Icon("mdi-wall-fire"),
		spacer,
		pango.Text(i.Firewall),
	)
	threshold(firewall,
		false,
		i.Firewall == "InputUnrestricted" || i.Firewall == "OutputUnrestricted",
		false,
		i.Firewall == "ForwardUnrestricted",
		i.Firewall == "Restricted",
	)
	out.Append(firewall)

	return out
}
