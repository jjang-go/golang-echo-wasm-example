package multiply

import "syscall/js"

// 곱셈 함수
func Multiply(this js.Value, args []js.Value) interface{} {
	result := args[0].Float() * args[1].Float()
	return result
}
