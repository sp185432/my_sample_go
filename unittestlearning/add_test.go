package addition

import "testing"

func TestAdd(t *testing.T) {
	/*
		actual := Add(4, 8)
		expected := 12

		if actual != expected {
			t.Fatal("FAIL: Actual value is different than expected.")
			t.Fatalf("FAIL: Addition test results: \nExpected: %d \nActual: %d", expected, actual)
		}
		t.Log("PASS: Addition result is correct.")
		t.Logf("PASS: Expected: %d \nActual %d", expected, actual)
	*/

	for _, tc := range testCases {
		actual := Add(tc.num1, tc.num2)
		expected := tc.num1 + tc.num2

		if actual != expected {
			t.Fatal("FAIL: Actual value is different than expected.")
			t.Fatalf("FAIL: Addition test results: \nExpected: %d \nActual: %d", expected, actual)
		}
		t.Log("PASS: Addition result is correct.")
		t.Logf("PASS: Expected: %d \nActual %d", expected, actual)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 8)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(4, 8)
	}
}
