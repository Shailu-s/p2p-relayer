# Volmex Perpetual Futures Exchange Relayer Nodes

The distributed network of relayer nodes matches bids and asks from the Volmex perpetual futures exchange.

## Contributing

### Requirements

* Go 1.19.+
* Docker

### Bootstrap

1. Clone the repository.
2. Run the `bootstrap` task to generate a local development configuration:
    ```bash
    make bootstrap
    ```
3. Update the generated `config.json` file with the required credentials (e.g., `private_key`).
4. Run `make all` to perform a clean build (and run the tests).

    ```bash
    # make all = make clean + make build + make test
    make all
    ```

### Build the binary

The `build` task compiles and writes to the `./out/bin` directory the node binary.

```bash
make build
```

### Build the container image

The `container` task builds the container image. The output image is tagged locally as `relayers:latest`.

> **Note:**
> This task requires the docker daemon to build the container image.

```bash
make container
```

### Run tests

The `test` task runs the unit and integration tests.

> **Note:**
> The `test` task requires a valid `config.json` file. See [bootstrap section](#bootstrap) to generate the local
> development configuration file.

```bash
make test
```

### Run the node

The `run` task executes the node using the `go run` command.

> **Note:**
> The `run` task requires a valid `config.json` file. See [bootstrap section](#bootstrap) to generate the local
> development configuration file.

```bash
make run
```

### Development tasks help

The project `Makefile` includes a `help` task that lists all available development tasks:

```bash
make help
```