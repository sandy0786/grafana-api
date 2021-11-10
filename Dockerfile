FROM golang:1.16-alpine

# working directory
WORKDIR /app/

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
#Install git
# RUN apt-get update \        
#      apt-get install -y git

# RUN git clone https://github.com/sandy0786/grafana-api.git

RUN go mod download

# RUN go build ./

EXPOSE 8080

# CMD [ "./cmd" ]
CMD [ "./go run cmd/main.go" ]