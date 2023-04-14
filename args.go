package builtin

import (
	"github.com/infrago/base"
	"github.com/infrago/infra"
)

func PagedVar(offset, limit int, exts ...base.Vars) base.Var {
	config := base.Vars{
		"offset": infra.Define("int", true, offset, "offset"),
		"limit":  infra.Define("int", true, limit, "limit"),
		"count":  infra.Define("int", true, 0, "返回记录总数"),
	}
	config = infra.VarsExtend(config, exts...)
	return base.Var{
		Type: "json", Required: true, Name: "分页", Children: config,
	}
}

func PagedVars(offset, limit int, exts ...base.Vars) base.Vars {
	config := base.Vars{
		"offset": infra.Define("int", true, offset, "offset"),
		"limit":  infra.Define("int", true, limit, "limit"),
	}
	return infra.VarsExtend(config, exts...)
}

func DigitDecodeConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: true, Name: name, Text: text,
		Decode: "digit",
	}
}
func DigitsDecodeConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "[int]", Required: true, Name: name, Text: text,
		Decode: "digits",
	}
}
func DigitDecodeNullableConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: false, Name: name, Text: text,
		Decode: "digit",
	}
}
func DigitEncodeConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: true, Name: name, Text: text,
		Encode: "digit",
	}
}
func DigitsEncodeConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "[int]", Required: true, Name: name, Text: text,
		Encode: "digits",
	}
}
func DigitEncodeNullableConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: false, Name: name, Text: text,
		Encode: "digit",
	}
}
func DigitCryptoConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: true, Name: name, Text: text,
		Encode: "digit", Decode: "digit",
	}
}
func DigitsCryptoConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "[int]", Required: true, Name: name, Text: text,
		Encode: "digits", Decode: "digits",
	}
}
func DigitCryptoNullableConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "int", Required: false, Name: name, Text: text,
		Encode: "digit", Decode: "digit",
	}
}

func DigitsCryptoNullableConfig(name string, texts ...string) base.Var {
	text := name
	if len(texts) > 0 {
		text = texts[0]
	}
	return base.Var{
		Type: "[int]", Required: false, Name: name, Text: text,
		Encode: "digits", Decode: "digits",
	}
}
