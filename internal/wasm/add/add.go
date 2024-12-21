package add

import "syscall/js"

// 덧셈 함수
func Add(this js.Value, args []js.Value) interface{} {
	result := args[0].Float() + args[1].Float()
	return result
}
