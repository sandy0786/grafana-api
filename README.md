# grafana-api

Grafana-api is a backend api for building grafana custom datasource plugin.

## Running the application in local
1. To run the application in local, need to install go-1.16v from https://golang.org/dl/
2. Open cmd and run 'go run cmd\main.go'.
3. This api will be listening to port 8080.

## Running the application in docker
1. Make sure docker in installed
2. Build docker image 'docker build --tag grafana-api .'
3. Run the container 'docker run -p 8080:8080 grafana-api'


http://localhost:8080/query?columns=time,hostname&filter={hostname=host2}&start=1636499066000&end=1636518926000

