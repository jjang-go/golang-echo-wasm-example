package subtract

import "syscall/js"

// 뺄셈 함수
func Subtract(this js.Value, args []js.Value) interface{} {
	result := args[0].Float() - args[1].Float()
	return result
}
