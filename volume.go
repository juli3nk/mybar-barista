package main

import (
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/click"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
)

func outputVolumeSpeaker(vol int, muted, isMic bool) bar.Output {
	iconName := "low"
	if muted {
		iconName = "off"
	} else {
		if vol >= 40 {
			iconName = "medium"
		}
		if vol >= 70 {
			iconName = "high"
		}
	}

	return outputs.Pango(
		pango.Icon("mdi-volume-"+iconName).Color(colors.Scheme("color10")),
		spacer,
		pango.Textf("%d%%", vol),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("volume")
	}))
}

func outputVolumeMicrophone(vol int, muted, isMic bool) bar.Output {
	iconName := "microphone"
	if muted {
		iconName = "microphone-off"
	}

	return outputs.Pango(
		pango.Icon("mdi-"+iconName),
		spacer,
		pango.Textf("%d%%", vol),
	)
}
