package cmd

import "github.com/gogf/gf/v2/net/ghttp"

func customOpenApiDoc(s *ghttp.Server) {
	openApi := s.GetOpenApi()
	openApi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openApi.Config.CommonResponseDataField = `Data`

	openApi.Info.Title = "Base Project"
	openApi.Info.Description = ""
}
