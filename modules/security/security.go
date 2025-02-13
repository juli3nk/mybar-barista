package security

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	"github.com/juli3nk/go-network/iwgetid"
	"github.com/juli3nk/go-network/nmcli"
)

func (m *Module) getWifiStatus() string {
	ssid, err := iwgetid.GetSSID("")
	if err != nil {
		return "Error"
	}

	if len(*ssid) == 0 {
		return "NotConnected"
	}

	for _, name := range m.config.Wifi {
		if *ssid == name {
			return "Trusted"
		}
	}

	return "Untrusted"
}

func (m *Module) getVpnStatus() string {
	if m.getWifiStatus() == "Trusted" {
		return "NotRequired"
	}

	_, err := nmcli.ConnectionShow("wireguard", m.config.Vpn.Name)
	if err != nil {
		if err.Error() != "no active connection" {
			return "Error"
		}
	}

	return "NotConnected"
}

func (m *Module) getDnsStatus() string {
	resolv, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return "Error"
	}

	dnsAddr := m.config.IpAddresses["dns"].IpAddress

	if strings.Contains(string(resolv), "nameserver "+dnsAddr) {
		return "Local"
	}

	return "Custom"
}

func (m *Module) getFirewallStatus() string {
	cmd := exec.Command("sudo", "nft", "-j", "list", "ruleset")
	output, err := cmd.Output()
	if err != nil {
		return "Error"
	}

	var nftData Nftables
	err = json.Unmarshal(output, &nftData)
	if err != nil {
		return "Error"
	}

	chainInput, err := isFirewallChainPolicyRestricted(nftData, "input")
	if err != nil {
		return "Error"
	}
	chainForward, err := isFirewallChainPolicyRestricted(nftData, "forward")
	if err != nil {
		return "Error"
	}
	chainOutput, err := isFirewallChainPolicyRestricted(nftData, "output")
	if err != nil {
		return "Error"
	}

	if !chainInput {
		return "InputUnrestricted"
	}
	if !chainForward {
		return "ForwardUnrestricted"
	}
	if !chainOutput {
		return "OutputUnrestricted"
	}

	return "Restricted"
}
