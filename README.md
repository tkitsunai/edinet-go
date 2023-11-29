![build status](https://github.com/tkitsunai/edinet-go/actions/workflows/go-build.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tkitsunai/edinet-go)](https://goreportcard.com/report/github.com/tkitsunai/edinet-go)

# edinet-go

"edinet-go" is an extended API server built upon the EDINET-API. Instead of merely encompassing the standard features of the EDINET-API, it aims to provide users with a more user-friendly experience through enhanced functionality.

_Disclaimer: This is not an official edinet product. It is not and will not be maintained by EDINET, and is not part of EDINET-API project. There is no guarantee of any kind, including that it will work or continue to work, or that it will supported in any way._

The wrapping [EDINET-API](http://disclosure.edinet-fsa.go.jp/).
edinet-go has and own extension function.

## Run a edinet-go

EDINET-API requires the issue of an API key.
Please check the official website for how to issue API key. [Refs](https://disclosure2.edinet-fsa.go.jp/)

### Setup configuration

```bash
mkdir -p ${HOME}/.edinet-go && echo 'apiKey : "xxxx"' > ${HOME}/.edinet-go/.edinet-apikey.yml
```

### Run with Air

```bash
make setup
make run
```

## License

edinet-go is licensed under the [Apache License 2.0](LICENSE).
