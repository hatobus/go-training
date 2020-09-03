package tempconv

import "testing"

func TestConvertKelvin(t *testing.T) {
	type testData struct {
		inputKelvin Kelvin
		expectCelsius Celsius
		expectFahrenheit Fahrenheit
	}

	testCases := map[string]testData{
		"0Kを華氏と摂氏温度に変換する": {
			inputKelvin: Kelvin(0),
			expectCelsius: Celsius(-273.15),
			expectFahrenheit: Fahrenheit(-459.67),
		},
		"273.15Kを華氏と摂氏温度に変換する": {
			inputKelvin: Kelvin(273.15),
			expectCelsius: Celsius(0),
			expectFahrenheit: Fahrenheit(31.999999999999943),
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			gotCelsius := KtoC(tc.inputKelvin)

			gotFahrenheit := KtoF(tc.inputKelvin)

			if (gotCelsius - tc.expectCelsius) != 0 {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectCelsius.String(), gotCelsius.String())
			}

			if (gotFahrenheit - tc.expectFahrenheit) != 0 {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectFahrenheit.String(), gotFahrenheit.String())
			}
		})
	}
}

func TestConvertCelsius(t *testing.T) {
	type testData struct {
		inputCelsius Celsius
		expectKelvin Kelvin
		expectFahrenheit Fahrenheit
	}

	testCases := map[string]testData{
		"0℃をケルビンと摂氏温度に変換する": {
			inputCelsius: Celsius(0),
			expectKelvin: Kelvin(273.15),
			expectFahrenheit: Fahrenheit(32),
		},
		"-273.15℃をケルビンと摂氏温度に変換する": {
			inputCelsius: Celsius(-273.15),
			expectKelvin: Kelvin(0),
			expectFahrenheit: Fahrenheit(-459.66999999999996),
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			gotKelvin := CtoK(tc.inputCelsius)

			gotFahrenheit := CtoF(tc.inputCelsius)

			if (gotKelvin - tc.expectKelvin) != 0 {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectKelvin.String(), gotKelvin.String())
			}

			if (gotFahrenheit - tc.expectFahrenheit) != 0 {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectFahrenheit.String(), gotFahrenheit.String())
			}
		})
	}
}

func TestConvertFahrenheit(t *testing.T) {
	type testData struct {
		inputFahrenheit Fahrenheit
		expectCelsius Celsius
		expectKelvin Kelvin
	}

	testCases := map[string]testData{
		"0Fをケルビンと摂氏温度に変換する": {
			inputFahrenheit: Fahrenheit(0),
			expectCelsius: Celsius(-17.77777777777778),
			expectKelvin: Kelvin(255.3722222222222),
		},
		"-273.15Fをケルビンと摂氏温度に変換する": {
			inputFahrenheit: Fahrenheit(-273.15),
			expectCelsius: Celsius(-169.52777777777777),
			expectKelvin: Kelvin(103.62222222222223),
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			gotKelvin := FtoK(tc.inputFahrenheit)

			gotCelsius := FtoC(tc.inputFahrenheit)

			if gotKelvin != tc.expectKelvin {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectKelvin.String(), gotKelvin.String())
			}

			if gotCelsius != tc.expectCelsius {
				t.Fatalf("unexpected output, want: %v, but %v", tc.expectCelsius.String(), gotCelsius.String())
			}
		})
	}
}
