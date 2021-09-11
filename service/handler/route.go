package handler

func (ox gatewayHandler) initNsqUsage() {
	router := ox.service.Group("/send")
	{
		router.POST("/nsq", ox.SenderNsq)
	}
}
