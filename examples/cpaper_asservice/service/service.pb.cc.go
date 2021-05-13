// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: cpaper_asservice/service/service.proto

/*
Package service contains
  *   chaincode interface definition
  *   chaincode gateway definition
  *   chaincode service to cckit router registration func
*/
package service

import (
	context "context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/s7techlab/cckit/examples/cpaper_asservice/schema"
	cckit_gateway "github.com/s7techlab/cckit/gateway"
	cckit_ccservice "github.com/s7techlab/cckit/gateway/service"
	cckit_router "github.com/s7techlab/cckit/router"
	cckit_param "github.com/s7techlab/cckit/router/param"
	cckit_defparam "github.com/s7techlab/cckit/router/param/defparam"
)

// CPaperChaincode  method names
const (
	CPaperChaincode_List = "List"

	CPaperChaincode_Get = "Get"

	CPaperChaincode_GetByExternalId = "GetByExternalId"

	CPaperChaincode_Issue = "Issue"

	CPaperChaincode_Buy = "Buy"

	CPaperChaincode_Redeem = "Redeem"

	CPaperChaincode_Delete = "Delete"
)

// CPaperChaincodeResolver interface for service resolver
type CPaperChaincodeResolver interface {
	CPaperChaincode(ctx cckit_router.Context) (CPaperChaincode, error)
}

// CPaperChaincode chaincode methods interface
type CPaperChaincode interface {
	List(cckit_router.Context, *empty.Empty) (*schema.CommercialPaperList, error)

	Get(cckit_router.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)

	GetByExternalId(cckit_router.Context, *schema.ExternalId) (*schema.CommercialPaper, error)

	Issue(cckit_router.Context, *schema.IssueCommercialPaper) (*schema.CommercialPaper, error)

	Buy(cckit_router.Context, *schema.BuyCommercialPaper) (*schema.CommercialPaper, error)

	Redeem(cckit_router.Context, *schema.RedeemCommercialPaper) (*schema.CommercialPaper, error)

	Delete(cckit_router.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)
}

// RegisterCPaperChaincode registers service methods as chaincode router handlers
func RegisterCPaperChaincode(r *cckit_router.Group, cc CPaperChaincode) error {

	r.Query(CPaperChaincode_List,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.List(ctx, ctx.Param().(*empty.Empty))
		},
		cckit_defparam.Proto(&empty.Empty{}))

	r.Query(CPaperChaincode_Get,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.Get(ctx, ctx.Param().(*schema.CommercialPaperId))
		},
		cckit_defparam.Proto(&schema.CommercialPaperId{}))

	r.Query(CPaperChaincode_GetByExternalId,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.GetByExternalId(ctx, ctx.Param().(*schema.ExternalId))
		},
		cckit_defparam.Proto(&schema.ExternalId{}))

	r.Invoke(CPaperChaincode_Issue,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.Issue(ctx, ctx.Param().(*schema.IssueCommercialPaper))
		},
		cckit_defparam.Proto(&schema.IssueCommercialPaper{}))

	r.Invoke(CPaperChaincode_Buy,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.Buy(ctx, ctx.Param().(*schema.BuyCommercialPaper))
		},
		cckit_defparam.Proto(&schema.BuyCommercialPaper{}))

	r.Invoke(CPaperChaincode_Redeem,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.Redeem(ctx, ctx.Param().(*schema.RedeemCommercialPaper))
		},
		cckit_defparam.Proto(&schema.RedeemCommercialPaper{}))

	r.Invoke(CPaperChaincode_Delete,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.Delete(ctx, ctx.Param().(*schema.CommercialPaperId))
		},
		cckit_defparam.Proto(&schema.CommercialPaperId{}))

	return nil
}

// NewCPaperGateway creates gateway to access chaincode method via chaincode service
func NewCPaperGateway(ccService cckit_ccservice.Chaincode, channel, chaincode string, opts ...cckit_gateway.Opt) *CPaperGateway {
	return &CPaperGateway{Gateway: cckit_gateway.NewChaincode(ccService, channel, chaincode, opts...)}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type CPaperGateway struct {
	Gateway cckit_gateway.Chaincode
}

// ServiceDef returns service definition
func (c *CPaperGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.ServiceDef{
		Desc:                        &_CPaper_serviceDesc,
		Service:                     c,
		HandlerFromEndpointRegister: RegisterCPaperHandlerFromEndpoint,
	}
}

// ApiDef deprecated, use ServiceDef
func (c *CPaperGateway) ApiDef() cckit_gateway.ServiceDef {
	return c.ServiceDef()
}

// Events returns events subscription
func (c *CPaperGateway) Events(ctx context.Context) (cckit_gateway.ChaincodeEventSub, error) {
	return c.Gateway.Events(ctx)
}

func (c *CPaperGateway) List(ctx context.Context, in *empty.Empty) (*schema.CommercialPaperList, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Query(ctx, CPaperChaincode_List, []interface{}{in}, &schema.CommercialPaperList{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaperList), nil
	}
}

func (c *CPaperGateway) Get(ctx context.Context, in *schema.CommercialPaperId) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Query(ctx, CPaperChaincode_Get, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}

func (c *CPaperGateway) GetByExternalId(ctx context.Context, in *schema.ExternalId) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Query(ctx, CPaperChaincode_GetByExternalId, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}

func (c *CPaperGateway) Issue(ctx context.Context, in *schema.IssueCommercialPaper) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Invoke(ctx, CPaperChaincode_Issue, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}

func (c *CPaperGateway) Buy(ctx context.Context, in *schema.BuyCommercialPaper) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Invoke(ctx, CPaperChaincode_Buy, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}

func (c *CPaperGateway) Redeem(ctx context.Context, in *schema.RedeemCommercialPaper) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Invoke(ctx, CPaperChaincode_Redeem, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}

func (c *CPaperGateway) Delete(ctx context.Context, in *schema.CommercialPaperId) (*schema.CommercialPaper, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Invoke(ctx, CPaperChaincode_Delete, []interface{}{in}, &schema.CommercialPaper{}); err != nil {
		return nil, err
	} else {
		return res.(*schema.CommercialPaper), nil
	}
}
