package redis

import (
	"bytes"
)


/************************* Redis Encoder & Decoder *************************/

type RedisEncoder struct {

}

func (encoder *RedisEncoder) Encode(v interface{}) ([]byte, error) {
	panic("Not Implemented!")

	return nil, nil
}

type RedisDecoder struct {

}

func (decoder *RedisDecoder) TryReadLen(bs []byte) (int, error) {
	
}

func (decoder *RedisDecoder) Decode(bs []byte, v interface{}) error {

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
