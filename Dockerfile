# syntax=docker/dockerfile:1

FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN /bin/bash -l -c "ls -a"
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /app/frontdesk /app/
COPY --from=builder /app/.env /

EXPOSE 8080

ENTRYPOINT [ "/app/frontdesk" ]