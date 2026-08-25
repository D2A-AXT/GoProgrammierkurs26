package temperatur_bewerten

// Temperatur bewerten
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func TemperatureLabel(temp int) string {
	// TODO: Entscheide mit if/else zwischen "warm" und "kühl".
	if temp >= 25 {
		return "warm"
	} else {
		return "kühl"
	}
}
