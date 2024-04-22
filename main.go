package main

import (
	"errors"
	"fmt"
)

type PaymentSystem struct {
	Name    string
	Minimum float64
	Credit  bool
}

var PaymentSystems = []PaymentSystem{
	{"cod", 0, false},
	{"transfer", 0, false},
	{"debit", 0, false},
	{"credit", 500000, true},
	{"gerai", 0, false},
}

func findPaymentSystem(name string) *PaymentSystem {
	for _, ps := range PaymentSystems {
		if ps.Name == name {
			return &ps
		}
	}
	return nil
}

func PembayaranBarang(ht float64, m string, cicil bool) error {
	if ht <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	ps := findPaymentSystem(m)
	if ps == nil {
		return errors.New("metode tidak dikenali")
	}

	if cicil && !ps.Credit {
		return errors.New("credit harus dicicil")
	}

	if !cicil && ps.Credit {
		return errors.New("credit harus dicicil")
	}

	if cicil && ht <= ps.Minimum {
		return errors.New("cicilan tidak memenuhi syarat")
	}

	return nil
}

func main() {
	// Contoh penggunaan
	err := PembayaranBarang(600000, "credit", true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil.")
	}
}
