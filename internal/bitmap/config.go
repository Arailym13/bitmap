package bitmap

type BMPHeader struct {
	FileType   [2]byte
	FileSize   uint32
	Reserved1  uint16
	Reserved2  uint16
	DataOffset uint32
}

type DIBHeader struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XpixelsPerM     uint32
	YpixelsPerM     uint32
	ColorsUsed      uint32
	ImportantColors uint32
}
