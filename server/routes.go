package routes

import (
	"encoding/json"
	"net/http"

	"github.com/wecredit/prometheus-sdk/alerts"
	"github.com/wecredit/prometheus-sdk/config"
	"github.com/wecredit/prometheus-sdk/metrics"
	"github.com/wecredit/prometheus-sdk/utils"
)

func RegisterRoutes() {
	http.HandleFunc("/info", handleInfo)
	http.HandleFunc("/error", handleError)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/status", handleStatus)
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")
	event := r.URL.Query().Get("event")

	if project == "" || event == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "Missing 'project' or 'event'")
		return
	}

	metrics.IncInfo(project, event)
	utils.RespondJSON(w, http.StatusOK, "Info metric logged")
}

func handleError(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")
	errType := r.URL.Query().Get("type")

	if project == "" || errType == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "Missing 'project' or 'type'")
		return
	}

	metrics.IncError(project, errType)

	if config.Cfg.AlertEmail.Enable {
		subject := "[Alert] Error in project " + project
		body := "Project " + project + " reported error: " + errType
		go alerts.SendEmailAlertWithRetry(config.Cfg.AlertEmail, subject, body, 3)
	}

	utils.RespondJSON(w, http.StatusOK, "Error metric logged")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, "Prometheus SDK is running")
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Cfg)
}
