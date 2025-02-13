package main

import (
	"fmt"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"

	"github.com/barista-run/barista/modules/battery"
)

func outputBattery(i battery.Info) bar.Output {
	if i.Status == battery.Disconnected || i.Status == battery.Unknown {
		return nil
	}

	iconName := "battery"
	if i.Status == battery.Charging {
		iconName += "-charging"
	}
	tenth := i.RemainingPct() / 10
	switch {
	case tenth == 0:
		iconName += "-outline"
	case tenth < 10:
		iconName += fmt.Sprintf("-%d0", tenth)
	}

	rem := i.RemainingTime()
	out := outputs.Group()

	// First segment will be used in summary mode.
	batFirst := outputs.Pango(
		pango.Icon("mdi-"+iconName).Color(colors.Scheme("color10")),
		spacer,
		pango.Textf("%d:%02d", int(rem.Hours()), int(rem.Minutes())%60),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("battery")
	}))
	threshold(batFirst,
		false,
		i.RemainingPct() <= 10,
		i.RemainingPct() <= 15,
		i.RemainingPct() <= 25,
	)
	out.Append(batFirst)

	// Others in detail mode.
	out.Append(outputs.Pango(
		pango.Icon("mdi-"+iconName),
		spacer,
		pango.Textf("%d%%", i.RemainingPct()),
		spacer,
		pango.Textf("(%d:%02d)", int(rem.Hours()), int(rem.Minutes())%60),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("battery")
	})))
	out.Append(outputs.Pango(
		pango.Textf("%4.1f/%4.1f", i.EnergyNow, i.EnergyFull),
		pango.Text("Wh").Smaller(),
	))
	out.Append(outputs.Pango(
		pango.Textf("% +6.2f", i.SignedPower()),
		pango.Text("W").Smaller(),
	))

	return out
}
