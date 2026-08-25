package grundrechenarten

import "testing"

func TestBasicOperations(t *testing.T) {
    a := 10
    b := 3

    wantSum := 13
    wantDifference := 7
    wantProduct := 30

    gotSum, gotDifference, gotProduct := BasicOperations(a, b)

    if gotSum != wantSum || gotDifference != wantDifference || gotProduct != wantProduct {
        t.Errorf("bekommen: (%d, %d, %d); erwartet: (%d, %d, %d)",
            gotSum, gotDifference, gotProduct,
            wantSum, wantDifference, wantProduct)
    }
}
