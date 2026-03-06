// Тесты для 4-х базовых правил
package rules

import (
	"testing"
)

// TestCheckLowercase проверяет правило низкого регистра
func TestCheckLowercase(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool // true = нарушение
	}{
		{"ok", "starting server", false},
		{"fail", "Starting server", true},
		{"empty", "", false},
		{"number start", "123 started", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLowercase(tt.msg); got != tt.want {
				t.Errorf("CheckLowercase(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

// TestCheckLanguage проверяет правило соответствия латинскому алфавиту
func TestCheckLanguage(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"ok", "hello world", false},
		{"fail cyrillic", "привет мир", true},
		{"fail chinese", "你好", true},
		{"ok with numbers", "port 8080", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLanguage(tt.msg); got != tt.want {
				t.Errorf("CheckLanguage(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

// TestCheckSymbols проверяет правило отсутствующих специсимволов
func TestCheckSymbols(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"ok", "server started", false},
		{"fail exclamation", "started!", true},
		{"fail question", "failed?", true},
		{"fail emoji", "started 🚀", true},
		{"ok dots", "loading...", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSymbols(tt.msg); got != tt.want {
				t.Errorf("CheckSymbols(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

// TestCheckSensitive проверяет правило отсутствующих чувствительных данных
func TestCheckSensitive(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"ok", "user authenticated", false},
		{"fail password", "password: 12345", true},
		{"fail token", "api_key=abc", true},
		{"fail case insensitive", "TOKEN in header", true},
		{"fail secret", "secret value", true},
		{"ok key word", "keyboard", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSensitive(tt.msg); got != tt.want {
				t.Errorf("CheckSensitive(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
