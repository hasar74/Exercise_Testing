package main

import (
	"errors"
	"testing"
)

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		ht       float64
		m        string
		cicil    bool
		expected error
	}{
		{600000, "credit", true, nil}, // Case valid
		{0, "cod", false, errors.New("harga tidak bisa nol")},
		{1000000, "debit", true, errors.New("credit harus dicicil")},
		{400000, "credit", true, errors.New("cicilan tidak memenuhi syarat")},
		{200000, "unknown", false, errors.New("metode tidak dikenali")},
	}

	for _, test := range tests {
		err := PembayaranBarang(test.ht, test.m, test.cicil)
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err != nil && err.Error() != test.expected.Error()) {
			t.Errorf("PembayaranBarang(%f, %s, %t) returned %v, expected %v", test.ht, test.m, test.cicil, err, test.expected)
		}
	}
}
