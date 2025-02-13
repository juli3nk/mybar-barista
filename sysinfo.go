package main

import (
	"time"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/format"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"

	"github.com/barista-run/barista/modules/meminfo"
	"github.com/barista-run/barista/modules/sysinfo"

	"github.com/martinlindhe/unit"
)

func outputLoadAvg(s sysinfo.Info) bar.Output {
	out := outputs.Group()

	loadFirst := outputs.Pango(
		pango.Icon("mdi-cpu-64-bit").Color(colors.Scheme("color10")),
		spacer,
		pango.Textf("%0.2f", s.Loads[0]),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("cpu")
	}))

	// Load averages are unusually high for a few minutes after boot.
	// so don't add colours until 10 minutes after system start.
	if s.Uptime > 10*time.Minute {
		threshold(loadFirst,
			false,
			s.Loads[0] > 128 || s.Loads[2] > 64,
			s.Loads[0] > 64 || s.Loads[2] > 32,
			s.Loads[0] > 32 || s.Loads[2] > 16,
			false,
		)
	}

	out.Append(loadFirst)

	// Others in detail mode
	out.Append(outputs.Pango(
		pango.Textf("%0.2f %0.2f", s.Loads[1], s.Loads[2]).Smaller(),
	))

	return out
}

// Free memory
func outputFreeMem(m meminfo.Info) bar.Output {
	out := outputs.Pango(
		pango.Icon("mdi-memory").Color(colors.Scheme("color10")),
		spacer,
		format.IBytesize(m.Available()),
	)

	freeGigs := m.Available().Gigabytes()
	threshold(out,
		false,
		freeGigs < 0.5,
		freeGigs < 1,
		freeGigs < 2,
	)

	out.OnClick(click.Left(func() {
		mainModalController.Toggle("mem")
	}))

	return out
}

// Swap memory
func outputSwapMem(m meminfo.Info) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-swap-horizontal"),
		spacer,
		format.IBytesize(m["SwapTotal"]-m["SwapFree"]),
		spacer,
		pango.Textf("(%2.0f%%)", (1-m.FreeFrac("Swap"))*100.0).Small(),
	)
}

func outputTemp(temp unit.Temperature) bar.Output {
	out := outputs.Pango(
		pango.Icon("mdi-fan"),
		spacer,
		pango.Textf("%2d℃", int(temp.Celsius())),
	)

	threshold(out,
		false,
		temp.Celsius() > 90,
		temp.Celsius() > 70,
		temp.Celsius() > 60,
	)

	return out
}
