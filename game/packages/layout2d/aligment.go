package layout2d

type Alignment int

const (
	VCenter Alignment = 0x01
	HCenter Alignment = 0x02
	Left    Alignment = 0x04
	Right   Alignment = 0x08
	Top     Alignment = 0x10
	Bottom  Alignment = 0x20
)
