package security

import (
	"time"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/value"
	"github.com/barista-run/barista/timing"
	"github.com/juli3nk/local-net/pkg/config"
)

type Info struct {
	General  string
	WiFi     string
	VPN      string
	DNS      string
	Firewall string
}

type Module struct {
	outputFunc value.Value
	scheduler  *timing.Scheduler
	config     *config.Config
}

func New(localNetFile string) *Module {
	m := &Module{scheduler: timing.NewScheduler()}
	m.RefreshInterval(5 * time.Second)
	m.loadConfig(localNetFile)
	return m
}

func (m *Module) RefreshInterval(interval time.Duration) *Module {
	m.scheduler.Every(interval)
	return m
}

// Output configures a module to display the output of a user-defined function.
func (m *Module) Output(outputFunc func(Info) bar.Output) *Module {
	m.outputFunc.Set(outputFunc)
	return m
}

func (m *Module) Stream(s bar.Sink) {
	outputFunc := m.outputFunc.Get().(func(Info) bar.Output)

	for {
		generalStatus := "Bad"
		wifiStatus := m.getWifiStatus()
		vpnStatus := m.getVpnStatus()
		dnsStatus := m.getDnsStatus()
		firewallStatus := m.getFirewallStatus()

		totalCount := 3
		count := 0

		if wifiStatus == "Trusted" {
			count++
		} else if wifiStatus == "Untrusted" {
			if vpnStatus == "Connected" {
				count++
			}
		}

		if dnsStatus == "Local" {
			count++
		}

		if firewallStatus == "ForwardUnrestricted" || firewallStatus == "Restricted" {
			count++
		}

		if count == totalCount {
			generalStatus = "Good"
		}
		if count > 0 && count < totalCount {
			generalStatus = "Degraded"
		}

		i := Info{
			General:  generalStatus,
			WiFi:     wifiStatus,
			VPN:      vpnStatus,
			DNS:      dnsStatus,
			Firewall: firewallStatus,
		}

		s.Output(outputFunc(i))

		<-m.scheduler.C
	}
}

func (m *Module) loadConfig(filename string) {
	cfg, _ := config.New(filename)

	m.config = cfg
}
