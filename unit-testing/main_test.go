package main

import "testing"

var (
	kubus        Kubus   = Kubus{4}
	realVolume   float64 = 64
	realLuas     float64 = 96
	realKeliling float64 = 48
)

/*
Aturan file unit testing
. nama file sama dengan nama file yang akan di-test, dengan akhiran _test.go
. unit test ditulis dalam bentuk method, dengan menyertakan parameter *testing.T
. nama method untuk unit testing diawali dengan Test (misal: TestHitungFungsi)
*/

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume: %.2f", kubus.Volume())

	if kubus.Volume() != realVolume {
		t.Errorf("Volume Salah!!! correct: %.2f", realVolume)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas: %.2f", kubus.Luas())

	if kubus.Luas() != realLuas {
		t.Errorf("Luas Salah!!! correct: %.2f", realLuas)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling: %.2f", kubus.Keliling())

	if kubus.Keliling() != realKeliling {
		t.Errorf("Keliling Salah!!! correct: %.2f", realKeliling)
	}
}
