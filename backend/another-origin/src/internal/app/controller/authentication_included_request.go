package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
)

type AuthenticationIncludedRequest struct{}

type Counter struct {
	CurrentCount int
}

func NewAuthenticationIncludedRequest() *AuthenticationIncludedRequest {
	return &AuthenticationIncludedRequest{}
}

func (c *AuthenticationIncludedRequest) Action(request *http.Request) *http.Response {
	cookieHeaders := map[string]string{}
	cookie, isThere := request.Cookies["counter"]
	var response *http.Response

	if isThere {
		currentCount, _ := strconv.Atoi(cookie.Value)
		currentCount += 1
		cookieHeaders["counter"] = fmt.Sprintf("%v", currentCount)

		counter := Counter{
			CurrentCount: currentCount,
		}

		body, _ := json.Marshal(counter)

		response = http.NewResponse(
			http.VersionsFor11,
			http.StatusSuccessCode,
			http.StatusReasonOk,
			request.TargetPath,
			[]byte(body),
		)
	} else {
		currentCount := 1
		cookieHeaders["counter"] = fmt.Sprintf("%v", currentCount)

		counter := Counter{
			CurrentCount: currentCount,
		}

		body, _ := json.Marshal(counter)

		response = http.NewResponse(
			http.VersionsFor11,
			http.StatusSuccessCode,
			http.StatusReasonOk,
			request.TargetPath,
			[]byte(body),
		)
	}

	// クッキーの設定
	for key, value := range cookieHeaders {
		response.SetCookieHeader(fmt.Sprintf("%s=%s", key, value))
	}

	// コンテンツタイプの設定
	response.SetContentTypeHeader("application/json")

	// コルスの設定
	response.SetHeader("Access-Control-Allow-Origin", "http://localhost:8080")
	// 認証情報を送るから、ヘッダもセットしてお以下のヘッダもセットしておく
	response.SetHeader("Access-Control-Allow-Credentials", "true")

	return response
}
