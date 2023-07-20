package client

import (
	types "github.com/thnkrn/comet/puller/pkg/types"
)

type HTTP interface {
	Get(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error
	Put(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error
	Post(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error
}
