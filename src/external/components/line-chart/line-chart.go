package linechart

// Chart Graphique chart.js
type Chart struct {
	ElementID string        `json:"elementID"` // ID de la balise HTML à laquelle rattacher le graphique
	Data      Configuration `json:"data"`      // Configuration du graphique
}

// Configuration Chart.js integration
type Configuration struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
	// Options []any  `json:"options"`
	// Plugins []any  `json:"plugins`
}

// Data Chart.js integration
type Data struct {
	Datasets []Dataset `json:"datasets"` // Jeux de données à afficher
	Labels   []string  `json:"labels"`   // Labels axe X
}

type Dataset struct {
	Label       string    `json:"label"`           // Légende
	Data        []float64 `json:"data"`            // Points
	Fill        bool      `json:"fill"`            // Remplissage des points
	BorderColor string    `json:"backgroundColor"` // Couleur du tracé
	Tension     float64   `json:"tension"`         // 0 -> droites entre les points, 1 -> courbes
}
