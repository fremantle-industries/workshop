FROM prom/prometheus:v2.42.0

COPY ./alerts.yml /etc/prometheus/alerts.yml
COPY ./prometheus.local.yml /etc/prometheus/prometheus.local.yml
COPY ./prometheus.development.yml /etc/prometheus/prometheus.development.yml
COPY ./prometheus.staging.yml /etc/prometheus/prometheus.staging.yml
COPY ./prometheus.production.yml /etc/prometheus/prometheus.production.yml
