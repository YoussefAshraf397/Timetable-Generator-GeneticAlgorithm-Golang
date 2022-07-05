package Application

import (
	"github.com/bykovme/gotrans"
)

func (req Request) OK(body interface{}) {
	prepareResponse := buildResponse(body, gotrans.T("ok"), 200, nil)
	req.Response(200, prepareResponse)
}

func (req Request) Created(body interface{}) {
	prepareResponse := buildResponse(body, gotrans.T("Created_suc"), 201, nil)
	req.Response(201, prepareResponse)

}

func (req Request) NoAuth() {
	prepareResponse := buildResponse(nil, gotrans.T("no_auth"), 401, nil)
	req.Response(401, prepareResponse)
}

func (req Request) BadRequest(err interface{}) {
	prepareResponse := buildResponse(nil, gotrans.T("something_wrong"), 422, err)
	req.Response(422, prepareResponse)
}

func (req Request) UserNotFound() {
	prepareResponse := buildResponse(nil, gotrans.T("Not_Found"), 404, nil)
	req.Response(404, prepareResponse)
}

func buildResponse(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})

	response["payload"] = payload
	response["message"] = message
	response["code"] = code
	response["errors"] = err
	return response
}
