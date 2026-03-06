// Правило проверки на регистр (с малой буквы) в логе
package rules

import (
	"github.com/PhosFactum/loglinter/internal/utils"
)

// CheckLowercase возвращает true, если правило нарушено
func CheckLowercase(msg string) bool {
	return !utils.IsLowercase(msg)
}
