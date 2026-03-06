// Правило проверки на латиницу в логе
package rules

import (
	"github.com/PhosFactum/loglinter/internal/utils"
)

// CheckLanguage возвращает true, если правило нарушено (есть не-английские символы)
func CheckLanguage(msg string) bool {
	return !utils.IsEnglishOnly(msg)
}
