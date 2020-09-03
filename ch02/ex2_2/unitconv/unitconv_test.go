package unitconv

import (
	"testing"
)

func TestTemperature(t *testing.T) {
	// 摂氏から華氏へ変換できることのテスト
	inputC := AbsoluteZeroC
	gotK := CtoK(inputC)
	wantK := Kelvin(0)

	if gotK != wantK {
		t.Fatalf("CtoK(%s) want %s, but got %v", inputC, wantK, gotK)
	}

	// 絶対温度から摂氏へ変換できることのテスト
	inputK := Kelvin(0)
	gotC := KtoC(inputK)
	wantC := AbsoluteZeroC

	if gotC != wantC {
		t.Fatalf("KtoC(%s) want %s, but got %v", inputK, wantC, gotC)
	}
}

func TestLenth(t *testing.T) {
	// 1m から 1ft へ変換できることのテスト
	inputM := Metre(1)
	gotF := MToF(inputM)
	wantF := Feet(3.280839895013123)

	if gotF != wantF {
		t.Fatalf("MToF(%v) want %v but got %v", inputM, wantF, gotF)
	}

	// 1ft から 1m へ変換できることのテスト
	inputF := Feet(1)
	gotM := FToM(inputF)
	wantM := Metre(0.3048)

	if gotM != wantM {
		t.Fatalf("FToM(%v) want %v but got %v", inputF, wantM, gotM)
	}
}

func TestWeight(t *testing.T) {
	// 1kgから1pdに変換できることのテスト
	inputKG := Kilogram(1)
	gotP := KToP(inputKG)
	wantP := Pound(2.20462)

	if gotP != wantP {
		t.Fatalf("KToP(%v) want %v but got %v", inputKG, wantP, gotP)
	}

	// 1pdから1kgに変換できることのテスト
	inputP := Pound(1)
	gotK := PToK(inputP)
	wantK := Kilogram(0.45359290943563974)

	if gotK != wantK {
		t.Fatalf("PToK(%v) want %v but got %v", inputP, wantK, gotK)
	}
}
