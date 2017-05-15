package Ascii

import "testing"

func ExtendedGreetingASCII() string {
	a:= []byte("\x53\x61\x6c\x75\x74\x2c\x20\xc3\xa7\x61\x20\x76\x61" +
		"\x20\xc2\xb0\x2d\x29\x20\xe2\x80\x8b\xe2\x82\xac\x35\x30")
	return string(a)
}

func TestExtendedGreetingASCII(t *testing.T) {
    if !(isExtendedASCII(ExtendedGreetingASCII())){
        t.Fail()
    }
}

func isExtendedASCII(s string) bool {
    for _, c := range s {
        if c < 128 {
            return false
        }
    }
    return true
}
