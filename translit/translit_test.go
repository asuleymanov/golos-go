package translit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	enc, _ := encode("teststr")
	require.Equal(t, "teststr", enc)

	enc, _ = encode("test str")
	require.Equal(t, "test-str", enc)

	enc, _ = encode("тест")
	require.Equal(t, "test", enc)

	enc, _ = encode("тест с пробелом жшщ")
	require.Equal(t, "test-s-probelom-zhshshch", enc)
}

func TestDecode(t *testing.T) {
	dec, _ := decode("test-s-probelom-zhshshch")
	require.Equal(t, "тест с пробелом жшщ", dec)
}

func TestEncodeDecode(t *testing.T) {
	str := "абвг деёжз ийкл мноп рсту фхцч шщ ь ъ эюя"
	enc, _ := encode(str)
	dec, _ := decode(enc)

	require.Equal(t, str, dec)
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
		require.Equal(t, decTags[i], tag)
	}
}
