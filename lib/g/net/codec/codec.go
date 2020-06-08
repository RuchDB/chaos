package codec

import (
	"errors"
)

const (
	CODEC_ID_REDIS = 1

	CODEC_NAME_REDIS = "Redis Codec"
)

/************************* Codec Types *************************/

var (
	ERR_CODEC_TYPE_UNSUPPORTED = errors.New("unsupported data type")
	ERR_CODEC_DATA_INCOMPLETE  = errors.New("incomplete data")
	ERR_CODEC_DATA_INVALID     = errors.New("invalid data")
)

type Codec interface {
	Id() int
	Name() string

	Encoder
	Decoder
}

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

type Decoder interface {
	TryReadLen(bs []byte) (int, error)
	Decode(bs []byte, v interface{}) error
}

/************************* Codec Map *************************/

var (
	ERR_CODEC_EXIST   = errors.New("codec already exists")
	ERR_CODEC_UNEXIST = errors.New("codec unexists")
)

var codecs []Codec

func init() {
	codecs = make([]Codec, 0, 8)

	RegisterCodec(&RedisCodec{})
}

func RegisterCodec(codec Codec) error {
	for _, c := range codecs {
		if codec.Id() == c.Id() || codec.Name() == c.Name() {
			return ERR_CODEC_EXIST
		}
	}

	codecs = append(codecs, codec)

	return nil
}

func GetCodecById(codecId int) (Codec, error) {
	for _, c := range codecs {
		if c.Id() == codecId {
			return c, nil
		}
	}

	return nil, ERR_CODEC_UNEXIST
}

func GetCodecByName(codecName string) (Codec, error) {
	for _, c := range codecs {
		if c.Name() == codecName {
			return c, nil
		}
	}

	return nil, ERR_CODEC_UNEXIST
}
