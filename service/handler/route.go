package handler

func (ox gatewayHandler) initNsqUsage() {
	router := ox.service.Group("/send")
	{
		router.GET("/nsq", ox.SenderNsq)
	}
}
