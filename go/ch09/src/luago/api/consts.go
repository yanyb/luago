package api

type LuaType int
type ArithOp = int
type CompareOp = int

const (
	LUA_MINSTACK            = 20
	LUAI_MAXSTACK           = 1000000
	LUA_REGISTRYINDEX       = -LUAI_MAXSTACK - 1000
	LUA_RIDX_GLOBALS  int64 = 2
)

const (
	LUA_TNONE = iota - 1
	LUA_TNIL
	LUA_TBOOLEAN
	LUA_TLIGHTUSERDATA
	LUA_TNUMBER
	LUA_TSTRING
	LUA_TTABLE
	LUA_TFUNCTION
	LUA_TUSERDATA
	LUA_TTHREAD
)

const (
	LUA_OPADD  = iota // +
	LUA_OPSUB         //-
	LUA_OPMUL         //*
	LUA_OPMOD         //%
	LUA_OPPOW         //^
	LUA_OPDIV         // /
	LUA_OPIDIV        // //
	LUA_OPBAND        //&
	LUA_OPBOR         //|
	LUA_OPBXOR        //~
	LUA_OPSHL         // <<
	LUA_OPSHR         // >>
	LUA_OPUNM         // -(unary minus)
	LUA_OPBNOT        //~
)

const (
	LUA_OPEQ = iota // ==
	LUA_OPLT        // <
	LUA_OPLE        // <=
)