// Правило проверки на отсутствие чувствительных данных
package rules

import (
	"regexp"
)

// Компиляция регулярного выражения с чувствительными данными
var sensitiveRegex = regexp.MustCompile(`(?i)\b(password|passwd|pass|api[\s_-]?key|apikey|token|secret|credential)\b`)

// CheckSensitive возвращает true, если найдены чувствительные данные
func CheckSensitive(msg string) bool {
	return sensitiveRegex.MatchString(msg)
}
