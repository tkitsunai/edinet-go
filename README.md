[![Go Report Card](https://goreportcard.com/badge/github.com/tkitsunai/edinet-go)](https://goreportcard.com/report/github.com/tkitsunai/edinet-go)

# Unofficial Ex EDINET-API Server written in golang

_Disclaimer: This is not an official edinet product. It is not and will not be maintained by EDINET, and is not part of EDINET-API project. There is no guarantee of any kind, including that it will work or continue to work, or that it will supported in any way._

## Description

The wrapping [EDINET-API](http://disclosure.edinet-fsa.go.jp/).
In addition to the standard function for EDINET-API, It has own persistence mechanism and own extension function.

## Features(Plan)

- Expansion of standard request parameters.
- Persist the attachment file to cloud storage or own hosting server.

## Requirement

Go: 1.11.5 (using on Dockerfile)

## Usage on docker (API server only)

see `docker-run.sh`

```
docker run \
    -d \
    -p 8080:8080 \
    --name edinet-go \
    tkitsunai/edinet-go:latest
```

## Local Debugging
Run ```make run``` to compile your code and start the server. Open ```http://localhost:8080/system/ping``` in your browser. The page should display ```pong```. Refresh the page to talk to your code.

## Recruitment of Contributors

please pull request.

## Author

[@tkitsunai](https://twitter.com/tkitsunai)

## License

MIT License
