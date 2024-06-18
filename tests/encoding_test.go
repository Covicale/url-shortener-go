package encoding_test

import (
	"testing"

	"github.com/covicale/url-shortener-go/internal/api/utils"
)

func TestEncodeBase62(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {
		expected := "1"
		got := utils.EncodeToBase62(1)
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Test 10", func(t *testing.T) {
		expected := "a"
		got := utils.EncodeToBase62(10)
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Test 100000000000", func(t *testing.T) {
		expected := "1L9zO9O"
		got := utils.EncodeToBase62(100000000000)
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}

func TestDecodeBase62(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		expected := 1
		got := utils.DecodeFromBase62("1")
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Test a", func(t *testing.T) {
		expected := 10
		got := utils.DecodeFromBase62("a")
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Test 1L9zO9O", func(t *testing.T) {
		expected := 100000000000
		got := utils.DecodeFromBase62("1L9zO9O")
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
