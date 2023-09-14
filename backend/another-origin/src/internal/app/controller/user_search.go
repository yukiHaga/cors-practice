package controller

import (
	"os"
	"path"

	"github.com/yukiHaga/web_server/src/internal/app/config/settings"
	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
)

type UserSearch struct{}

func NewUserSearch() *UserSearch {
	return &UserSearch{}
}

func (c *UserSearch) Action(request *http.Request) *http.Response {
	STATIC_ROOT, _ := settings.GetStaticRoot()

	body, _ := os.ReadFile(path.Join(STATIC_ROOT, "user_search_result.html"))

	response := http.NewResponse(
		http.VersionsFor11,
		http.StatusSuccessCode,
		http.StatusReasonOk,
		request.TargetPath,
		body,
	)

	// コルスの設定
	response.SetHeader("Access-Control-Allow-Origin", "http://localhost:8080")

	return response
}
