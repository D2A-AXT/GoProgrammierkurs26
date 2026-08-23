package main

func CreateWorld() []Room {
	return []Room{
		{
			ID:          "office",
			Name:        "Verlassenes Büro",
			Description: "Ein alter Schreibtisch, ein Aktenschrank und eine Tür zum Flur.",
			Items: []Item{
				{
					ID:          "drawer_key",
					Name:        "Kleiner Schlüssel",
					Description: "Ein kleiner Messingschlüssel. Vielleicht passt er irgendwo.",
				},
			},
			Exits: []Exit{
				{
					Name:       "Flur",
					TargetRoom: "hallway",
				},
			},
		},
		{
			ID:          "hallway",
			Name:        "Dunkler Flur",
			Description: "Links ist ein Lagerraum. Geradeaus befindet sich eine Stahltür.",
			Items:       nil,
			Exits: []Exit{
				{
					Name:       "Büro",
					TargetRoom: "office",
				},
				{
					Name:       "Lager",
					TargetRoom: "storage",
				},
				{
					Name:         "Stahltür",
					TargetRoom:   "control",
					RequiredItem: "access_card",
					LockedText:   "Die Stahltür verlangt eine Zugangskarte.",
				},
			},
		},
		{
			ID:          "storage",
			Name:        "Lagerraum",
			Description: "Regale voller alter Geräte. Auf einer Kiste liegt eine Zugangskarte.",
			Items: []Item{
				{
					ID:          "access_card",
					Name:        "Zugangskarte",
					Description: "Eine alte Zugangskarte mit Firmenlogo.",
				},
			},
			Exits: []Exit{
				{
					Name:       "Flur",
					TargetRoom: "hallway",
				},
			},
		},
		{
			ID:          "control",
			Name:        "Kontrollraum",
			Description: "Ein Bedienpult blinkt. Neben der Außentür befindet sich ein Schlüsselschalter.",
			Items:       nil,
			Exits: []Exit{
				{
					Name:       "Flur",
					TargetRoom: "hallway",
				},
				{
					Name:         "Außentür",
					TargetRoom:   "outside",
					RequiredItem: "drawer_key",
					LockedText:   "Der Schlüsselschalter ist verriegelt.",
				},
			},
		},
		{
			ID:          "outside",
			Name:        "Draußen",
			Description: "Frische Luft. Du hast den Escape Room verlassen.",
		},
	}
}
