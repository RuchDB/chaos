package redis

import (
	"bytes"
)


/************************* Redis Encoder & Decoder *************************/

type RedisEncoder struct { }

func (encoder *RedisEncoder) Encode(v interface{}) ([]byte, error) {
	switch data := v.(type) {
	case string:
		return append([]byte(data), '\r', '\n'), nil
	case *string:
		return append([]byte(*data), '\r', '\n'), nil
	case []byte:
		return append(data, '\r', '\n'), nil
	case *[]byte:
		return append(*data, '\r', '\n'), nil
	default:
		return nil, ERR_CODEC_TYPE_UNSUPPORTED
	}
}

type RedisDecoder struct { }

func (decoder *RedisDecoder) TryReadLen(bs []byte) (int, error) {
	pos := bytes.IndexByte(bs, '\n')
	if pos < 0 {
		return 0, ERR_CODEC_DATA_INCOMPLETE
	}

	if pos == 0 || bs[pos - 1] != '\r' {
		return 0, ERR_CODEC_DATA_INVALID
	}

	return pos + 1, nil
}

func (decoder *RedisDecoder) Decode(bs []byte, v interface{}) error {
	if len(bs) < 2 || !(bs[len(bs) - 2] == '\r' && bs[len(bs) - 1] == '\n') {
		return ERR_CODEC_DATA_INVALID
	}

	switch data := v.(type) {
	case *[]byte:
		*data = bs[:len(bs) - 2]
	case *string:
		*data = string(bs[:len(bs) - 2])
	default:
		return ERR_CODEC_TYPE_UNSUPPORTED
	}

	return nil
}


/************************* Redis Codec *************************/

type RedisCodec struct {
	RedisEncoder
	RedisDecoder
}

func (codec *RedisCodec) Id() int {
	return CODEC_ID_REDIS
}

func (codec *RedisCodec) Name() string {
	return CODEC_NAME_REDIS
}

func (codec *RedisCodec) Encode(v interface{}) ([]byte, error) {
	return codec.RedisEncoder.Encode(v)
}

func (codec *RedisCodec) TryReadLen(bs []byte) (int, error) {
	return codec.RedisDecoder.TryReadLen(bs)
}

func (codec *RedisCodec) Decode(bs []byte, v interface{}) error {
	return codec.RedisDecoder.Decode(bs, v)
}
