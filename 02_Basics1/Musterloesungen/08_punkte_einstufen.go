package musterloesungen

// 08. Punkte einstufen

func ClassifyPoints(points int) string {
    if points >= 90 {
        return "sehr gut"
    } else if points >= 75 {
        return "gut"
    } else if points >= 50 {
        return "bestanden"
    } else {
        return "nicht bestanden"
    }
}
