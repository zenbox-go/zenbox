package print

import "github.com/labstack/gommon/color"

func D(args ...interface{}) {
	color.Println(formatColor(0, args...)...)
}

func I(args ...interface{}) {
	color.Println(formatColor(1, args...)...)
}

func E(args ...interface{}) {
	color.Println(formatColor(2, args...)...)
}

func W(args ...interface{}) {
	color.Println(formatColor(3, args...)...)
}

func DF(format string, args ...interface{}) {
	color.Printf(color.Magenta("[调试] ", color.B)+format+"\n", args...)
}

func IF(format string, args ...interface{}) {
	color.Printf(color.Green("[提示] ", color.B)+format+"\n", args...)
}

func EF(format string, args ...interface{}) {
	color.Printf(color.Red("[错误] ", color.B)+format+"\n", args...)
}

func WF(format string, args ...interface{}) {
	color.Printf(color.Yellow("[警告] ", color.B)+format+"\n", args...)
}

func formatColor(typ int, args ...interface{}) []interface{} {
	texts := make([]interface{}, 0)

	switch typ {
	case 0:
		texts = append(texts, color.Magenta("[调试]", color.B))
	case 1:
		texts = append(texts, color.Green("[提示]", color.B))
	case 2:
		texts = append(texts, color.Red("[错误]", color.B))
	case 3:
		texts = append(texts, color.Yellow("[警告]", color.B))
	default:

	}

	return append(texts, args...)
}
