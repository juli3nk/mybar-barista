package main

import (
	"os/exec"
	"strings"
	"time"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/colors"
	"barista.run/format"
	"barista.run/outputs"
	"barista.run/pango"

	"barista.run/modules/meminfo"
	"barista.run/modules/sysinfo"

	"github.com/martinlindhe/unit"
)

func deviceForMountPath(path string) string {
	mnt, _ := exec.Command("df", "-P", path).Output()
	lines := strings.Split(string(mnt), "\n")
	if len(lines) > 1 {
		devAlias := strings.Split(lines[1], " ")[0]
		dev, _ := exec.Command("realpath", devAlias).Output()
		devStr := strings.TrimSpace(string(dev))
		if devStr != "" {
			return devStr
		}
		return devAlias
	}
	return ""
}

func threshold(out *bar.Segment, urgent bool, color ...bool) *bar.Segment {
	if urgent {
		return out.Urgent(true)
	}
	colorKeys := []string{"bad", "degraded", "good"}
	for i, c := range colorKeys {
		if len(color) > i && color[i] {
			return out.Color(colors.Scheme(c))
		}
	}
	return out
}

func outputLoadAvg(s sysinfo.Info) bar.Output {
	out := outputs.Pango(
		pango.Icon("mdi-desktop-tower").Alpha(0.6),
		pango.Textf("%0.2f", s.Loads[0]),
	)

	// Load averages are unusually high for a few minutes after boot.
	if s.Uptime < 10*time.Minute {
		// so don't add colours until 10 minutes after system start.
		return out
	}

	threshold(out,
		s.Loads[0] > 128 || s.Loads[2] > 64,
		s.Loads[0] > 64 || s.Loads[2] > 32,
		s.Loads[0] > 32 || s.Loads[2] > 16,
	)
	out.OnClick(click.Left(func() {
		mainModalController.Toggle("sysinfo")
	}))

	return out
}

func outputUptime(s sysinfo.Info) bar.Output {
	u := s.Uptime
	var uptimeOut *pango.Node
	if u.Hours() < 24 {
		uptimeOut = pango.Textf("%d:%02d",
			int(u.Hours()), int(u.Minutes())%60)
	} else {
		uptimeOut = pango.Textf("%dd%02dh",
			int(u.Hours()/24), int(u.Hours())%24)
	}
	return pango.Icon("mdi-trending-up").Alpha(0.6).Concat(uptimeOut)
}

// Free memory
func outputFreeMem(m meminfo.Info) bar.Output {
	out := outputs.Pango(
		pango.Icon("material-memory").Alpha(0.8),
		format.IBytesize(m.Available()),
	)
	freeGigs := m.Available().Gigabytes()
	threshold(out,
		freeGigs < 0.5,
		freeGigs < 1,
		freeGigs < 2,
		freeGigs > 12)
	out.OnClick(click.Left(func() {
		mainModalController.Toggle("sysinfo")
	}))
	return out
}

// Swap memory
func outputSwapMem(m meminfo.Info) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-swap-horizontal").Alpha(0.8),
		format.IBytesize(m["SwapTotal"]-m["SwapFree"]), spacer,
		pango.Textf("(% 2.0f%%)", (1-m.FreeFrac("Swap"))*100.0).Small(),
	)
}

func outputTemp(temp unit.Temperature) bar.Output {
	out := outputs.Pango(
		pango.Icon("mdi-fan").Alpha(0.6), spacer,
		pango.Textf("%2d℃", int(temp.Celsius())),
	)
	threshold(out,
		temp.Celsius() > 90,
		temp.Celsius() > 70,
		temp.Celsius() > 60,
	)
	return out
}
