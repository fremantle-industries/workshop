FROM grafana/grafana:7.4.0
# default environment configuration that can be overridden at run time
ENV GRAFANA_DATASOURCE_PROMETHEUS_URL=${GRAFANA_DATASOURCE_PROMETHEUS_URL:-http://prometheus:9090}

COPY ./plugins /var/lib/grafana/plugins/
COPY ./grafana.ini /etc/grafana/grafana.ini
COPY ./provisioning /etc/grafana/provisioning
