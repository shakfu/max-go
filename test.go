package main

import "C"
import "unsafe"

type foo struct{}

func test() {
	// basics
	Log("logging...")
	Error("errored...")
	Alert("alerting...")

	// create class
	class := NewClass("maxgo", func(object Object) unsafe.Pointer {
		Pretty("init", object)

		object.IntIn()
		object.AnyIn()
		object.IntOut()

		o := unsafe.Pointer(&foo{})
		Pretty(o)

		return o
	}, func(ptr unsafe.Pointer, msg string, inlet int, atoms []Atom) {
		Pretty("handler", ptr, msg, inlet, atoms)
	}, func(pointer unsafe.Pointer) {
		Pretty("free", pointer)
	})

	// add methods
	class.AddMethod("in1") // TODO: Relay to bridge_int
	class.AddMethod("in2") // TODO: Relay to bridge_gimme

	// register
	class.Register()
}
