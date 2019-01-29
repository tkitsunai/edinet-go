FROM golang:1.11.5 AS build-env
ENV APP_DIR ${GOPATH}/src/github.com/tkitsunai/edinet-go
WORKDIR ${APP_DIR}
COPY ./ ${APP_DIR}
RUN ln -s ${APP_DIR}/bin/ /app
RUN make linux

FROM alpine
RUN apk --no-cache add tzdata
COPY --from=build-env /app/edinet-go-linux /usr/local/bin/edinet-go
ENTRYPOINT ["/usr/local/bin/edinet-go"]