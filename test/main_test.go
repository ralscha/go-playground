package main

import "testing"

func TestDiscountedPrice(t *testing.T) {
	testCases := []struct {
		price           float64
		discountPercent float64
		expected        float64
		expectError     bool
		desc            string
	}{
		{100.0, 0.0, 100.0, false, "no discount"},
		{100.0, 50.0, 50.0, false, "50% discount"},
		{100.0, 100.0, 0.0, false, "100% discount"},
		{100.0, 110.0, 0.0, true, "discount greater than 100%"},
		{100.0, -10.0, 0.0, true, "negative discount"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := DiscountedPrice(
				tc.price, tc.discountPercent)
			if tc.expectError && err == nil {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) "+
						"should return an error",
					tc.price, tc.discountPercent)
			}
			if !tc.expectError && err != nil {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) "+
						"returned an error: %v",
					tc.price, tc.discountPercent, err)
			}
			if !tc.expectError && result != tc.expected {
				t.Errorf(
					"DiscountedPrice(%.2f, %.2f) = "+
						"%.2f; want %.2f",
					tc.price, tc.discountPercent, result,
					tc.expected)
			}
		})
	}
}
