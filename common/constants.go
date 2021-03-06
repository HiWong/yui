package common

const (
	Help = `Yui version ` + Version + `, https://github.com/RitterHou/yui

Usage:

    yui command [arguments]

The commands are:

    build   build the source file
    run     run compiled file or source file
    shell   open interactive shell
    dec     decompile binary code to instructions

`
	FileExtension = ".yuicode"
	magicHex      = 0x10151657 // 10月15日 16时57分

	Begin = "BEGIN"
	End   = "END"
)

var (
	Magic []byte // Magic Number

	PUSH     = byte(0x00)
	PLUS     = byte(0x01)
	MINUS    = byte(0x02)
	MULTIPLY = byte(0x03)
	DIVIDE   = byte(0x04)

	INT   = byte(0x00)
	FLOAT = byte(0x01)
)

func init() {
	Magic = Int2ByteArray(int32(magicHex))
}
