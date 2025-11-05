package types

type CodeFormat string

const (
	CodeFormatDataMatrix    CodeFormat = "datamatrix"
	CodeFormatGS1DataMatrix CodeFormat = "gs1datamatrix"
	CodeFormatQRCode        CodeFormat = "qrcode"
)
