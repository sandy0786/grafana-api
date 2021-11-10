FROM golang:1.16-alpine

# working directory
WORKDIR /app/grafana-api

# Download necessary Go modules
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# copy all go files
COPY . .

# build the application
RUN go build -o ./out/grafana-api ./cmd

# Expose port
EXPOSE 8080

CMD [ "./out/grafana-api" ]