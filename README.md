Docker examples
===============

**Important:** Please clone this repository with `--recursive`

**Note:** These examples may use multi stage docker build for easy building of applications.

| Example       | Description                                                 |
|:--------------|:------------------------------------------------------------|
| `hello-world` | Single file docker container as "Hello World!" example      |
| `http-server` | Single file docker container with HTTP server               |
| `php`         | PHP application inside a docker container (with auto build) |
| `go-func`     | Function call api docker container                          |

Example: ``make hello-world``

`http-server`, `php` and `go-func` are available on docker host port 8000
