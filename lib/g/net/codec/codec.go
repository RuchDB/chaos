package proto

/************************* Codec Types *************************/

const (
	ERR_ENCODE_TYPE_UNSUPPORTED = errors.New("unsupported data type")

	ERR_DECODE_DATA_INCOMPLETE = errors.New("incomplete data")
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

const (
	ERR_CODEC_EXIST   = errors.New("codec already exists")
	ERR_CODEC_UNEXIST = errors.New("codec unexists")
)

var codecMap map[int]Codec

func init() {
	codecMap = make(map[int]Codec)
}

func RegisterCodec(codec Codec) error {
	if _, exist := codecMap[codec.Id()]; exist {
		return ERR_CODEC_EXIST
	}

	codecMap[codec.Id()] = codec
	
	return nil
}

func GetCodec(codecId int) (Codec, error) {
	codec, exist := codecMap[codecId]
	if !exist {
		return nil, ERR_CODEC_UNEXIST
	}

	return codec, nil
}
