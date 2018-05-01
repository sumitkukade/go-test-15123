package servicehandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"

	"gopkg.in/mgo.v2/bson"
)

type UserValidationHandler struct {
}

type GetSurvayTopics struct {
}

func (p UserValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p GetSurvayTopics) Get(r *http.Request) SrvcRes {

	formatted := bson.M{
		"responseData": "heyy",
		"message":      "hello",
		"status":       true}

	data, _ := marshalResponse(formatted)

	return SrvcRes{http.StatusOK, data, "hello", nil}
}
func (p GetSurvayTopics) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
func (p GetSurvayTopics) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p GetSurvayTopics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserValidationHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserValidationHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserValidationHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var cred dao.UserCredentials
	err = json.Unmarshal(body, &cred)

	token := dao.AuthenticateUser(cred)

	if token == "" {
		return UnauthorizedAccess("Bad username or password")
	} else {
		return Response200OK(token)
	}

}
