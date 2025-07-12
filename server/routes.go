package routes

import (
	"encoding/json"
	"net/http"

	"github.com/wecredit/prometheus-sdk/alerts"
	"github.com/wecredit/prometheus-sdk/config"
	"github.com/wecredit/prometheus-sdk/metrics"
	"github.com/wecredit/prometheus-sdk/utils"
)

// RegisterRoutes sets up HTTP handlers for various endpoints.
func RegisterRoutes() {
	http.HandleFunc("/info", handleInfo)   // Endpoint to log info metrics.
	http.HandleFunc("/error", handleError) // Endpoint to log error metrics and send alerts.
	http.HandleFunc("/health", handleHealth) // Endpoint to check the health of the service.
	http.HandleFunc("/status", handleStatus) // Endpoint to retrieve the current configuration.
}

// handleInfo logs an info metric based on the 'project' and 'event' query parameters.
func handleInfo(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project") // Extract 'project' from query parameters.
	event := r.URL.Query().Get("event")     // Extract 'event' from query parameters.

	if project == "" || event == "" { // Validate required parameters.
		utils.RespondJSON(w, http.StatusBadRequest, "Missing 'project' or 'event'")
		return
	}

	metrics.IncInfo(project, event) // Increment the info metric.
	utils.RespondJSON(w, http.StatusOK, "Info metric logged") // Respond with success message.
}
 
// handleError logs an error metric and optionally sends an email alert.
func handleError(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project") // Extract 'project' from query parameters.
	errType := r.URL.Query().Get("type")    // Extract 'type' from query parameters.

	if project == "" || errType == "" { // Validate required parameters.
		utils.RespondJSON(w, http.StatusBadRequest, "Missing 'project' or 'type'")
		return
	}

	metrics.IncError(project, errType) // Increment the error metric.

	if config.Cfg.AlertEmail.Enable { // Check if email alerts are enabled.
		subject := "[Alert] Error in project " + project // Construct email subject.
		body := "Project " + project + " reported error: " + errType // Construct email body.
		go alerts.SendEmailAlertWithRetry(config.Cfg.AlertEmail, subject, body, 3) // Send email alert asynchronously.
	}

	utils.RespondJSON(w, http.StatusOK, "Error metric logged") // Respond with success message.
}

// handleHealth responds with a simple message indicating the service is running.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, "Prometheus SDK is running") // Respond with health status.
}

// handleStatus responds with the current configuration in JSON format.
func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON.
	json.NewEncoder(w).Encode(config.Cfg) // Encode and send the configuration as JSON.
}
