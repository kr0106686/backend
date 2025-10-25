package test__test

import (
	"net/http/httptest"
	"testing"

	"github.com/kr0106686/backend/handler"
	"github.com/stretchr/testify/require"
)

func Test_Http(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/hello", nil)
	handler.Hello(rr, req)

	require.Equal(t, "hello world", rr.Body.String())
}
