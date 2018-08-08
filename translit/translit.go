package translit

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

//EncodeTags transliteration of an array of tags
func EncodeTags(tag []string) []string {
	var arrEncTag []string
	for _, val := range tag {
		str, count := encode(val)
		if count > 0 {
			str = "ru--" + str
		}
		arrEncTag = append(arrEncTag, str)
	}
	return arrEncTag
}

//DecodeTags transliteration of an array of tags
func DecodeTags(tags []string) []string {
	arrDecTag := make([]string, 0)
	for _, tag := range tags {
		dec := DecodeTag(tag)
		arrDecTag = append(arrDecTag, dec)
	}
	return arrDecTag
}

//EncodeTag transliteration of a tag
func EncodeTag(tag string) string {
	str, count := encode(tag)
	if count > 0 {
		str = "ru--" + str
	}
	return str
}

//DecodeTag transliteration of a tag
func DecodeTag(tag string) string {
	if len(tag) < 5 {
		return tag
	}
	if tag[0:4] == "ru--" {
		str, _ := decode(tag[4:])
		return str
	}
	return tag
}

//EncodeTitle transliteration of the title
func EncodeTitle(title string) string {
	if title == "" {
		return title
	}

	var str string
	reg, err := regexp.Compile("[^a-zA-Z0-9а-яА-Я.,]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(title, "-")
	s1, _ := encode(processedString)
	s2 := strings.Replace(s1, ".", "", -1)
	s3 := strings.Split(s2, "")
	if s3[0] == "-" {
		str = strings.Join(s3[1:], "")
	} else {
		str = strings.Join(s3, "")
	}
	return str
}

//DecodeTitle transliteration of the title
func DecodeTitle(title string) string {
	str, _ := decode(title)
	return str
}

func encode(text string) (string, int) {
	if text == "" {
		return "", 0
	}
	text = strings.ToLower(text)
	text = strings.Replace(text, "ые", "yie", -1)

	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	// Previous, next letter for special processor
	var rr string
	var ok bool

	i := 0

	for {
		r, _, err := input.ReadRune()
		if err != nil {
			break
		}

		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
			i++
			continue
		} else {
			output.WriteString(string(r))
			continue
		}
		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
			i++
		} else {
			output.WriteString(string(r))
			continue
		}
	}

	return output.String(), i
}

func decode(str string) (string, int) {
	if str == "" {
		return "", 0
	}

	str = strings.Replace(str, "yie", "ые", -1)

	for _, trans := range decOrder {
		str = strings.Replace(str, trans, decMap[trans], -1)
	}

	return str, len(str)
}
