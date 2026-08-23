package account

import "testing"

func TestDepositCopy(t *testing.T) {
	original := Account{Owner: "Ada", Balance: 100}
	changed := DepositCopy(original, 50)

	if changed.Balance != 150 {
		t.Errorf("changed.Balance = %d, want 150", changed.Balance)
	}
	if original.Balance != 100 {
		t.Errorf("original.Balance = %d, want unverändert 100", original.Balance)
	}
}

func TestRenameCopy(t *testing.T) {
	original := Account{Owner: "Ada", Balance: 100}
	changed := RenameCopy(original, "Grace")

	if changed.Owner != "Grace" {
		t.Errorf("changed.Owner = %q, want Grace", changed.Owner)
	}
	if original.Owner != "Ada" {
		t.Errorf("original.Owner = %q, want unverändert Ada", original.Owner)
	}
}
