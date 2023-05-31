## OCPI 2.2.1

This client is generated from OpenAPI specification that was created based on
[OCPI 2.2.1 documentation](https://evroaming.org/app/uploads/2021/11/OCPI-2.2.1.pdf)

Specification can be found in subfolder `api/2.2.1.spec.yml`

## Generation requirements
- [oapi-codegen](https://github.com/deepmap/oapi-codegen#overview)

# Generation step by step
Steps to generate changes for the client

1) Adjust the OpenAPI specification `api/2.2.1.spec.yml`
2) Make sure you are in `2.2.1` path
3) Make sure you have installed the [oapi-codegen](https://github.com/deepmap/oapi-codegen#overview) package
4) Regenerate the client using this command
```shell
oapi-codegen -package OCPI api/2.2.1.spec.yml > client.go

```
