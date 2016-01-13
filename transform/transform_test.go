package transform

import (
	"io/ioutil"
	"strings"
	"testing"
)

var (
	sjisStr = "\x82\xb1\x82\xea\x82\xcd\x20\x53\x68\x69\x66\x74\x4a\x49\x53\x20\x82\xcc\x95\xb6\x8e\x9a\x97\xf1\x82\xc5\x82\xb7\x81\x42"
	utf8Str = "これは ShiftJIS の文字列です。"
)

func TestRead(t *testing.T) {
	str := strings.NewReader(sjisStr)

	reader := NewReader(str)

	buff, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}

	actual := string(buff)

	if actual != utf8Str {
		t.Errorf("got %v\nwant %v", actual, utf8Str)
	}
}
