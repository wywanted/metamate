package funcs

import (
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

func CollectSvcRsps(ctx types.ReqCtx, ctxs []types.ReqCtx) types.ReqCtx {
	for _, ctx0 := range ctxs {
		ctx.GSvcRsps = append(ctx.GSvcRsps, ctx0.GSvcRsp)
	}

	return ctx
}

func Collect(from, to string) func(ctx types.ReqCtx, ctxs []types.ReqCtx) types.ReqCtx {
	switch from {
	case types.GEntity:
		switch to {
		case types.GCliRsp:
			return func(ctx types.ReqCtx, ctxs []types.ReqCtx) types.ReqCtx {
				gs := []generic.Generic{}
				for _, ctx0 := range ctxs {
					gs = append(gs, ctx0.GEntity)
					ctx.Errs = append(ctx.Errs, ctx0.Errs...)
				}

				gSlice, ok := ctx.GCliRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
				if ok {
					gSlice.Set(gs)
					ctx.GCliRsp.MustSetGenericSlice([]string{ctx.ForTypeNode.PluralFieldName()}, gSlice)
				}

				return ctx
			}
		case types.GSvcRsp:
			return func(ctx types.ReqCtx, ctxs []types.ReqCtx) types.ReqCtx {
				gs := []generic.Generic{}
				for _, ctx0 := range ctxs {
					gs = append(gs, ctx0.GEntity)
					ctx.Errs = append(ctx.Errs, ctx0.Errs...)
				}

				gSlice, ok := ctx.GSvcRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
				if ok {
					gSlice.Set(gs)
					ctx.GSvcRsp.MustSetGenericSlice([]string{ctx.ForTypeNode.PluralFieldName()}, gSlice)
				}

				return ctx
			}
		}
	case types.GSvcRsp:
		switch to {
		case types.GEntity:

		}
	}

	panic("ho")
}

func Map(from, to string) func(ctx types.ReqCtx) (ctxs []types.ReqCtx) {
	switch from {
	case types.GCliRsp:
		switch to {
		case types.GEntity:
			return func(ctx types.ReqCtx) (ctxs []types.ReqCtx) {
				gSlice, ok := ctx.GCliRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
				if !ok {
					return
				}

				gs := gSlice.Get()
				for i, _ := range gs {
					ctxs = append(ctxs, types.ReqCtx{
						GCliReq:       ctx.GCliReq,
						GEntity:       gs[i],
						ForTypeNode:   ctx.ForTypeNode,
					})
				}

				return
			}
		}
	case types.GSvcRsp:
		switch to {
		case types.GEntity:
			return func(ctx types.ReqCtx) (ctxs []types.ReqCtx) {
				gSlice, ok := ctx.GSvcRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
				if !ok {
					return
				}

				gs := gSlice.Get()
				for i, _ := range gs {
					ctxs = append(ctxs, types.ReqCtx{
						GSvcReq:       ctx.GSvcReq.Copy(),
						GEntity:       gs[i].Copy(),
						ForTypeNode:   ctx.ForTypeNode,
					})
				}

				return
			}
		}
	case types.Svcs:
		switch to {
		case types.Svc:
			return func(ctx types.ReqCtx) (ctxs []types.ReqCtx) {
				if ctx.Svcs == nil {
					return
				}

				for i, _ := range ctx.Svcs {
					ctxs = append(ctxs, types.ReqCtx{
						Stage:       ctx.Stage,
						GCliReq:     ctx.GCliReq,
						Svc:         &ctx.Svcs[i],
						ForTypeNode: ctx.ForTypeNode,
					})
				}

				return
			}
		}
	}

	panic("h")
}
