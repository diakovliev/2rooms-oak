package layout2d

// Alignment is the alignment of an entity
type Alignment int

const (
	// VCenter is the vertical center alignment
	VCenter Alignment = 0x01
	// HCenter is the horizontal center alignment
	HCenter Alignment = 0x02
	// Left is the left alignment
	Left Alignment = 0x04
	// Right is the right alignment
	Right Alignment = 0x08
	// Top is the top alignment
	Top Alignment = 0x10
	// Bottom is the bottom alignment
	Bottom Alignment = 0x20
)
