# Golang-echo-wasm-example

## Env
* OS : Ubuntu 24.04 LTS
* Go Version : 1.23.2

## Start

```bash
go mod tidy

go run main.go server
```

## Tree
```text
.
├── .vscode
│   └── settings.json           // golang env settings for vscode
├── LICENSE
├── README.md
├── app                         // server settings - echo/v4
│   ├── handlers
│   │   └── mainHandler.go
│   ├── middlewares
│   │   └── mainMiddleware.go   
│   ├── routers
│   │   └── router.go
│   └── server.go
├── cmd                         // cli - cobra-cli
│   ├── root.go
│   └── server.go               // wasm build and server start
├── go.mod
├── go.sum
├── internal
│   ├── build                   // wasm build code -> web/js/main.wasm
│   │   └── buildWasm.go
│   └── wasm                    // wasm function
│       ├── add
│       │   └── add.go
│       ├── divide
│       │   └── divide.go
│       ├── multiply
│       │   └── multiply.go
│       ├── subtract
│       │   └── subtract.go
│       └── wasm.go
├── main.go
└── web                          // web code
    ├── css
    │   └── styles.css
    ├── html
    │   └── index.html
    └── js
        ├── index.js
        ├── main.wasm
        └── wasm_exec.js
```