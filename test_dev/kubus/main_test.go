package main

import "testing"

var (
	kubus 			Kubus 	= Kubus{4}
	correctVolume 	float64 = 64
	correctLuas 	float64 = 96
	correctKeliling float64 = 48
)

func TestHitungVolume (t *testing.T) {
	t.Logf("Volume: %.2f", kubus.Volume())

	if kubus.Volume() != correctVolume {
		t.Errorf("SALAH! harusnya %.2f", correctVolume)
	}
}
func TestHitungLuas (t *testing.T) {
	t.Logf("Luas: %.2f", kubus.Luas())

	if kubus.Luas() != correctLuas {
		t.Errorf("SALAH! harusnya %.2f", correctLuas)
	}
}
func TestHitungKeliling (t *testing.T) {
	t.Logf("Keliling: %.2f", kubus.Keliling())

	if kubus.Keliling() != correctKeliling {
		t.Errorf("SALAH! harusnya %.2f", correctKeliling)
	}
}



