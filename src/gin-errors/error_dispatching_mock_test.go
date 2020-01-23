package errors

import (
	"bufio"
	"github.com/stretchr/testify/mock"
	"net"
	"net/http"
)

type MockError struct {
	mock.Mock
}

func (m MockError) Error() string {
	panic("implement me")
}

func (m MockError) Code() string {

	args := m.Called()
	return args.String(0)
}

func (m MockError) Message() string {
	panic("implement me")
}

func (m MockError) OrigErr() error {
	panic("implement me")
}

// net....

type MockWriter struct {
	*mock.Mock
}

func NewMockWriter() *MockWriter {
	return &MockWriter{Mock: &mock.Mock{}}
}

func (m MockWriter) Header() http.Header {
	panic("implement me")
}

func (m MockWriter) Write([]byte) (int, error) {
	panic("implement me")
}

func (m MockWriter) WriteHeader(statusCode int) {

	m.Called(statusCode)
}

func (m MockWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	panic("implement me")
}

func (m MockWriter) Flush() {
	panic("implement me")
}

func (m MockWriter) CloseNotify() <-chan bool {
	panic("implement me")
}

func (m MockWriter) Status() int {
	panic("implement me")
}

func (m MockWriter) Size() int {
	panic("implement me")
}

func (m MockWriter) WriteString(string) (int, error) {
	panic("implement me")
}

func (m MockWriter) Written() bool {
	panic("implement me")
}

func (m MockWriter) WriteHeaderNow() {
	panic("implement me")
}

func (m MockWriter) Pusher() http.Pusher {
	panic("implement me")
}
