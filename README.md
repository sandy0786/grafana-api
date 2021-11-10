# grafana-api

Grafana-api is a backend api for building grafana custom datasource plugin.

http://localhost:8080/query?columns=time,hostname&filter={hostname=host2}&start=1636499066000&end=1636518926000

docker build --tag grafana-api .