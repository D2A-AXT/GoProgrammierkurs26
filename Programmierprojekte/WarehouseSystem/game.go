package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 1000
	WindowHeight = 650
)

type Game struct {
	Warehouse Warehouse
	Selected  int
	Message   string

	demoProductAdded bool
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func initialWarehouse() Warehouse {
	return Warehouse{
		Products: []Product{
			{ID: "P-100", Name: "Kabelbinder 200mm", Quantity: 86, Minimum: 20, PriceCents: 9},
			{ID: "P-110", Name: "RJ45-Stecker", Quantity: 42, Minimum: 15, PriceCents: 39},
			{ID: "P-120", Name: "Netzwerkkabel 5m", Quantity: 11, Minimum: 8, PriceCents: 849},
			{ID: "P-130", Name: "Sicherung 2A", Quantity: 4, Minimum: 6, PriceCents: 115},
			{ID: "P-140", Name: "Relais 24V", Quantity: 8, Minimum: 5, PriceCents: 1290},
			{ID: "P-150", Name: "Sensor M12", Quantity: 3, Minimum: 4, PriceCents: 2490},
			{ID: "P-160", Name: "Patchpanel 24-Port", Quantity: 6, Minimum: 2, PriceCents: 4990},
			{ID: "P-170", Name: "Hutschienenklemme", Quantity: 125, Minimum: 40, PriceCents: 31},
		},
	}
}

func (g *Game) Reset() {
	g.Warehouse = initialWarehouse()
	g.Selected = 0
	g.Message = "Wähle einen Artikel und führe eine Lagerbewegung aus."
	g.demoProductAdded = false
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	if len(g.Warehouse.Products) == 0 {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.Selected--
		if g.Selected < 0 {
			g.Selected = len(g.Warehouse.Products) - 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.Selected++
		if g.Selected >= len(g.Warehouse.Products) {
			g.Selected = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.stockIn()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.stockOut()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.addDemoProduct()
	}

	return nil
}

func (g *Game) selectedProduct() *Product {
	if g.Selected < 0 || g.Selected >= len(g.Warehouse.Products) {
		return nil
	}

	id := g.Warehouse.Products[g.Selected].ID
	return FindProduct(&g.Warehouse, id)
}

func (g *Game) stockIn() {
	product := g.selectedProduct()
	if product == nil {
		g.Message = "Produkt konnte nicht gefunden werden."
		return
	}

	if AddStock(product, 10) {
		g.Message = fmt.Sprintf("%s: 10 Stück eingelagert.", product.Name)
	} else {
		g.Message = "Einlagerung fehlgeschlagen."
	}
}

func (g *Game) stockOut() {
	product := g.selectedProduct()
	if product == nil {
		g.Message = "Produkt konnte nicht gefunden werden."
		return
	}

	if RemoveStock(product, 1) {
		g.Message = fmt.Sprintf("%s: 1 Stück ausgelagert.", product.Name)
	} else {
		g.Message = fmt.Sprintf("%s: Auslagerung nicht möglich.", product.Name)
	}
}

func (g *Game) addDemoProduct() {
	if g.demoProductAdded {
		g.Message = "Die Demo-Neuanlage wurde bereits ausgeführt."
		return
	}

	product := Product{
		ID:         "P-500",
		Name:       "Industrie-Switch",
		Quantity:   2,
		Minimum:    1,
		PriceCents: 18900,
	}

	if AddProduct(&g.Warehouse, product) {
		g.demoProductAdded = true
		g.Selected = len(g.Warehouse.Products) - 1
		g.Message = "Neuer Artikel angelegt: Industrie-Switch."
	} else {
		g.Message = "Artikel konnte nicht angelegt werden."
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}
