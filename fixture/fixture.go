package reFastHttpFixture

import (
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"net"
	"testing"
)

type fixture struct {
	l net.Listener
	*fasthttp.Server
	t *testing.T
}

func (fx *fixture) Finish() {
	require.NoError(fx.t, fx.l.Close())
}

func (fx *fixture) Address() string {
	return "http://" + fx.l.Addr().String()
}

func New(t *testing.T, handler fasthttp.RequestHandler) *fixture {
	fx := &fixture{
		Server: &fasthttp.Server{},
		t:      t,
	}

	fx.Handler = handler
	l, err := net.Listen("tcp4", "localhost:0")
	require.NoError(t, err)
	fx.l = l
	go fx.Server.Serve(l)

	return fx
}
