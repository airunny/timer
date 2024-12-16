package icontext

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
)

const (
	userIdKey string = "UserId" // 用户ID
	roleKey   string = "RoleId" // 角色
)

func withValue(ctx context.Context, key, value string) context.Context {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	md.Set(key, value)
	return metadata.NewServerContext(ctx, md)
}

func fromValue(ctx context.Context, key string) (string, bool) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return "", false
	}

	out := md.Get(key)
	return out, out != ""
}

func WithUserId(ctx context.Context, in string) context.Context {
	return withValue(ctx, userIdKey, in)
}

func UserIdFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, userIdKey)
}

func WithRole(ctx context.Context, in string) context.Context {
	return withValue(ctx, roleKey, in)
}

func RoleFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, roleKey)
}
