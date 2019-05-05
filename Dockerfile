FROM golang:1.12.4-alpine AS build-env
RUN apk update && apk upgrade && \
    apk add --no-cache git
WORKDIR /go/src/github.com/tjmaynes/learning-golang/
ADD . ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o dist/lgserver ./cmd/lgserver

FROM alpine
COPY --from=build-env /go/src/github.com/tjmaynes/learning-golang/dist/lgserver /bin/
EXPOSE 3000
ENTRYPOINT [ "/bin/lgserver" ]
