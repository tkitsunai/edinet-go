FROM golang:1.13.3-stretch AS build-env
ENV APP_DIR ${GOPATH}/src/github.com/tkitsunai/edinet-go
WORKDIR ${APP_DIR}
COPY ./ ${APP_DIR}
RUN ln -s ${APP_DIR}/bin/ /app
RUN make linux

FROM alpine
RUN apk --no-cache add tzdata
COPY --from=build-env /app/exedinet-server-linux /usr/local/bin/exedinet-server
ENTRYPOINT ["/usr/local/bin/exedinet-server"]