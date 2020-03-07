// generated by metactl sdk gen 
package mastodon

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
	deleteStatusesEndpoint := s.opts.Service.GetDeleteStatusesEndpoint()
	getFeedsEndpoint := s.opts.Service.GetGetFeedsEndpoint()
	getPeopleEndpoint := s.opts.Service.GetGetPeopleEndpoint()
	getStatusesEndpoint := s.opts.Service.GetGetStatusesEndpoint()
	postStatusesEndpoint := s.opts.Service.GetPostStatusesEndpoint()
	putPeopleEndpoint := s.opts.Service.GetPutPeopleEndpoint()
	putStatusesEndpoint := s.opts.Service.GetPutStatusesEndpoint()

	return sdk.Service{
		Name: ptr.String(s.opts.Service.Name()),
		Endpoints: &sdk.Endpoints{
			LookupService: &sdk.LookupServiceEndpoint{},
			DeleteStatuses: &deleteStatusesEndpoint,
			GetFeeds: &getFeedsEndpoint,
			GetPeople: &getPeopleEndpoint,
			GetStatuses: &getStatusesEndpoint,
			PostStatuses: &postStatusesEndpoint,
			PutPeople: &putPeopleEndpoint,
			PutStatuses: &putStatusesEndpoint,
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
    case sdk.DeleteStatusesRequestName:
        var req sdk.DeleteStatusesRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.DeleteStatuses(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.GetFeedsRequestName:
        var req sdk.GetFeedsRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.GetFeeds(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.GetPeopleRequestName:
        var req sdk.GetPeopleRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.GetPeople(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.GetStatusesRequestName:
        var req sdk.GetStatusesRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.GetStatuses(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.PostStatusesRequestName:
        var req sdk.PostStatusesRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.PostStatuses(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.PutPeopleRequestName:
        var req sdk.PutPeopleRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.PutPeople(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
    case sdk.PutStatusesRequestName:
        var req sdk.PutStatusesRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.PutStatuses(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
	}
}