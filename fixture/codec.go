package reFastHttpFixture

import (
	"encoding/json"
	"github.com/remicro/api/serialization"
)

type coder struct {
	expErr error
}

func (c coder) Encode(object interface{}) (data []byte, err error) {
	if c.expErr != nil {
		err = c.expErr
		return
	}
	return json.Marshal(object)
}

func (c coder) Decode(object interface{}, data []byte) (err error) {
	if c.expErr != nil {
		err = c.expErr
		return
	}
	return json.Unmarshal(data, object)
}

func Encoder() serialization.Encoder {
	return coder{}
}

func Decoder() serialization.Decoder {
	return coder{}
}

func FailEncoder(err error) serialization.Encoder {
	return coder{expErr: err}
}

func FailDecoder(err error) serialization.Decoder {
	return coder{expErr: err}
}
