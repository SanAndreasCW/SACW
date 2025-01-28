package colors

type InfixColor string

type Color struct {
	Hex   uint
	Infix InfixColor
}

var (
	WhiteColor = Color{
		Hex:   0xFFFFFFFF,
		Infix: "{ffffff}",
	}
	InfoColor = Color{
		Hex:   0xf1c40fFF,
		Infix: "{f1c40f}",
	}
	NoteColor = Color{
		Hex:   0xf39c12FF,
		Infix: "{f39c12}",
	}
	SuccessColor = Color{
		Hex:   0x2ecc71FF,
		Infix: "{2ecc71}",
	}
	ErrorColor = Color{
		Hex:   0xe74c3cFF,
		Infix: "{e74c3c}",
	}
	NoticeColor = Color{
		Hex:   0x3498dbFF,
		Infix: "{3498db}",
	}
)
