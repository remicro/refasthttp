package refasthttp

import (
	"github.com/remicro/api/cloud/balancer"
	"github.com/remicro/api/logging"
	"github.com/remicro/api/net/rehttp"
	"github.com/remicro/api/serialization"
	"github.com/valyala/fasthttp"
)

func New() rehttp.Builder {
	return &fastHttpClient{
		req:    fasthttp.AcquireRequest(),
		uri:    fasthttp.AcquireURI(),
		logger: dummyLogger{},
	}
}

type fastHttpClient struct {
	url        string
	method     int
	decodeType rehttp.ContentType
	logger     logging.Logger
	req        *fasthttp.Request
	before     rehttp.Before
	decoder    serialization.Decoder
	encoder    serialization.Encoder
	encObj     interface{}
	decObj     interface{}
	uri        *fasthttp.URI
	bln balancer.Balancer
}

func (fhc *fastHttpClient) Balancer(bln balancer.Balancer) rehttp.Builder {
	fhc.bln = bln
	return fhc
}

func (fhc *fastHttpClient) Service(name string) rehttp.Builder {
	node, err := fhc.bln.Find(name)
	if err != nil {
		return fhc
	}
	fhc.uri.Parse(nil, []byte(node.Address()))
	return fhc
}

func (fhc *fastHttpClient) Logger(logger logging.Logger) rehttp.Builder {
	fhc.logger = logger
	return fhc
}

func (fhc *fastHttpClient) DecodeType(contentType rehttp.ContentType) rehttp.Builder {
	fhc.decodeType = contentType
	return fhc
}

func (fhc *fastHttpClient) Address(address string) rehttp.Builder {
	fhc.uri.Parse(nil, []byte(address))
	return fhc
}

func (fhc *fastHttpClient) PUT(url string) rehttp.Builder {
	fhc.uri.SetPath(url)
	fhc.req.Header.SetMethod("PUT")
	return fhc
}

func (fhc *fastHttpClient) ToEncode(object interface{}) rehttp.Builder {
	fhc.encObj = object
	return fhc
}

func (fhc *fastHttpClient) ToDecode(object interface{}) rehttp.Builder {
	fhc.decObj = object
	return fhc
}

func (fhc *fastHttpClient) Encoder(encoder serialization.Encoder) rehttp.Builder {
	fhc.encoder = encoder
	return fhc
}

func (fhc *fastHttpClient) Decoder(decoder serialization.Decoder) rehttp.Builder {
	fhc.decoder = decoder
	return fhc
}

func (fhc *fastHttpClient) GET(u string) rehttp.Builder {
	fhc.uri.SetPath(u)
	return fhc
}

func (fhc *fastHttpClient) POST(u string) rehttp.Builder {
	fhc.uri.SetPath(u)
	fhc.req.Header.SetMethod("POST")
	return fhc
}

func (fhc *fastHttpClient) DELETE(u string) rehttp.Builder {
	fhc.uri.SetPath(u)
	fhc.req.Header.SetMethod("DELETE")
	return fhc
}

func (fhc *fastHttpClient) PATCH(u string) rehttp.Builder {
	fhc.uri.SetPath(u)
	fhc.req.Header.SetMethod("PATCH")
	return fhc
}

func (fhc *fastHttpClient) OPTIONS(u string) rehttp.Builder {
	fhc.uri.SetPath(u)
	fhc.req.Header.SetMethod("OPTIONS")
	return fhc
}

func (fhc *fastHttpClient) QueryParam(key, value string) rehttp.Builder {
	fhc.uri.QueryArgs().Set(key, value)
	return fhc
}

func (fhc *fastHttpClient) ContentType(contentType rehttp.ContentType) rehttp.Builder {
	fhc.req.Header.SetContentType(string(contentType))
	return fhc
}

func (fhc *fastHttpClient) Header(key, value string) rehttp.Builder {
	fhc.req.Header.Set(key, value)
	return fhc
}

func (fhc *fastHttpClient) Cookie(key string, value []byte) rehttp.Builder {
	fhc.req.Header.SetCookieBytesKV([]byte(key), value)
	return fhc
}
func (fhc *fastHttpClient) Before(before rehttp.Before) rehttp.Builder {
	fhc.before = before
	return fhc
}

func (fhc *fastHttpClient) Go() (response rehttp.Response, err error) {
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if fhc.encObj != nil && fhc.encoder != nil {
		var data []byte
		data, err = fhc.encoder.Encode(fhc.encObj)
		if err != nil {
			fhc.logger.Debug().Log("can't encode object for request")
			return
		}
		fhc.req.SetBody(data)
	}

	if fhc.decodeType != "" {
		fhc.req.Header.Add("Accept", fhc.decodeType.String())
	}

	fhc.req.SetRequestURIBytes(fhc.uri.FullURI())
	if fhc.before != nil {
		fhc.before(fhc, string(fhc.uri.FullURI()), fhc.req.Body())
	}
	err = client.Do(fhc.req, resp)
	if err != nil {
		return
	}
	response = &responseImpl{
		response: resp,
	}

	if fhc.decObj != nil && fhc.decoder != nil && string(resp.Header.ContentType()) == fhc.decodeType.String() {
		err = fhc.decoder.Decode(fhc.decObj, resp.Body())
		if err != nil {
			fhc.logger.Debug().
				Int("status", response.Status()).
				String("content-type", string(resp.Header.ContentType())).
				Log("can't decode response")
		}
	}
	return
}
