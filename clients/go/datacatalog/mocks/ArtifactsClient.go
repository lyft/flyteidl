// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	datacatalog "github.com/lyft/flyteidl/gen/pb-go/flyteidl/datacatalog"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ArtifactsClient is an autogenerated mock type for the ArtifactsClient type
type ArtifactsClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactsClient) Create(ctx context.Context, in *datacatalog.CreateRequest, opts ...grpc.CallOption) (*datacatalog.CreateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.CreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateRequest, ...grpc.CallOption) *datacatalog.CreateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateProvenance provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactsClient) GenerateProvenance(ctx context.Context, in *datacatalog.GenerateProvenanceRequest, opts ...grpc.CallOption) (*datacatalog.GenerateProvenanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GenerateProvenanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GenerateProvenanceRequest, ...grpc.CallOption) *datacatalog.GenerateProvenanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GenerateProvenanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GenerateProvenanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactsClient) Get(ctx context.Context, in *datacatalog.GetRequest, opts ...grpc.CallOption) (*datacatalog.GetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetRequest, ...grpc.CallOption) *datacatalog.GetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactsClient) Query(ctx context.Context, in *datacatalog.QueryRequest, opts ...grpc.CallOption) (*datacatalog.QueryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.QueryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.QueryRequest, ...grpc.CallOption) *datacatalog.QueryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.QueryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.QueryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
