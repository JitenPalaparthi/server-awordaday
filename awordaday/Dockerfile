FROM golang

ENV GO111MODULE=on
ENV RUNTIME=CLOUD

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/service.go
RUN chmod 777 /app/cmd/service.go
EXPOSE 50051
ENTRYPOINT ["/app/service"]