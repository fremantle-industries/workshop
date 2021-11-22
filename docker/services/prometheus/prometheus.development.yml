global:
  scrape_interval:     15s
  evaluation_interval: 15s
  # scrape_timeout is set to the global default (10s).

  # Attach these labels to any time series or alerts when communicating with
  # external systems (federation, remote storage, Alertmanager).
  # external_labels:
  #     monitor: 'my-project'

rule_files:
  - 'alerts.yml'

# alert
# alerting:
#   alertmanagers:
#   - scheme: http
#     static_configs:
#     - targets:
#       - "alertmanager:9093"

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
         - targets: ['localhost:9090']
