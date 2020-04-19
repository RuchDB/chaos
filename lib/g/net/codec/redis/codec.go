package redis

import (
	"bytes"

	"github.com/RuchDB/chaos/proto"
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

