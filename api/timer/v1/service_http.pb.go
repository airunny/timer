// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.3
// source: timer/v1/service.proto

package v1

import (
	context "context"
	common "github.com/airunny/timer/api/common"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceAddApplication = "/api.timer.v1.Service/AddApplication"
const OperationServiceAddTimer = "/api.timer.v1.Service/AddTimer"
const OperationServiceAddUser = "/api.timer.v1.Service/AddUser"
const OperationServiceDeleteApplication = "/api.timer.v1.Service/DeleteApplication"
const OperationServiceDeleteTimer = "/api.timer.v1.Service/DeleteTimer"
const OperationServiceDeleteUser = "/api.timer.v1.Service/DeleteUser"
const OperationServiceGenApplicationSecret = "/api.timer.v1.Service/GenApplicationSecret"
const OperationServiceGetApplication = "/api.timer.v1.Service/GetApplication"
const OperationServiceGetTimer = "/api.timer.v1.Service/GetTimer"
const OperationServiceGetUser = "/api.timer.v1.Service/GetUser"
const OperationServiceHealthy = "/api.timer.v1.Service/Healthy"
const OperationServiceListApplication = "/api.timer.v1.Service/ListApplication"
const OperationServiceListTimer = "/api.timer.v1.Service/ListTimer"
const OperationServiceListTimerCallback = "/api.timer.v1.Service/ListTimerCallback"
const OperationServiceListUser = "/api.timer.v1.Service/ListUser"
const OperationServiceLogin = "/api.timer.v1.Service/Login"
const OperationServiceRefreshToken = "/api.timer.v1.Service/RefreshToken"
const OperationServiceReplayTimer = "/api.timer.v1.Service/ReplayTimer"
const OperationServiceUpdateApplication = "/api.timer.v1.Service/UpdateApplication"
const OperationServiceUpdateApplicationStatus = "/api.timer.v1.Service/UpdateApplicationStatus"
const OperationServiceUpdatePassword = "/api.timer.v1.Service/UpdatePassword"
const OperationServiceUpdateTimerStatus = "/api.timer.v1.Service/UpdateTimerStatus"
const OperationServiceUpdateUserPassword = "/api.timer.v1.Service/UpdateUserPassword"
const OperationServiceUpdateUserStatus = "/api.timer.v1.Service/UpdateUserStatus"

type ServiceHTTPServer interface {
	// AddApplication 应用[添加]
	AddApplication(context.Context, *AddApplicationRequest) (*Application, error)
	// AddTimer 定时器[添加]
	AddTimer(context.Context, *AddTimerRequest) (*AddTimerReply, error)
	// AddUser 用户[添加]
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	// DeleteApplication 应用[删除]
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*DeleteApplicationReply, error)
	// DeleteTimer 定时器[移除]
	DeleteTimer(context.Context, *DeleteTimerRequest) (*DeleteTimerReply, error)
	// DeleteUser 用户[删除]
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// GenApplicationSecret 应用[重新生成秘钥]
	GenApplicationSecret(context.Context, *GenApplicationSecretRequest) (*GenApplicationSecretReply, error)
	// GetApplication 应用[详情]
	GetApplication(context.Context, *GetApplicationRequest) (*Application, error)
	// GetTimer 定时器[详情]
	GetTimer(context.Context, *GetTimerRequest) (*Timer, error)
	// GetUser 用户[详情]
	GetUser(context.Context, *GetUserRequest) (*User, error)
	Healthy(context.Context, *common.EmptyRequest) (*common.HealthyReply, error)
	// ListApplication 应用[列表]
	ListApplication(context.Context, *ListApplicationRequest) (*ListApplicationReply, error)
	// ListTimer 定时器[列表]
	ListTimer(context.Context, *ListTimerRequest) (*ListTimerReply, error)
	// ListTimerCallback 定时器[回调列表]
	ListTimerCallback(context.Context, *ListTimerCallbackRequest) (*ListTimerCallbackReply, error)
	// ListUser 用户[列表]
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// RefreshToken 刷新登录token
	RefreshToken(context.Context, *RefreshTokenRequest) (*LoginReply, error)
	// ReplayTimer 定时器[重放]
	ReplayTimer(context.Context, *ReplayTimerRequest) (*ReplayTimerReply, error)
	// UpdateApplication 应用[更新]
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*UpdateApplicationReply, error)
	// UpdateApplicationStatus 应用[更新状态]
	UpdateApplicationStatus(context.Context, *UpdateApplicationStatusRequest) (*UpdateApplicationStatusReply, error)
	// UpdatePassword 更新密码，用户自己
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error)
	// UpdateTimerStatus 定时器[更新状态]
	UpdateTimerStatus(context.Context, *UpdateTimerStatusRequest) (*UpdateTimerStatusReply, error)
	// UpdateUserPassword 用户[更新密码]
	UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*UpdateUserPasswordReply, error)
	// UpdateUserStatus 用户[更新状态]
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusReply, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/healthz", _Service_Healthy0_HTTP_Handler(srv))
	r.POST("/v1/login", _Service_Login0_HTTP_Handler(srv))
	r.GET("/v1/refresh", _Service_RefreshToken0_HTTP_Handler(srv))
	r.PUT("/v1/password", _Service_UpdatePassword0_HTTP_Handler(srv))
	r.POST("/v1/app", _Service_AddApplication0_HTTP_Handler(srv))
	r.DELETE("/v1/app/{id}", _Service_DeleteApplication0_HTTP_Handler(srv))
	r.PUT("/v1/app/{id}", _Service_UpdateApplication0_HTTP_Handler(srv))
	r.PUT("/v1/app/{id}/status", _Service_UpdateApplicationStatus0_HTTP_Handler(srv))
	r.GET("/v1/app/{id}", _Service_GetApplication0_HTTP_Handler(srv))
	r.PUT("/v1/app/{id}/secret", _Service_GenApplicationSecret0_HTTP_Handler(srv))
	r.GET("/v1/app", _Service_ListApplication0_HTTP_Handler(srv))
	r.POST("/v1/timer", _Service_AddTimer0_HTTP_Handler(srv))
	r.GET("/v1/timer/{id}", _Service_GetTimer0_HTTP_Handler(srv))
	r.PUT("/v1/timer/{id}/status", _Service_UpdateTimerStatus0_HTTP_Handler(srv))
	r.DELETE("/v1/timer/{id}", _Service_DeleteTimer0_HTTP_Handler(srv))
	r.GET("/v1/timer/{id}/replay", _Service_ReplayTimer0_HTTP_Handler(srv))
	r.GET("/v1/timer", _Service_ListTimer0_HTTP_Handler(srv))
	r.GET("/v1/timer/{id}/callback", _Service_ListTimerCallback0_HTTP_Handler(srv))
	r.POST("/v1/user", _Service_AddUser0_HTTP_Handler(srv))
	r.GET("/v1/user/{id}", _Service_GetUser0_HTTP_Handler(srv))
	r.PUT("/v1/user/{id}/status", _Service_UpdateUserStatus0_HTTP_Handler(srv))
	r.PUT("/v1/user/{id}/password", _Service_UpdateUserPassword0_HTTP_Handler(srv))
	r.GET("/v1/user", _Service_ListUser0_HTTP_Handler(srv))
	r.DELETE("/v1/user/{id}", _Service_DeleteUser0_HTTP_Handler(srv))
}

func _Service_Healthy0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in common.EmptyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceHealthy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Healthy(ctx, req.(*common.EmptyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.HealthyReply)
		return ctx.Result(200, reply)
	}
}

func _Service_Login0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Service_RefreshToken0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdatePassword0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdatePassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePassword(ctx, req.(*UpdatePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePasswordReply)
		return ctx.Result(200, reply)
	}
}

func _Service_AddApplication0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddApplicationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddApplication(ctx, req.(*AddApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Application)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteApplication0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteApplicationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteApplication(ctx, req.(*DeleteApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteApplicationReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateApplication0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateApplicationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApplication(ctx, req.(*UpdateApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateApplicationReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateApplicationStatus0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateApplicationStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateApplicationStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApplicationStatus(ctx, req.(*UpdateApplicationStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateApplicationStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Service_GetApplication0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetApplicationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApplication(ctx, req.(*GetApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Application)
		return ctx.Result(200, reply)
	}
}

func _Service_GenApplicationSecret0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenApplicationSecretRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGenApplicationSecret)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenApplicationSecret(ctx, req.(*GenApplicationSecretRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenApplicationSecretReply)
		return ctx.Result(200, reply)
	}
}

func _Service_ListApplication0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListApplicationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceListApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListApplication(ctx, req.(*ListApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListApplicationReply)
		return ctx.Result(200, reply)
	}
}

func _Service_AddTimer0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddTimerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddTimer(ctx, req.(*AddTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddTimerReply)
		return ctx.Result(200, reply)
	}
}

func _Service_GetTimer0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTimerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTimer(ctx, req.(*GetTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Timer)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateTimerStatus0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTimerStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateTimerStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTimerStatus(ctx, req.(*UpdateTimerStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTimerStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteTimer0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTimerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTimer(ctx, req.(*DeleteTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTimerReply)
		return ctx.Result(200, reply)
	}
}

func _Service_ReplayTimer0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReplayTimerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceReplayTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReplayTimer(ctx, req.(*ReplayTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReplayTimerReply)
		return ctx.Result(200, reply)
	}
}

func _Service_ListTimer0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTimerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceListTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTimer(ctx, req.(*ListTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTimerReply)
		return ctx.Result(200, reply)
	}
}

func _Service_ListTimerCallback0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTimerCallbackRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceListTimerCallback)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTimerCallback(ctx, req.(*ListTimerCallbackRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTimerCallbackReply)
		return ctx.Result(200, reply)
	}
}

func _Service_AddUser0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddUser(ctx, req.(*AddUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddUserReply)
		return ctx.Result(200, reply)
	}
}

func _Service_GetUser0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateUserStatus0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateUserStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateUserPassword0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserPassword(ctx, req.(*UpdateUserPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserPasswordReply)
		return ctx.Result(200, reply)
	}
}

func _Service_ListUser0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteUser0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	AddApplication(ctx context.Context, req *AddApplicationRequest, opts ...http.CallOption) (rsp *Application, err error)
	AddTimer(ctx context.Context, req *AddTimerRequest, opts ...http.CallOption) (rsp *AddTimerReply, err error)
	AddUser(ctx context.Context, req *AddUserRequest, opts ...http.CallOption) (rsp *AddUserReply, err error)
	DeleteApplication(ctx context.Context, req *DeleteApplicationRequest, opts ...http.CallOption) (rsp *DeleteApplicationReply, err error)
	DeleteTimer(ctx context.Context, req *DeleteTimerRequest, opts ...http.CallOption) (rsp *DeleteTimerReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	GenApplicationSecret(ctx context.Context, req *GenApplicationSecretRequest, opts ...http.CallOption) (rsp *GenApplicationSecretReply, err error)
	GetApplication(ctx context.Context, req *GetApplicationRequest, opts ...http.CallOption) (rsp *Application, err error)
	GetTimer(ctx context.Context, req *GetTimerRequest, opts ...http.CallOption) (rsp *Timer, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *User, err error)
	Healthy(ctx context.Context, req *common.EmptyRequest, opts ...http.CallOption) (rsp *common.HealthyReply, err error)
	ListApplication(ctx context.Context, req *ListApplicationRequest, opts ...http.CallOption) (rsp *ListApplicationReply, err error)
	ListTimer(ctx context.Context, req *ListTimerRequest, opts ...http.CallOption) (rsp *ListTimerReply, err error)
	ListTimerCallback(ctx context.Context, req *ListTimerCallbackRequest, opts ...http.CallOption) (rsp *ListTimerCallbackReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	ReplayTimer(ctx context.Context, req *ReplayTimerRequest, opts ...http.CallOption) (rsp *ReplayTimerReply, err error)
	UpdateApplication(ctx context.Context, req *UpdateApplicationRequest, opts ...http.CallOption) (rsp *UpdateApplicationReply, err error)
	UpdateApplicationStatus(ctx context.Context, req *UpdateApplicationStatusRequest, opts ...http.CallOption) (rsp *UpdateApplicationStatusReply, err error)
	UpdatePassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *UpdatePasswordReply, err error)
	UpdateTimerStatus(ctx context.Context, req *UpdateTimerStatusRequest, opts ...http.CallOption) (rsp *UpdateTimerStatusReply, err error)
	UpdateUserPassword(ctx context.Context, req *UpdateUserPasswordRequest, opts ...http.CallOption) (rsp *UpdateUserPasswordReply, err error)
	UpdateUserStatus(ctx context.Context, req *UpdateUserStatusRequest, opts ...http.CallOption) (rsp *UpdateUserStatusReply, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) AddApplication(ctx context.Context, in *AddApplicationRequest, opts ...http.CallOption) (*Application, error) {
	var out Application
	pattern := "/v1/app"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) AddTimer(ctx context.Context, in *AddTimerRequest, opts ...http.CallOption) (*AddTimerReply, error) {
	var out AddTimerReply
	pattern := "/v1/timer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) AddUser(ctx context.Context, in *AddUserRequest, opts ...http.CallOption) (*AddUserReply, error) {
	var out AddUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...http.CallOption) (*DeleteApplicationReply, error) {
	var out DeleteApplicationReply
	pattern := "/v1/app/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...http.CallOption) (*DeleteTimerReply, error) {
	var out DeleteTimerReply
	pattern := "/v1/timer/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GenApplicationSecret(ctx context.Context, in *GenApplicationSecretRequest, opts ...http.CallOption) (*GenApplicationSecretReply, error) {
	var out GenApplicationSecretReply
	pattern := "/v1/app/{id}/secret"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceGenApplicationSecret))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...http.CallOption) (*Application, error) {
	var out Application
	pattern := "/v1/app/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetTimer(ctx context.Context, in *GetTimerRequest, opts ...http.CallOption) (*Timer, error) {
	var out Timer
	pattern := "/v1/timer/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) Healthy(ctx context.Context, in *common.EmptyRequest, opts ...http.CallOption) (*common.HealthyReply, error) {
	var out common.HealthyReply
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceHealthy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) ListApplication(ctx context.Context, in *ListApplicationRequest, opts ...http.CallOption) (*ListApplicationReply, error) {
	var out ListApplicationReply
	pattern := "/v1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceListApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) ListTimer(ctx context.Context, in *ListTimerRequest, opts ...http.CallOption) (*ListTimerReply, error) {
	var out ListTimerReply
	pattern := "/v1/timer"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceListTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) ListTimerCallback(ctx context.Context, in *ListTimerCallbackRequest, opts ...http.CallOption) (*ListTimerCallbackReply, error) {
	var out ListTimerCallbackReply
	pattern := "/v1/timer/{id}/callback"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceListTimerCallback))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/refresh"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) ReplayTimer(ctx context.Context, in *ReplayTimerRequest, opts ...http.CallOption) (*ReplayTimerReply, error) {
	var out ReplayTimerReply
	pattern := "/v1/timer/{id}/replay"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceReplayTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...http.CallOption) (*UpdateApplicationReply, error) {
	var out UpdateApplicationReply
	pattern := "/v1/app/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateApplicationStatus(ctx context.Context, in *UpdateApplicationStatusRequest, opts ...http.CallOption) (*UpdateApplicationStatusReply, error) {
	var out UpdateApplicationStatusReply
	pattern := "/v1/app/{id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateApplicationStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...http.CallOption) (*UpdatePasswordReply, error) {
	var out UpdatePasswordReply
	pattern := "/v1/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdatePassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateTimerStatus(ctx context.Context, in *UpdateTimerStatusRequest, opts ...http.CallOption) (*UpdateTimerStatusReply, error) {
	var out UpdateTimerStatusReply
	pattern := "/v1/timer/{id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateTimerStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...http.CallOption) (*UpdateUserPasswordReply, error) {
	var out UpdateUserPasswordReply
	pattern := "/v1/user/{id}/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...http.CallOption) (*UpdateUserStatusReply, error) {
	var out UpdateUserStatusReply
	pattern := "/v1/user/{id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateUserStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
