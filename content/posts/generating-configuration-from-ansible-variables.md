+++
title = "Generating configuration from Ansible variables"
Description = "Personal findings, thoughts & blog"
Tags = [
  "personal",
  "ansible",
]
date = "2016-11-03T08:48:29+02:00"

+++

If you have ever tried to render Ansible multi hash or list variable you probably something alike.

```yaml
global:
  scrape_interval: "{{ prometheus_config_global_scrape_interval | to_nice_yaml }}"
  evaluation_interval: "{{ prometheus_config_global_evaluation_interval | to_nice_yaml }}"
  scrape_timeout: "{{ prometheus_config_global_scrape_timeout | to_nice_yaml }}"
  external_labels: "{{ prometheus_config_global_external_labels | to_nice_yaml }}"
```

This generates invalid and ugly YAML. 

```yaml
global:
  scrape_interval: "15s
...
"
  evaluation_interval: "30s
...
"
  scrape_timeout: "10s
...
"
  external_labels: "null
...
"
```

### TLDR

If you are persistent thus configuration maniac you probably found a way either by destructuring hash or made extra redundant variables around complex one. 
But there is a better way that I came up with.

```yaml
{{
{
'global': {
  'scrape_interval': prometheus_config_global_scrape_interval,
  'evaluation_interval': prometheus_config_global_evaluation_interval,
  'scrape_timeout': prometheus_config_global_scrape_timeout,
  'external_labels': prometheus_config_global_external_labels }
} | to_nice_yaml
}}
```

Here we are using Jinja 2 hash syntax and creating new hash with wanted keys inside of block later piping through `to_nice_yaml` filter as well.
This generates pretty and valid YAML.

```yaml
global:
    evaluation_interval: 30s
    external_labels: null
    scrape_interval: 15s
    scrape_timeout: 10s
```

Same applies to more complex variable definitions like this hash configuration inside of array/list.

```yaml
prometheus_config_scrape_configs:
  - job_name: 'prometheus'
    honor_labels: true
    scrape_interval: '15s'
    scrape_timeout: '2s'
    metrics_path: '/metrics'
    scheme: 'http'
    static_configs:
      - targets:
          - "{{ prometheus_web__listen_address }}" # Prometheus itself
          - "{{ prometheus_alert_manager_web__listen_address }}" # Alert manager

  - job_name: 'consul-services'
    consul_sd_configs:
      - server: "consul.service.consul:8500"
        services:
          - nodeexporter
```

Variable used in template.

```yaml
{% if prometheus_config_scrape_configs is not none and prometheus_config_scrape_configs | length > 0 %}
{{ {'scrape_configs': prometheus_config_scrape_configs} | to_nice_yaml }}
{% endif %}
```

End result here.

```yaml
scrape_configs:
-   honor_labels: true
    job_name: prometheus
    metrics_path: /metrics
    scheme: http
    scrape_interval: 15s
    scrape_timeout: 2s
    static_configs:
    -   targets:
        - 192.168.250.11:9090
        - 192.168.250.11:9093
-   consul_sd_configs:
    -   server: consul.service.consul:8500
        services:
        - nodeexporter
    job_name: consul-services
```

Examples are taken from by ansible-prometheus playbook: https://github.com/ernestas-poskus/ansible-prometheus.
