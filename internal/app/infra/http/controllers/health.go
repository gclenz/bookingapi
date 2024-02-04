package controllers

import (
	"fmt"
	"net/http"
	"time"
)

type HealthController struct {
	startedAt time.Time
}

func (hc *HealthController) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"message": "ok", "uptime":"%v", "time":"%s"}`, time.Since(hc.startedAt), time.Now().UTC())))
}

func NewHealthController(startedAt time.Time) HealthController {
	return HealthController{
		startedAt: startedAt,
	}
}
