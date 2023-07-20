package types

type Unit struct {
	GB int
	KB int
	MB int
	PB int
	TB int
}

var SizeUnit = Unit{
	GB: 1073741824,
	KB: 1024,
	MB: 1048576,
	PB: 1125899906842624,
	TB: 1099511627776,
}
