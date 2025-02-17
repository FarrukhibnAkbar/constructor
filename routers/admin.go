package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/system/v1/")

	adminGroup.POST("/API", r.handler.CreateAPI)
	adminGroup.POST("/API/details", r.handler.CreateAPIDetails)
	adminGroup.Any("/core/*api_name", r.handler.ExecuteAPI)
}
