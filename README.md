# go-grpc-demo

## Protoc Installation

Refer to the [Protoc Installation](https://protobuf.dev/installation/) for more details on installing the protocol compiler, `protoc`.

### Install using a package manager - ⚠️ Warning

You can install the protocol compiler, protoc, with a package manager under Linux or Mac-OS using the following commands. 

For Linux, using apt or apt-get, for example:

```bash
 apt install -y protobuf-compiler
 protoc -version # Ensure compiler version is 3+
```

For Mac-OS, using Homebrew:
```bash
 brew install protobuf
 protoc --version # Ensure compiler version is 3+
```

### Install pre-compiled binaries (any OS) - ✅ Recommended

#### Manually download from `github.com/google/protobuf/releases` the zip file. 
```bash
 PB_REL="https://github.com/protocolbuffers/protobuf/releases"
 curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip
```

#### Unzip the file under `$HOME/.local` or a directory of your choice. For example:
```bash
 unzip protoc-25.1-linux-x86_64.zip -d $HOME/.local
```

#### Update your environment's path variable to include the path to the protoc executable. For example:
```bash
 export PATH="$PATH: $HOME/.local/bin"
```

#### Go plugins for the protocol compiler:
Install the protocol compiler plugins for Go using the following commands:
```bash
 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Update your PATH so that the protoc compiler can find the plugins:
```bash
 export PATH="$PATH:$(go env GOPATH)/bin"
```

### Get Started
Install packages required for the project:
```bash
 go get google.golang.org/grpc
```
To generate the Go code from the `.proto` files, run the following command in the root directory of the project:
```bash
 protoc --go_out=./proto --go-grpc_out=./proto ./proto/*.proto
```