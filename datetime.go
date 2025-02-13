package main

import (
	"time"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
)

func outputLocaldate(now time.Time) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-calendar-today").Color(colors.Scheme("color10")),
		spacer,
		now.Format("Mon Jan 2"),
	).OnClick(click.RunLeft("gsimplecal"))
}

func outputLocaltime(now time.Time) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-clock-time-five-outline").Color(colors.Scheme("color10")),
		spacer,
		now.Format("15:04:05"),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("timezones")
	}))
}
