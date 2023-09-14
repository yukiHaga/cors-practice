package controller

import (
	"os"
	"path"

	"github.com/yukiHaga/web_server/src/internal/app/config/settings"
	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
)

type PreflightRequest struct{}

func NewPreflightRequest() *PreflightRequest {
	return &PreflightRequest{}
}

func (c *PreflightRequest) Action(request *http.Request) *http.Response {
	STATIC_ROOT, _ := settings.GetStaticRoot()

	body, _ := os.ReadFile(path.Join(STATIC_ROOT, "preflight_request.html"))
	return http.NewResponse(
		http.VersionsFor11,
		http.StatusSuccessCode,
		http.StatusReasonOk,
		request.TargetPath,
		body,
	)
}
