// Тесты для вспомогательных функций проверки условий
package utils

import (
	"testing"
)

// TestIsLowercase проверяет, начинается ли строка с маленькой буквы
func TestIsLowercase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"ok", "hello", true},
		{"fail", "Hello", false},
		{"empty", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowercase(tt.s); got != tt.want {
				t.Errorf("IsLowercase(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// TestIsEnglishOnly проверяет, содержит ли строка только английские символы (ASCII)
func TestIsEnglishOnly(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"ok", "hello", true},
		{"fail", "привет", false},
		{"ok numbers", "port8080", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnglishOnly(tt.s); got != tt.want {
				t.Errorf("IsEnglishOnly(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// TestIsEmoji проверяет, распознаётся ли руна как эмодзи
func TestIsEmoji(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"rocket", '🚀', true},
		{"star", '⭐', true},
		{"letter", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmoji(tt.r); got != tt.want {
				t.Errorf("IsEmoji(%U) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}

// TestHasForbiddenSymbols проверяет наличие запрещённых символов (!, ?, …) или эмодзи
func TestHasForbiddenSymbols(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"ok", "hello", false},
		{"fail !", "hello!", true},
		{"fail ?", "hello?", true},
		{"fail emoji", "hello🚀", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasForbiddenSymbols(tt.s); got != tt.want {
				t.Errorf("HasForbiddenSymbols(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
