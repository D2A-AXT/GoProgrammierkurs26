package musterloesungen

// 07. Temperatur bewerten

func TemperatureLabel(temp int) string {
    if temp >= 25 {
        return "warm"
    } else {
        return "kühl"
    }
}
