package controller

import (
	"encoding/json"

	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
)

type SearchAfterPreflight struct{}

type User struct {
	Name string
	Age  string
}

func NewSearchAfterPreflight() *SearchAfterPreflight {
	return &SearchAfterPreflight{}
}

func (c *SearchAfterPreflight) Action(request *http.Request) *http.Response {
	var response *http.Response

	if request.Method == "OPTIONS" {
		response = http.NewResponse(
			http.VersionsFor11,
			http.StatusSuccessCode,
			http.StatusReasonOk,
			request.TargetPath,
			[]byte{},
		)

		response.SetHeader("Access-Control-Allow-Methods", "GET")
		response.SetHeader("Access-Control-Allow-Headers", "Content-Type")
	} else if request.Method == "GET" {
		users := map[string][]User{
			"users": {
				{
					Name: "yuki",
					Age:  "26",
				},
			},
		}

		body, _ := json.Marshal(users)

		response = http.NewResponse(
			http.VersionsFor11,
			http.StatusSuccessCode,
			http.StatusReasonOk,
			request.TargetPath,
			body,
		)
	}

	// コルスの設定
	response.SetHeader("Access-Control-Allow-Origin", "http://localhost:8080")

	return response
}
