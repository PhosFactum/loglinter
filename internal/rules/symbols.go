// Правило проверки на отсутствие специсимволов
package rules

import (
	"github.com/PhosFactum/loglinter/internal/utils"
)

// CheckSymbols возвращает true, если правило нарушено (есть спецсимволы/эмодзи)
func CheckSymbols(msg string) bool {
	return utils.HasForbiddenSymbols(msg)
}
