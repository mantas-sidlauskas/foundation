// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package logging

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/go-ids"
)

var Log = log.New(os.Stderr, "[grpclog] ", log.Ldate|log.Ltime|log.Lmicroseconds)

type interceptor struct{}

func (interceptor) unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	t := time.Now()
	reqid := logHeader(ctx, "request", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err == nil {
		Log.Printf("%s: id=%s: took %v; response: %+v", info.FullMethod, reqid, time.Since(t), resp)
	} else {
		Log.Printf("%s: id=%s: took %v; error: %v", info.FullMethod, reqid, time.Since(t), err)
	}
	return resp, err
}

func (interceptor) streaming(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	t := time.Now()
	reqid := logHeader(stream.Context(), "stream", info.FullMethod, nil)
	err := handler(srv, stream)
	if err == nil {
		Log.Printf("%s: id=%s: took %v, finished ok", info.FullMethod, reqid, time.Since(t))
	} else {
		Log.Printf("%s: id=%s: took %v; error: %v", info.FullMethod, reqid, time.Since(t), err)
	}
	return err
}

func logHeader(ctx context.Context, what, fullMethod string, req interface{}) string {
	// XXX establish request id propagation.
	reqid := ids.NewRandomBase32ID(8)
	peerAddr := "unknown"
	authType := "none"
	if p, has := peer.FromContext(ctx); has {
		peerAddr = p.Addr.String()
		if p.AuthInfo != nil {
			authType = p.AuthInfo.AuthType()
		}
	}

	if req != nil {
		core.Log.Printf("%s: id=%s: request from %s (auth: %s): %+v", fullMethod, reqid, peerAddr, authType, req)
	} else {
		core.Log.Printf("%s: id=%s: request from %s (auth: %s)", fullMethod, reqid, peerAddr, authType)
	}
	return reqid
}

func Prepare(ctx context.Context, deps ExtensionDeps) error {
	var interceptor interceptor
	deps.Interceptors.Add(interceptor.unary, interceptor.streaming)
	return nil
}
