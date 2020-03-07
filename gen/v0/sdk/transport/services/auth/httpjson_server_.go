// generated by metactl sdk gen 
package auth

import (
	"encoding/json"
	"net/http"
	"reflect"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport"
)

type HttpJsonServer struct {
	opts HttpJsonServerOpts
}

type HttpJsonServerOpts struct {
	Service Service
}

func NewHttpJsonServer(opts HttpJsonServerOpts) (http.Handler) {
	return HttpJsonServer{opts: opts}
}

func (s HttpJsonServer) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(transport.CONTENT_TYPE_HEADER, transport.CONTENT_TYPE_JSON)
	w.Header().Set(transport.METAMATE_TYPE_HEADER, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
	    return
	}

	return
}

func (s HttpJsonServer) getService() (sdk.Service) {
	authenticateClientAccountEndpoint := s.opts.Service.GetAuthenticateClientAccountEndpoint()
	pipeClientAccountsEndpoint := s.opts.Service.GetPipeClientAccountsEndpoint()
	verifyTokenEndpoint := s.opts.Service.GetVerifyTokenEndpoint()

	return sdk.Service{
		Name: ptr.String(s.opts.Service.Name()),
		Endpoints: &sdk.Endpoints{
			LookupService: &sdk.LookupServiceEndpoint{},
			AuthenticateClientAccount: &authenticateClientAccountEndpoint,
			PipeClientAccounts: &pipeClientAccountsEndpoint,
			VerifyToken: &verifyTokenEndpoint,
		},
	}
}

func (s HttpJsonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(transport.METAMATE_TYPE_HEADER) {
	case sdk.LookupServiceRequestName:
			var req sdk.LookupServiceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				return
			}
	
			svc := s.getService()
			rsp := sdk.LookupServiceResponse{
				Output: &sdk.LookupServiceOutput{
					Service: &svc,
				},
			}
	
			err = s.send(w, rsp)
			if err != nil {
				return
			}
    case sdk.AuthenticateClientAccountRequestName:
        var req sdk.AuthenticateClientAccountRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.AuthenticateClientAccount(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.PipeClientAccountsRequestName:
        var req sdk.PipeClientAccountsRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.PipeClientAccounts(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.VerifyTokenRequestName:
        var req sdk.VerifyTokenRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.VerifyToken(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
	}
}