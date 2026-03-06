package api

import "net/http"

func RegisterRoutes(r *Router) {
	r.Use(LoggerMiddleware)

	r.GET("/health", healthCheckHandler)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    "OK",
	})
}
