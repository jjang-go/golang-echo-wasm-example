package divide

import "syscall/js"

// 나눗셈 함수
func Divide(this js.Value, args []js.Value) interface{} {
	result := args[0].Float() / args[1].Float()
	return result
}
