package main

import  "github.com/256dpi/max-go"

type instance struct {
	in1   *max.Inlet
	in2   *max.Inlet
	out1  *max.Outlet
	out2  *max.Outlet
}

func (i *instance) Init(obj *max.Object, args []max.Atom) bool {
	// print to Max console
	max.Pretty("init", args)

	// declare inlets
	i.in1 = obj.Inlet(max.Any, "example1 inlet 1", true)
	i.in2 = obj.Inlet(max.Float, "example1 inlet 2", false)

	// declare outlets
	i.out1 = obj.Outlet(max.Any, "example outlet 1")
	i.out2 = obj.Outlet(max.Bang, "example1 outlet 2")

	return true
}

func (i *instance) Handle(inlet int, msg string, data []max.Atom) {
	// print to Max console
	max.Pretty("handle", inlet, msg, data)

	// send to first outlet
	i.out1.Any(msg, data)
}

func (i *instance) Free() {
	// print to Max console
	max.Pretty("free")
}

func init() {
	// initialize Max class
	max.Register("example1", &instance{})
}

func main() {
	// not called
}

