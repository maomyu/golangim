package contextcli

import "context"

const Name = "addr"

const addrKey = "addr"

func ContextWithAddr(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, addrKey, addr)
}
