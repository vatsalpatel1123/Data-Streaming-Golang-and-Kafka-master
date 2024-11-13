package handlers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"streaming-api/monitoring"
	"streaming-api/services"
	"streaming-api/utils"
)

// POST /stream/start
func StartStream(w http.ResponseWriter, r *http.Request) {
	monitoring.RequestsProcessed.WithLabelValues("start_stream").Inc()
	// Start the stream
}

// POST /stream/{stream_id}/send
func SendStreamData(w http.ResponseWriter, r *http.Request) {
	monitoring.RequestsProcessed.WithLabelValues("send_stream_data").Inc()
	// Send stream data
}

// GET /stream/{stream_id}/results
func StreamResults(w http.ResponseWriter, r *http.Request) {
	monitoring.RequestsProcessed.WithLabelValues("stream_results").Inc()
	// Stream results via WebSocket
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func StreamResults(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Logger.Error("WebSocket upgrade failed", zap.Error(err))
		return
	}
	defer conn.Close()

	streamID := mux.Vars(r)["stream_id"]
	go services.ConsumeKafkaStream(conn, streamID)
}
