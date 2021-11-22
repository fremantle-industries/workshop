apiVersion: 1

deleteDatasources:
  - name: Prometheus
    orgId: 1

datasources:
- name: Prometheus
  type: prometheus
  access: proxy
  orgId: 1
  url: ${GRAFANA_DATASOURCE_PROMETHEUS_URL}
  password:
  user:
  database:
  basicAuth: false
  basicAuthUser:
  basicAuthPassword:
  withCredentials:
  isDefault: true
  jsonData:
     graphiteVersion: "1.1"
     tlsAuth: false
     tlsAuthWithCACert: false
  secureJsonData:
    tlsCACert: "..."
    tlsClientCert: "..."
    tlsClientKey: "..."
  version: 1
  editable: true

- name: Ingest MDC TimescaleDB
  type: postgres
  orgId: 1
  url: ${GRAFANA_DATASOURCE_INGEST_MDC_TIMESCALEDB_URL}
  password:
  user: ${GRAFANA_DATASOURCE_INGEST_MDC_TIMESCALEDB_USER}
  database: ${GRAFANA_DATASOURCE_INGEST_MDC_TIMESCALEDB_DATABASE}
  isDefault: false
  jsonData:
    sslmode: "disable"
    timescaledb: true
  secureJsonData:
    password: ${GRAFANA_DATASOURCE_INGEST_MDC_TIMESCALEDB_PASSWORD}
  version: 1
  editable: true

- name: Marketplaces TimescaleDB
  type: postgres
  orgId: 1
  url: ${GRAFANA_DATASOURCE_MARKETPLACES_TIMESCALEDB_URL}
  password:
  user: ${GRAFANA_DATASOURCE_MARKETPLACES_TIMESCALEDB_USER}
  database: ${GRAFANA_DATASOURCE_MARKETPLACES_TIMESCALEDB_DATABASE}
  isDefault: false
  jsonData:
    sslmode: "disable"
    timescaledb: true
  secureJsonData:
    password: ${GRAFANA_DATASOURCE_MARKETPLACES_TIMESCALEDB_PASSWORD}
  version: 1
  editable: true

- name: Admin TimescaleDB
  type: postgres
  orgId: 1
  url: ${GRAFANA_DATASOURCE_ADMIN_TIMESCALEDB_URL}
  password:
  user: ${GRAFANA_DATASOURCE_ADMIN_TIMESCALEDB_USER}
  database: ${GRAFANA_DATASOURCE_ADMIN_TIMESCALEDB_DATABASE}
  isDefault: false
  jsonData:
    sslmode: "disable"
    timescaledb: true
  secureJsonData:
    password: ${GRAFANA_DATASOURCE_ADMIN_TIMESCALEDB_PASSWORD}
  version: 1
  editable: true

- name: Indexer TimescaleDB
  type: postgres
  orgId: 1
  url: ${GRAFANA_DATASOURCE_INDEXER_TIMESCALEDB_URL}
  password:
  user: ${GRAFANA_DATASOURCE_INDEXER_TIMESCALEDB_USER}
  database: ${GRAFANA_DATASOURCE_INDEXER_TIMESCALEDB_DATABASE}
  isDefault: false
  jsonData:
    sslmode: "disable"
    timescaledb: true
  secureJsonData:
    password: ${GRAFANA_DATASOURCE_INDEXER_TIMESCALEDB_PASSWORD}
  version: 1
  editable: true

- name: Redis
  type: redis-datasource
  access: proxy
  isDefault: false
  orgId: 1
  version: 1
  url: ${GRAFANA_DATASOURCE_REDIS_URL}
  jsonData:
    poolSize: 5
    timeout: 10
    pingInterval: 0
    pipelineWindow: 0
  secureJsonData:
    password: ${GRAFANA_DATASOURCE_REDIS_PASSWORD}
  editable: true
