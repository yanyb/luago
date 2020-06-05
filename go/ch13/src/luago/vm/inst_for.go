package vm

import "luago/api"

func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sBx)
}

func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a + 2)
	vm.PushValue(a)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	isPositionStep := vm.ToNumber(a+2) >= 0
	if isPositionStep && vm.Compare(a, a+1, api.LUA_OPLE) ||
		!isPositionStep && vm.Compare(a+1, a, api.LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a+3)
	}
}

func tForLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1
	if !vm.IsNil(a + 1) {
		vm.Copy(a+1, a)
		vm.AddPC(sBx)
	}
}
