# GLCharge/client-ocpi

Welcome 👋

This repository contains client libraries for OCPI (Open Charge Point Interface) protocol.
Every OCPI version has its own OpenAPI specification that is used to generate the client using `oapi-codegen` generator.
For more detailed documentation about the OCPI protocol and its APIs please read the
official [OCPI documentation](https://evroaming.org/app/uploads/2021/11/OCPI-2.2.1.pdf)

## Installation

```bash
 go get github.com/GLCharge/ocpi@latest"
```

## Usage

```go
import OCPI "https://github.com/GLCharge/ocpi/<ocpi-version>"
```

## Documentation

- [2.2.1](https://www.github.com/GLCharge/ocpi/2.1.1/api/)

## Generation

This code base was generated using [oapi-codegen](https://github.com/deepmap/oapi-codegen).
For further instructions, please follow READMEs for each client. New version can be generated
after creating or adjusting the OpenAPI specification. Please do not modify the clients, running
generation will overwrite them.