FROM alpine:latest
RUN apk add build-base
WORKDIR ./cpp
RUN mkdir "bin"
WORKDIR ./src