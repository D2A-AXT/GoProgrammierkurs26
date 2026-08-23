package main

import "testing"

func sampleWarehouse() Warehouse {
	return Warehouse{
		Products: []Product{
			{
				ID:         "P-100",
				Name:       "Kabelbinder",
				Quantity:   50,
				Minimum:    10,
				PriceCents: 8,
			},
			{
				ID:         "P-200",
				Name:       "Relais 24V",
				Quantity:   7,
				Minimum:    5,
				PriceCents: 1290,
			},
			{
				ID:         "P-300",
				Name:       "Sensor M12",
				Quantity:   3,
				Minimum:    4,
				PriceCents: 2490,
			},
		},
	}
}

func TestFindProductIndex(t *testing.T) {
	warehouse := sampleWarehouse()

	if got := FindProductIndex(warehouse, "P-200"); got != 1 {
		t.Fatalf("index = %d, want 1", got)
	}

	if got := FindProductIndex(warehouse, "UNKNOWN"); got != -1 {
		t.Fatalf("index = %d, want -1", got)
	}
}

func TestFindProduct(t *testing.T) {
	warehouse := sampleWarehouse()

	product := FindProduct(&warehouse, "P-300")

	if product == nil {
		t.Fatal("expected product, got nil")
	}

	if product.Name != "Sensor M12" {
		t.Fatalf("name = %q, want %q", product.Name, "Sensor M12")
	}
}

func TestFindProductReturnsPointerIntoWarehouse(t *testing.T) {
	warehouse := sampleWarehouse()

	product := FindProduct(&warehouse, "P-100")
	if product == nil {
		t.Fatal("expected product, got nil")
	}

	product.Quantity = 999

	if warehouse.Products[0].Quantity != 999 {
		t.Fatal("returned pointer must modify the product inside the warehouse")
	}
}

func TestFindProductMissing(t *testing.T) {
	warehouse := sampleWarehouse()

	if product := FindProduct(&warehouse, "UNKNOWN"); product != nil {
		t.Fatalf("expected nil, got %+v", product)
	}
}

func TestAddStock(t *testing.T) {
	product := Product{Quantity: 5}

	if !AddStock(&product, 10) {
		t.Fatal("expected AddStock to succeed")
	}

	if product.Quantity != 15 {
		t.Fatalf("quantity = %d, want 15", product.Quantity)
	}
}

func TestAddStockRejectsInvalidAmount(t *testing.T) {
	product := Product{Quantity: 5}

	if AddStock(&product, 0) {
		t.Fatal("amount 0 must be rejected")
	}

	if AddStock(&product, -3) {
		t.Fatal("negative amount must be rejected")
	}

	if product.Quantity != 5 {
		t.Fatalf("quantity changed to %d", product.Quantity)
	}
}

func TestRemoveStock(t *testing.T) {
	product := Product{Quantity: 20}

	if !RemoveStock(&product, 7) {
		t.Fatal("expected RemoveStock to succeed")
	}

	if product.Quantity != 13 {
		t.Fatalf("quantity = %d, want 13", product.Quantity)
	}
}

func TestRemoveStockRejectsTooMuch(t *testing.T) {
	product := Product{Quantity: 4}

	if RemoveStock(&product, 5) {
		t.Fatal("removing more than available must fail")
	}

	if product.Quantity != 4 {
		t.Fatalf("quantity changed to %d", product.Quantity)
	}
}

func TestRemoveStockRejectsInvalidAmount(t *testing.T) {
	product := Product{Quantity: 4}

	if RemoveStock(&product, 0) {
		t.Fatal("amount 0 must fail")
	}

	if RemoveStock(&product, -1) {
		t.Fatal("negative amount must fail")
	}
}

func TestIsLowStock(t *testing.T) {
	if !IsLowStock(Product{Quantity: 3, Minimum: 5}) {
		t.Error("quantity below minimum must be low stock")
	}

	if !IsLowStock(Product{Quantity: 5, Minimum: 5}) {
		t.Error("quantity equal to minimum must be low stock")
	}

	if IsLowStock(Product{Quantity: 6, Minimum: 5}) {
		t.Error("quantity above minimum must not be low stock")
	}
}

func TestTotalQuantity(t *testing.T) {
	warehouse := sampleWarehouse()

	if got := TotalQuantity(warehouse); got != 60 {
		t.Fatalf("total quantity = %d, want 60", got)
	}
}

func TestInventoryValueCents(t *testing.T) {
	warehouse := sampleWarehouse()

	// 50*8 + 7*1290 + 3*2490 = 16900
	const want = 16900

	if got := InventoryValueCents(warehouse); got != want {
		t.Fatalf("inventory value = %d, want %d", got, want)
	}
}

func TestAddProduct(t *testing.T) {
	warehouse := sampleWarehouse()

	product := Product{
		ID:         "P-400",
		Name:       "Patchpanel",
		Quantity:   2,
		Minimum:    1,
		PriceCents: 4990,
	}

	if !AddProduct(&warehouse, product) {
		t.Fatal("expected AddProduct to succeed")
	}

	if len(warehouse.Products) != 4 {
		t.Fatalf("len = %d, want 4", len(warehouse.Products))
	}

	if warehouse.Products[3].ID != "P-400" {
		t.Fatalf("last product ID = %q", warehouse.Products[3].ID)
	}
}

func TestAddProductRejectsDuplicateID(t *testing.T) {
	warehouse := sampleWarehouse()

	duplicate := Product{
		ID:         "P-100",
		Name:       "Anderes Produkt",
		Quantity:   1,
		Minimum:    0,
		PriceCents: 100,
	}

	if AddProduct(&warehouse, duplicate) {
		t.Fatal("duplicate ID must be rejected")
	}

	if len(warehouse.Products) != 3 {
		t.Fatalf("warehouse length changed to %d", len(warehouse.Products))
	}
}

func TestAddProductRejectsInvalidData(t *testing.T) {
	tests := []Product{
		{ID: "", Name: "A", Quantity: 1, Minimum: 0, PriceCents: 100},
		{ID: "X", Name: "", Quantity: 1, Minimum: 0, PriceCents: 100},
		{ID: "X", Name: "A", Quantity: -1, Minimum: 0, PriceCents: 100},
		{ID: "X", Name: "A", Quantity: 1, Minimum: -1, PriceCents: 100},
		{ID: "X", Name: "A", Quantity: 1, Minimum: 0, PriceCents: -1},
	}

	for i, product := range tests {
		warehouse := sampleWarehouse()

		if AddProduct(&warehouse, product) {
			t.Fatalf("case %d: invalid product was accepted", i)
		}
	}
}
