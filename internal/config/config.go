// Конфигурация линтера (через кфг-файл)
package config

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *Config {
	return &Config{
		Rules: Rules{
			Lowercase: Rule{Enabled: true},
			Language:  Rule{Enabled: true},
			Symbols:   Rule{Enabled: true},
			Sensitive: Rule{Enabled: true},
		},
		Sensitive: SensitiveConfig{
			Keywords: []string{
				"password", "passwd", "pass",
				"api_key", "apikey", "api-key",
				"token", "secret", "credential",
			},
		},
		Symbols: SymbolsConfig{
			Forbidden: []string{"!", "?", "…"},
		},
	}
}

// Config представляет собой основную конфигурацию
type Config struct {
	Rules     Rules           `json:"rules" yaml:"rules"`
	Sensitive SensitiveConfig `json:"sensitive" yaml:"sensitive"`
	Symbols   SymbolsConfig   `json:"symbols" yaml:"symbols"`
}

// Rules представляет собой настройки правил
type Rules struct {
	Lowercase Rule `json:"lowercase" yaml:"lowercase"`
	Language  Rule `json:"language" yaml:"language"`
	Symbols   Rule `json:"symbols" yaml:"symbols"`
	Sensitive Rule `json:"sensitive" yaml:"sensitive"`
}

// Rule представляет собой настройку отдельного правила
type Rule struct {
	Enabled bool `json:"enabled" yaml:"enabled"`
}

// SensitiveConfig это настройка проверки чувствительных данных
type SensitiveConfig struct {
	Keywords []string `json:"keywords" yaml:"keywords"`
}

// SymbolsConfig это настройка проверки спецсимволов
type SymbolsConfig struct {
	Forbidden []string `json:"forbidden" yaml:"forbidden"`
}

// Load загружает конфигурацию из файла
func Load(path string) (*Config, error) {
	if path == "" {
		return DefaultConfig(), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := DefaultConfig()
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

// BuildSensitiveRegex строит регулярку из ключевых слов
func (c *Config) BuildSensitiveRegex() (*regexp.Regexp, error) {
	if len(c.Sensitive.Keywords) == 0 {
		return regexp.Compile(`^$`) // никогда не матчится
	}

	// Экранируем спецсимволы в ключевых словах
	escaped := make([]string, 0, len(c.Sensitive.Keywords))
	for _, kw := range c.Sensitive.Keywords {
		escaped = append(escaped, regexp.QuoteMeta(kw))
	}

	// ИСПРАВЛЕНО: strings.Join вместо несуществующего regexp.Join
	pattern := `(?i)\b(` + strings.Join(escaped, "|") + `)\b`
	return regexp.Compile(pattern)
}

// GetForbiddenRunes возвращает запрещённые символы как руны
func (c *Config) GetForbiddenRunes() []rune {
	runes := make([]rune, 0, len(c.Symbols.Forbidden))
	for _, s := range c.Symbols.Forbidden {
		for _, r := range s {
			runes = append(runes, r)
		}
	}
	return runes
}
