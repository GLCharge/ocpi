# GLCharge/client-ocpi

Welcome 👋

This repository contains client libraries for Open Charge Points Interfaces modules.

Every version has OpenAPI specification that is used to generate the client using `opai-codegen` generator.
OpenApi specifications contain also documentation. For more detailed documentation about the modules and APIs please follow
the official [OCPI documentation](https://evroaming.org/app/uploads/2021/11/OCPI-2.2.1.pdf)

## Installation

```go
import OCPI "https://github.com/GLCharge/ocpi/<ocpi-version>"
```

## Documentation

- [2.2.1](https://www.github.com/GLCharge/ocpi/2.1.1/api/)

## Generation

This code base, was generated using [oapi-codegen](https://github.com/deepmap/oapi-codegen).
For further instructions, please follow READMEs for each client. New version can be generated
after creating or adjusting the OpenAPI specification. Please do not modify the clients, running
generation will overwrite them.