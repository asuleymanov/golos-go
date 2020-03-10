package transaction

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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
	return encoder.WriteBytes(b[:n])
}

//EncodeUVarint converting uint64 to byte
func (encoder *Encoder) EncodeUVarint(i uint64) error {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, i)
	return encoder.WriteBytes(b[:n])
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

//EncodeString converting string to byte
func (encoder *Encoder) EncodeString(v string) error {
	if err := encoder.EncodeUVarint(uint64(len(v))); err != nil {
		return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", v, err)
	}

	return encoder.WriteString(v)
}

//EncodeBool converting bool to byte
func (encoder *Encoder) EncodeBool(b bool) error {
	if b {
		return encoder.EncodeNumber(byte(1))
	}
	return encoder.EncodeNumber(byte(0))
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
		return encoder.WriteBytes(v)
	case bool:
		return encoder.EncodeBool(v)
	default:
		return fmt.Errorf("encoder: unsupported type encountered")
	}
}

func (encoder *Encoder) WriteBytes(bs []byte) error {
	if _, err := encoder.w.Write(bs); err != nil {
		return fmt.Errorf("encoder: failed to write bytes: %v\n Error : %w", bs, err)
	}
	return nil
}

func (encoder *Encoder) WriteString(s string) error {
	if _, err := io.Copy(encoder.w, strings.NewReader(s)); err != nil {
		return fmt.Errorf("encoder: failed to write string: %v\n Error : %w", s, err)
	}
	return nil
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
			return encoder.WriteBytes(pkn2)
		}
		pkn3, _ := btcec.ParsePubKey(pkn2, btcec.S256())
		if _, err := encoder.w.Write(pkn3.SerializeCompressed()); err != nil {
			return fmt.Errorf("encoder: failed to write bytes: %v\n Error : %w", pkn3.SerializeCompressed(), err)
		}
		return nil
	}
	return errors.New("Public key is incorrect")
}
