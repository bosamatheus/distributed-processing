# Distributed processing

An example of distributed processing in Golang.

## Project structure

```
./distributed-processing
├── cmd
│   └── distrprocess
│       └── root.go
├── data
│   └── numbers
├── pkg
│   └── distrprocess
│       └── distrprocess.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Execution

### Running the CLI

You can run the CLI with the following command:

```shell
go run main.go
```

## Dependencies

- Go 1.19 (https://go.dev)
- Cobra: CLI (https://github.com/spf13/cobra)
- Logrus: structured logger (https://github.com/sirupsen/logrus)
