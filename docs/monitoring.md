# Monitoring with Prometheus and Grafana

1. **Prometheus Setup**:
   - Install Prometheus: https://prometheus.io/download/
   - Start Prometheus: `./prometheus --config.file=prometheus.yml`
   - Access at: `http://localhost:9090/`

2. **Grafana Setup**:
   - Install Grafana: https://grafana.com/grafana/download
   - Add Prometheus as a data source.
   - Import Grafana dashboard for Kafka and API metrics.
