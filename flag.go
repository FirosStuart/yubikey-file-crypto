package main

import (
	"flag"
)

const (
	defaultFlagString = ""
	flagBoolIsNil     = false
)

type FlagSetter struct {
	FileName *string // default defaultFlagString
	Setup    *bool   // default flagBoolIsNil
	Reset    *bool   // default flagBoolIsNil
	Encrypt  *bool   // default flagBoolIsNil
	Decrypt  *bool   // default flagBoolIsNil
}

func (p *FlagSetter) FlagSetting() {
	p.Setup = flag.Bool("setup", flagBoolIsNil, "configure a new YubiKey")
	p.Reset = flag.Bool("reset", flagBoolIsNil, "reset a YubiKey")
	p.Encrypt = flag.Bool("encrypt", flagBoolIsNil, "encrypt some file")
	p.FileName = flag.String("filename", defaultFlagString, "file to encrypt or decrypt")
	p.Decrypt = flag.Bool("decrypt", flagBoolIsNil, "decrypt some file")
	flag.Parse()
}
