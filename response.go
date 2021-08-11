package refasthttp

import (
	"github.com/remicro/api/net/rehttp"
	"github.com/valyala/fasthttp"
)

type responseImpl struct {
	response      *fasthttp.Response
	acquiredError error
	decodedObject interface{}
}

func (res *responseImpl) Status() (code int) {
	return res.response.StatusCode()
}

func (res *responseImpl) Decoded() (decoded interface{}) {
	return res.decodedObject
}

func (res *responseImpl) Body() []byte {
	return res.response.Body()
}

func (res *responseImpl) ContentType() (contentType rehttp.ContentType) {
	return rehttp.ContentType(res.response.Header.ContentType())
}

func (res *responseImpl) Header(key string) (values []string) {
	value := res.response.Header.Peek(key)
	if len(value) > 0 {
		values = append(values, string(value))
	}
	return
}

func (res *responseImpl) Error() (err error) {
	return res.acquiredError
}
