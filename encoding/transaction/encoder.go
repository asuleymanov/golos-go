package transaction

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

//Encoder structure for the converter
type Encoder struct {
	w io.Writer
}

//NewEncoder initializing a new converter
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

//EncodeVarint converting int64 to byte
func (encoder *Encoder) EncodeVarint(i int64) error {
	if i >= 0 {
		return encoder.EncodeUVarint(uint64(i))
	}

	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(b, i)
	return encoder.writeBytes(b[:n])
}

//EncodeUVarint converting uint64 to byte
func (encoder *Encoder) EncodeUVarint(i uint64) error {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, i)
	return encoder.writeBytes(b[:n])
}

//EncodeNumber converting number to byte
func (encoder *Encoder) EncodeNumber(v interface{}) error {
	if err := binary.Write(encoder.w, binary.LittleEndian, v); err != nil {
		return fmt.Errorf("encoder: failed to write number: %v\n Error : %w", v, err)
	}
	return nil
}

//EncodeArrString converting []string to byte
func (encoder *Encoder) EncodeArrString(v []string) error {
	if err := encoder.EncodeUVarint(uint64(len(v))); err != nil {
		return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", v, err)
	}
	for _, val := range v {
		if err := encoder.EncodeUVarint(uint64(len(val))); err != nil {
			return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", val, err)
		}
		if _, err := io.Copy(encoder.w, strings.NewReader(val)); err != nil {
			return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", val, err)
		}
	}
	return nil
}

//Encode function that determines the input values of which converter to use
func (encoder *Encoder) Encode(v interface{}) error {
	if marshaller, ok := v.(TransactionMarshaller); ok {
		return marshaller.MarshalTransaction(encoder)
	}

	switch v := v.(type) {
	case int:
		return encoder.EncodeNumber(v)
	case int8:
		return encoder.EncodeNumber(v)
	case int16:
		return encoder.EncodeNumber(v)
	case int32:
		return encoder.EncodeNumber(v)
	case int64:
		return encoder.EncodeNumber(v)

	case uint:
		return encoder.EncodeNumber(v)
	case uint8:
		return encoder.EncodeNumber(v)
	case uint16:
		return encoder.EncodeNumber(v)
	case uint32:
		return encoder.EncodeNumber(v)
	case uint64:
		return encoder.EncodeNumber(v)

	case string:
		return encoder.EncodeString(v)
	case []byte:
		return encoder.writeBytes(v)
	case bool:
		return encoder.EncodeBool(v)
	default:
		return fmt.Errorf("encoder: unsupported type encountered")
	}
}

//EncodeString converting string to byte
func (encoder *Encoder) EncodeString(v string) error {
	if err := encoder.EncodeUVarint(uint64(len(v))); err != nil {
		return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", v, err)
	}

	return encoder.writeString(v)
}

func (encoder *Encoder) writeBytes(bs []byte) error {
	if _, err := encoder.w.Write(bs); err != nil {
		return fmt.Errorf("encoder: failed to write bytes: %v\n Error : %w", bs, err)
	}
	return nil
}

func (encoder *Encoder) writeString(s string) error {
	if _, err := io.Copy(encoder.w, strings.NewReader(s)); err != nil {
		return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", s, err)
	}
	return nil
}

//EncodeBool converting bool to byte
func (encoder *Encoder) EncodeBool(b bool) error {
	if b {
		return encoder.EncodeNumber(byte(1))
	}
	return encoder.EncodeNumber(byte(0))
}

//EncodeMoney converting Asset to byte
func (encoder *Encoder) EncodeMoney(s string) error {
	r, _ := regexp.Compile(`^[0-9]+\.?[0-9]* [A-Za-z0-9]+$`)
	if r.MatchString(s) {
		asset := strings.Split(s, " ")
		amm, errParsInt := strconv.ParseInt(strings.Replace(asset[0], ".", "", -1), 10, 64)
		if errParsInt != nil {
			return errParsInt
		}
		ind := strings.Index(asset[0], ".")
		var perc int
		if ind == -1 {
			perc = 0
		} else {
			perc = len(asset[0]) - ind - 1
		}
		if err := binary.Write(encoder.w, binary.LittleEndian, amm); err != nil {
			return fmt.Errorf("encoder: failed to write number: %v\n Error : %w", amm, err)
		}
		if err := binary.Write(encoder.w, binary.LittleEndian, byte(perc)); err != nil {
			return fmt.Errorf("encoder: failed to write number: %v\n Error : %w", perc, err)
		}

		if _, err := io.Copy(encoder.w, strings.NewReader(asset[1])); err != nil {
			return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", asset[1], err)
		}

		for i := byte(len(asset[1])); i < 7; i++ {
			if err := binary.Write(encoder.w, binary.LittleEndian, byte(0)); err != nil {
				return fmt.Errorf("encoder: failed to write number: %v\n Error : %w", 0, err)
			}
		}
		return nil
	}
	return errors.New("Expecting amount like '99.000 SYMBOL'")
}

//EncodePubKey converting PubKey to byte
func (encoder *Encoder) EncodePubKey(s string) error {
	pkn1 := strings.Join(strings.Split(s, "")[3:], "")
	b58 := base58.Decode(pkn1)
	chs := b58[len(b58)-4:]
	pkn2 := b58[:len(b58)-4]
	chHash := ripemd160.New()
	_, errHash := chHash.Write(pkn2)
	if errHash != nil {
		return errHash
	}
	nchs := chHash.Sum(nil)[:4]
	if bytes.Equal(chs, nchs) {
		if string(pkn2) == string(make([]byte, 33)) {
			return encoder.writeBytes(pkn2)
		}
		pkn3, _ := btcec.ParsePubKey(pkn2, btcec.S256())
		if _, err := encoder.w.Write(pkn3.SerializeCompressed()); err != nil {
			return fmt.Errorf("encoder: failed to write bytes: %v\n Error : %w", pkn3.SerializeCompressed(), err)
		}
		return nil
	}
	return errors.New("Public key is incorrect")
}
