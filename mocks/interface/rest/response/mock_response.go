package mock_response

import (
	"net/http"

	"github.com/riyanda432/belajar-authentication/src/interface/rest/response"
	"github.com/stretchr/testify/mock"
)

type MockResponseIResponseClient struct {
	mock.Mock
}

func (m *MockResponseIResponseClient) MockBuildMeta(res *response.Meta) {
	m.Mock.On("BuildMeta", mock.Anything, mock.Anything, mock.Anything).Return(res)
}

// BuildMeta implements response.IResponseClient
func (m *MockResponseIResponseClient) BuildMeta(page int, perPage int, count int64) *response.Meta {
	args := m.Called(page, perPage, count)
	var r *response.Meta
	if n, ok := args.Get(0).(*response.Meta); ok {
		r = n
	}
	return r
}

func (m *MockResponseIResponseClient) MockHttpErrorSuccess() {
	m.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
}

// HttpError implements response.IResponseClient
func (m *MockResponseIResponseClient) HttpError(w http.ResponseWriter, err error) error {
	args := m.Called(w, err)
	var r error
	if n, ok := args.Get(0).(error); ok {
		r = n
	}
	return r
}

func (m *MockResponseIResponseClient) MockJSONSuccess() {
	m.Mock.On(
		"JSON",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
}

// JSON implements response.IResponseClient
func (m *MockResponseIResponseClient) JSON(
	w http.ResponseWriter,
	message string,
	data interface{},
	meta *response.Meta,
) error {
	args := m.Called(w, message, data, meta)
	var r error
	if n, ok := args.Get(0).(error); ok {
		r = n
	}
	return r
}

var _ response.IResponseClient = &MockResponseIResponseClient{}
