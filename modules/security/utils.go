package security

type Nftables struct {
	Nftables []map[string]interface{} `json:"nftables"`
}

func isFirewallChainPolicyRestricted(nftData Nftables, chainName string) (bool, error) {
	for _, obj := range nftData.Nftables {
		if chain, ok := obj["chain"].(map[string]interface{}); ok {
			if chain["name"] == chainName {
				policy := chain["policy"].(string)
				if policy == "drop" {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
