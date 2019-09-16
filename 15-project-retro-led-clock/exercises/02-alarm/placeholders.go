package main

type placeholder [5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var digits = [...]placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

var a = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}
var l = placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}
var r = placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var m = placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}
var bang = placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}
var empty = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var alarm = [...]placeholder{
	empty, a, l, a, r, m, bang, empty,
}
