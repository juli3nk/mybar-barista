package kbd

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type InputDevice struct {
	Type                string `json:"type"`
	XkbActiveLayoutName string `json:"xkb_active_layout_name"`
}

func getKeyboardLayout() (string, error) {
	// Exécute la commande swaymsg pour récupérer les périphériques d'entrée
	cmd := exec.Command("swaymsg", "-t", "get_inputs")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'exécution de swaymsg: %w", err)
	}

	// Parse le JSON
	var devices []InputDevice
	if err := json.Unmarshal(output, &devices); err != nil {
		return "", fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}

	// Recherche le premier clavier trouvé
	for _, device := range devices {
		if device.Type == "keyboard" {
			// Extraire le code (ex: "French (AZERTY)" → "fr")
			code := extractLayoutCode(device.XkbActiveLayoutName)
			return code, nil
		}
	}

	return "", fmt.Errorf("aucun clavier trouvé")
}

// Fonction pour extraire le code du layout
func extractLayoutCode(layout string) string {
	// Liste des correspondances de noms à codes (ajoute d'autres si besoin)
	mapping := map[string]string{
		"English (US)":    "us",
		"French (AZERTY)": "fr",
		"German":          "de",
		"Spanish":         "es",
		"Italian":         "it",
		"Portuguese (BR)": "br",
		"Russian":         "ru",
	}

	// Vérifier si le nom existe dans la table de correspondance
	if code, found := mapping[layout]; found {
		return code
	}

	// Si inconnu, retourner une version "brute" du nom en minuscule
	return strings.ToLower(strings.Split(layout, " ")[0])
}
