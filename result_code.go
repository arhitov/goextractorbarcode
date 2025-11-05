package goextractorbarcode

import "github.com/arhitov/goextractorbarcode/types"

type ResultCode struct {
	format types.CodeFormat
	text   string
}

func NewDataCode(
	format types.CodeFormat,
	text string,
) *ResultCode {
	return &ResultCode{
		format: format,
		text:   text,
	}
}

func (r ResultCode) Format() types.CodeFormat {
	return r.format
}

func (r ResultCode) Text() string {
	return r.text
}
