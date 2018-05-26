package translit

import (
	"testing"
)

func TestEncode(t *testing.T) {
	enc, _ := encode("teststr")
	if enc != "teststr" {
		t.Errorf("Wrong encode string for latin. Expected '%s', got '%s'", "teststr", enc)
	}

	enc, _ = encode("test str")
	if enc != "test-str" {
		t.Errorf("Wrong encode string for latin with space. Expected '%s', got '%s'", "test-str", enc)
	}

	enc, _ = encode("тест")
	if enc != "test" {
		t.Errorf("Wrong encode string for ru. Expected '%s', got '%s'", "test", enc)
	}

	enc, _ = encode("тест с пробелом жшщ")
	if enc != "test-s-probelom-zhshshch" {
		t.Errorf("Wrong encode string for ru with space. Expected '%s', got '%s'", "test-s-probelom-zhshshch", enc)
	}
}

func TestDecode(t *testing.T) {
	dec, _ := decode("test-s-probelom-zhshshch")

	if dec != "тест с пробелом жшщ" {
		t.Errorf("Wrong decode string for ru with space and big simbols. Expected '%s', got '%s'", "тест с пробелом жшщ", dec)
	}
}

func TestEncodeDecode(t *testing.T) {
	str := "абвг деёжз ийкл мноп рсту фхцч шщ ь ъ эюя"
	enc, _ := encode(str)
	dec, _ := decode(enc)

	if dec != str {
		t.Errorf("Wrong encode/decode string for ru with space and big simbols. Expected '%s', got '%s'", str, dec)
	}
}

func TestEncodeDecodeTags(t *testing.T) {
	tags := []string{
		"simple-tag1",
		"Русский тэг1",
		"РУССКИЙ ТЭГ 2",
		"твердые формы",
		"simple-tag-3",
	}
	results := []string{
		"simple-tag1",
		"русский тэг1",
		"русский тэг 2",
		"твердые формы",
		"simple-tag-3",
	}

	encTags := EncodeTags(tags)
	decTags := DecodeTags(encTags)

	for i, tag := range results {
		if tag != decTags[i] {
			t.Errorf("Wrong encode/decode tags. Expected '%s', got '%s'", tag, decTags[i])
		}
	}
}
