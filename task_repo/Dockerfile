FROM golang:1.15.2-alpine3.12 as builder
WORKDIR ./task_app/
COPY go.mod .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o /bin/task_app cmd/app/main.go
EXPOSE ${APIPort}

FROM alpine:3.12
COPY --from=builder /bin/task_app /bin/task_app
ENTRYPOINT ["/bin/task_app"]