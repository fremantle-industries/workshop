{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "datasource",
          "uid": "grafana"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "id": 4,
  "links": [
    {
      "asDropdown": true,
      "icon": "external link",
      "includeVars": false,
      "keepTime": false,
      "tags": [
        "marketplaces"
      ],
      "targetBlank": false,
      "title": "Marketplaces",
      "tooltip": "",
      "type": "dashboards",
      "url": ""
    },
    {
      "asDropdown": true,
      "icon": "external link",
      "includeVars": false,
      "keepTime": false,
      "tags": [
        "collections"
      ],
      "targetBlank": false,
      "title": "Collections",
      "tooltip": "",
      "type": "dashboards",
      "url": ""
    },
    {
      "asDropdown": false,
      "icon": "external link",
      "includeVars": false,
      "keepTime": false,
      "tags": [
        "tokens"
      ],
      "targetBlank": false,
      "title": "Tokens",
      "tooltip": "",
      "type": "dashboards",
      "url": ""
    },
    {
      "asDropdown": true,
      "icon": "external link",
      "includeVars": false,
      "keepTime": false,
      "tags": [
        "data-processing"
      ],
      "targetBlank": false,
      "title": "Data Processing",
      "tooltip": "",
      "type": "dashboards",
      "url": ""
    },
    {
      "asDropdown": true,
      "icon": "external link",
      "includeVars": false,
      "keepTime": false,
      "tags": [
        "infrastructure"
      ],
      "targetBlank": false,
      "title": "Infrastructure",
      "tooltip": "",
      "type": "dashboards",
      "url": ""
    }
  ],
  "liveNow": false,
  "panels": [
    {
      "datasource": {
        "type": "postgres",
        "uid": "PA8EA412CF102AB70"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "purple",
                "value": null
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 20,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 4,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "pluginVersion": "9.1.0",
      "targets": [
        {
          "datasource": {
            "type": "postgres",
            "uid": "PA8EA412CF102AB70"
          },
          "format": "table",
          "group": [],
          "metricColumn": "none",
          "rawQuery": true,
          "rawSql": "SELECT\n  COUNT(c.*)\nFROM\n  reservoir_collection c;\n",
          "refId": "A",
          "select": [
            [
              {
                "params": [
                  "value"
                ],
                "type": "column"
              }
            ]
          ],
          "timeColumn": "time",
          "where": [
            {
              "name": "$__timeFilter",
              "params": [],
              "type": "macro"
            }
          ]
        }
      ],
      "title": "Token Sets",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "postgres",
        "uid": "PA8EA412CF102AB70"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "purple",
                "value": null
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 20,
        "w": 24,
        "x": 0,
        "y": 20
      },
      "id": 5,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "pluginVersion": "9.1.0",
      "targets": [
        {
          "datasource": {
            "type": "postgres",
            "uid": "PA8EA412CF102AB70"
          },
          "format": "table",
          "group": [],
          "metricColumn": "none",
          "rawQuery": true,
          "rawSql": "SELECT\n  COUNT(a.*)\nFROM\n  reservoir_asset a;\n",
          "refId": "A",
          "select": [
            [
              {
                "params": [
                  "value"
                ],
                "type": "column"
              }
            ]
          ],
          "timeColumn": "time",
          "where": [
            {
              "name": "$__timeFilter",
              "params": [],
              "type": "macro"
            }
          ]
        }
      ],
      "title": "Token Assets",
      "type": "stat"
    }
  ],
  "schemaVersion": 37,
  "style": "dark",
  "tags": [
    "home"
  ],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Home",
  "uid": "ih7wMIY7k",
  "version": 12,
  "weekStart": ""
}
