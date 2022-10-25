// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/runtime/sandbox/v1/sandbox.proto
package sandbox

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
)

type SandboxService interface {
	CreateSandbox(context.Context, *CreateSandboxRequest) (*CreateSandboxResponse, error)
	StartSandbox(context.Context, *StartSandboxRequest) (*StartSandboxResponse, error)
	StopSandbox(context.Context, *StopSandboxRequest) (*StopSandboxResponse, error)
	WaitSandbox(context.Context, *WaitSandboxRequest) (*WaitSandboxResponse, error)
	SandboxStatus(context.Context, *SandboxStatusRequest) (*SandboxStatusResponse, error)
	PingSandbox(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterSandboxService(srv *ttrpc.Server, svc SandboxService) {
	srv.RegisterService("containerd.runtime.sandbox.v1.Sandbox", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"CreateSandbox": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateSandboxRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CreateSandbox(ctx, &req)
			},
			"StartSandbox": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartSandboxRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.StartSandbox(ctx, &req)
			},
			"StopSandbox": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StopSandboxRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.StopSandbox(ctx, &req)
			},
			"WaitSandbox": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitSandboxRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.WaitSandbox(ctx, &req)
			},
			"SandboxStatus": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req SandboxStatusRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.SandboxStatus(ctx, &req)
			},
			"PingSandbox": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PingRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.PingSandbox(ctx, &req)
			},
		},
	})
}

type sandboxClient struct {
	client *ttrpc.Client
}

func NewSandboxClient(client *ttrpc.Client) SandboxService {
	return &sandboxClient{
		client: client,
	}
}

func (c *sandboxClient) CreateSandbox(ctx context.Context, req *CreateSandboxRequest) (*CreateSandboxResponse, error) {
	var resp CreateSandboxResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "CreateSandbox", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *sandboxClient) StartSandbox(ctx context.Context, req *StartSandboxRequest) (*StartSandboxResponse, error) {
	var resp StartSandboxResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "StartSandbox", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *sandboxClient) StopSandbox(ctx context.Context, req *StopSandboxRequest) (*StopSandboxResponse, error) {
	var resp StopSandboxResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "StopSandbox", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *sandboxClient) WaitSandbox(ctx context.Context, req *WaitSandboxRequest) (*WaitSandboxResponse, error) {
	var resp WaitSandboxResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "WaitSandbox", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *sandboxClient) SandboxStatus(ctx context.Context, req *SandboxStatusRequest) (*SandboxStatusResponse, error) {
	var resp SandboxStatusResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "SandboxStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *sandboxClient) PingSandbox(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	var resp PingResponse
	if err := c.client.Call(ctx, "containerd.runtime.sandbox.v1.Sandbox", "PingSandbox", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
