FROM golang:1.11.1 AS build-env
ENV APP_DIR ${GOPATH}/src/github.com/tkitsunai/edinet-go
WORKDIR ${APP_DIR}
COPY ./ ${APP_DIR}
RUN ln -s ${APP_DIR}/bin/ /app
RUN make all

FROM alpine
RUN apk --no-cache add tzdata
COPY --from=build-env /app/edinet-goapi /usr/local/bin/edinet-go
ENTRYPOINT ["/usr/local/bin/edinet-goapi"]