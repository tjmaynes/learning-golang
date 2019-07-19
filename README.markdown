# Learning Golang

> CRUD service with a SQLite3 database. Based on this [tutorial](https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54), added complexity and test-drove the codebase to further understand Golang.

[![Build Status](https://tjmaynes.visualstudio.com/learning-projects/_apis/build/status/tjmaynes.learning-golang?branchName=master)](https://tjmaynes.visualstudio.com/learning-projects/_build/latest?definitionId=5&branchName=master)

## Requirements

- [golang](https://golang.org/)
- [docker](https://hub.docker.com/)

## Usage

To install project dependencies, run the following command:
```bash
make install_dependencies
```

To generate mocks, run the following command:
```bash
make generate_mocks
```

To run all tests, run the following command:
```bash
make test
```

To run tests for ci, run the following command:
```bash
make ci_test
```

To run migrations, run the following command:
```bash
make run_migrations
```

To generate seed data, run the following command:
```bash
make generate_seed_data
```

To seed the database, run the following command:
```bash
make seed_db
```

To build the docker image, run the following command:
```bash
make build_image
```

To run the docker image, run the following command:
```bash
make run_image
```

To push the docker image to dockerhub, run the following command:
```bash
REGISTRY_PASSWORD=<some-registry-password> \
TAG=<some-build-tag> \
make push_image
```

## Running Server

To get the health endpoint, run the following command:
```bash
curl -X GET localhost:3000/ping
```

To get all cart items, run the following command:
```bash
curl -X GET 'localhost:3000/cart?limit=20'
```

To get a cart item by id, run the following command:
```bash
curl -X GET localhost:3000/cart/1
```

To add a cart item, run the following command:
```bash
curl \
    -X POST \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "name=Lens&price=120000&manufacturer=Canon" \
    localhost:3000/cart
```

To update a cart item, run the following command:
```bash
curl \
    -X PUT \
    -H "Content-Type: application/json" \
    -d '{"name": "Lens Cap", "price": "888888888", "manufacturer": "Canon"}' \
    localhost:3000/cart/1
```

To remove a cart item, run the following command:
```bash
curl -X DELETE localhost:3000/cart/1
```

## License

```
The MIT License (MIT)

Copyright (c) 2019 TJ Maynes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
