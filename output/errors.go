package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	default:
		color.Red("Неизвестный тип ошибки")
	}
}

func sumValue[T int | float32](a, b T) T {
	return a+b
}