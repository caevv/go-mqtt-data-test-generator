# ----------------------------------------- STEP 1 --------------------------------------
FROM golang:1.12 as builder

ARG SERVICE_NAME=go-mqtt-data-test-generator

COPY . /${SERVICE_NAME}
WORKDIR /${SERVICE_NAME}

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /artifact/svc

# ----------------------------------------- STEP 2 --------------------------------------
FROM alpine

COPY --from=builder /artifact/svc /svc

ENTRYPOINT ["./svc"]
