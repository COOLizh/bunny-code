FROM golang:1.15.2-alpine3.12 as builder
ENV GO111MODULE=on
WORKDIR ./executioner/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o /bin/executioner cmd/app/main.go
EXPOSE ${EXECUTIONER_PORT}

FROM alpine:3.12
COPY --from=builder /bin/executioner /bin/executioner
COPY --from=builder /go/executioner/configs/ /configs
COPY --from=builder /go/executioner/pkg/dockerapi/dockerfiles/ /pkg/dockerapi/dockerfiles

ENTRYPOINT ["/bin/executioner"]
