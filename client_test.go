package refasthttp

import (
	"github.com/remicro/api/net/rehttp"
	"github.com/remicro/refasthttp/fixture"
	"github.com/remicro/trifle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestNew(t *testing.T) {
	fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
		assert.Equal(t, "GET", string(ctx.Method()))
		ctx.Write([]byte("OK"))
	})
	defer fx.Finish()
	res, err := New().
		GET(fx.Address()).
		Go()
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusOK, res.Status())
	assert.Equal(t, "OK", string(res.Body()))
}

type Object struct {
	Label string `json:"label"`
}

func TestFastHttpClient_GET(t *testing.T) {

	t.Run("decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()
		var result Object
		res, err := New().GET(fx.Address()).Decoder(reFastHttpFixture.Decoder()).ToDecode(&result).Go()
		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_POST(t *testing.T) {
	t.Run("encode request, decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		req := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "POST", string(ctx.Method()))
			var rr Object

			require.NoError(t, reFastHttpFixture.Decoder().Decode(&rr, ctx.PostBody()))
			assert.Equal(t, req, rr)

			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object
		res, err := New().
			POST(fx.Address()).
			Encoder(reFastHttpFixture.Encoder()).
			Decoder(reFastHttpFixture.Decoder()).
			ToEncode(&req).
			ToDecode(&result).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_PUT(t *testing.T) {
	t.Run("encode request, decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		req := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "PUT", string(ctx.Method()))
			var rr Object

			require.NoError(t, reFastHttpFixture.Decoder().Decode(&rr, ctx.PostBody()))
			assert.Equal(t, req, rr)

			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object
		res, err := New().
			PUT(fx.Address()).
			Encoder(reFastHttpFixture.Encoder()).
			Decoder(reFastHttpFixture.Decoder()).
			ToEncode(&req).
			ToDecode(&result).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_DELETE(t *testing.T) {
	t.Run("decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "DELETE", string(ctx.Method()))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()
		var result Object
		res, err := New().DELETE(fx.Address()).Decoder(reFastHttpFixture.Decoder()).ToDecode(&result).Go()
		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_OPTIONS(t *testing.T) {
	t.Run("decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "OPTIONS", string(ctx.Method()))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()
		var result Object
		res, err := New().OPTIONS(fx.Address()).Decoder(reFastHttpFixture.Decoder()).ToDecode(&result).Go()
		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_PATCH(t *testing.T) {
	t.Run("decode response", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "PATCH", string(ctx.Method()))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()
		var result Object
		res, err := New().PATCH(fx.Address()).Decoder(reFastHttpFixture.Decoder()).ToDecode(&result).Go()
		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_ContentType(t *testing.T) {

	t.Run("expect transfer content type of request to server side", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, rehttp.ContentTypeJson.String(), string(ctx.Request.Header.ContentType()))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()
		var result Object
		res, err := New().
			GET(fx.Address()).
			ContentType(rehttp.ContentTypeJson).
			Decoder(reFastHttpFixture.Decoder()).
			ToDecode(&result).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_Header(t *testing.T) {

	t.Run("expect transfer header to server side", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		headerKey, headerValue := trifle.String(), trifle.String()
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, rehttp.ContentTypeJson.String(), string(ctx.Request.Header.ContentType()))
			assert.Equal(t, headerValue, string(ctx.Request.Header.Peek(headerKey)))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object
		res, err := New().
			GET(fx.Address()).
			ContentType(rehttp.ContentTypeJson).Header(headerKey, headerValue).
			Decoder(reFastHttpFixture.Decoder()).
			ToDecode(&result).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_Cookie(t *testing.T) {
	t.Run("expect transfer header to server side", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		cookieKey, cookieValue := trifle.String(), trifle.String()
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, rehttp.ContentTypeJson.String(), string(ctx.Request.Header.ContentType()))
			assert.Equal(t, cookieValue, string(ctx.Request.Header.Cookie(cookieKey)))
			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object
		res, err := New().
			GET(fx.Address()).
			ContentType(rehttp.ContentTypeJson).
			Cookie(cookieKey, []byte(cookieValue)).
			Decoder(reFastHttpFixture.Decoder()).
			ToDecode(&result).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_QueryParam(t *testing.T) {
	t.Run("expect correct send query params", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		queryKey, queryValue := trifle.String(), trifle.String()
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, queryValue, string(ctx.QueryArgs().Peek(queryKey)))

			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object

		res, err := New().
			GET(fx.Address()).QueryParam(queryKey, queryValue).
			ToDecode(&result).
			Decoder(reFastHttpFixture.Decoder()).
			Go()

		require.NoError(t, err)
		assert.Equal(t, rehttp.ContentTypeJson, res.ContentType())
		assert.Equal(t, 200, res.Status())
		assert.Equal(t, exp, result)
	})
}

func TestFastHttpClient_Go(t *testing.T) {
	t.Run("expect error on encoding", func(t *testing.T) {
		obj := Object{
			Label: trifle.String(),
		}
		exp := trifle.UnexpectedError()

		res, err := New().
			GET(trifle.String()).
			ToEncode(&obj).
			Encoder(reFastHttpFixture.FailEncoder(exp)).
			Go()
		require.Nil(t, res)
		assert.Equal(t, exp, err)
	})

	t.Run("expect error on object decoding", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		expErr := trifle.UnexpectedError()

		queryKey, queryValue := trifle.String(), trifle.String()
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, queryValue, string(ctx.QueryArgs().Peek(queryKey)))

			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object

		res, err := New().
			GET(fx.Address()).QueryParam(queryKey, queryValue).
			ToDecode(&result).
			Decoder(reFastHttpFixture.FailDecoder(expErr)).
			Go()

		require.Error(t, expErr, err)
		require.NotNil(t, res, err)
		require.NoError(t, reFastHttpFixture.Decoder().Decode(&result, res.Body()))
		assert.Equal(t, exp, result)
	})
	t.Run("expect error on do request", func(t *testing.T) {
		res, err := New().
			GET(trifle.String()).
			Go()
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestFastHttpClient_Before(t *testing.T) {
	t.Run("expect error on object decoding", func(t *testing.T) {
		exp := Object{
			Label: trifle.String(),
		}
		expErr := trifle.UnexpectedError()

		queryKey, queryValue := trifle.String(), trifle.String()
		fx := reFastHttpFixture.New(t, func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "GET", string(ctx.Method()))
			assert.Equal(t, queryValue, string(ctx.QueryArgs().Peek(queryKey)))

			data, err := reFastHttpFixture.Encoder().Encode(&exp)
			require.NoError(t, err)
			ctx.Response.Header.SetContentType("application/json")
			ctx.Write(data)
		})
		defer fx.Finish()

		var result Object

		res, err := New().
			GET(fx.Address()).QueryParam(queryKey, queryValue).
			ToDecode(&result).
			Decoder(reFastHttpFixture.FailDecoder(expErr)).
			Before(func(b rehttp.Builder, url string, body []byte) {
				assert.Equal(t, fx.Address()+"/?"+queryKey+"="+queryValue, url)
			}).
			Go()

		require.Error(t, expErr, err)
		require.NotNil(t, res, err)
		require.NoError(t, reFastHttpFixture.Decoder().Decode(&result, res.Body()))
		assert.Equal(t, exp, result)
	})
}
