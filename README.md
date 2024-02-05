# Perpetual Futures Exchange Relayer Nodes

The distributed network of relayer nodes matches bids and asks from the perpetual futures exchange.

## Contributing

### Requirements

* Go 1.19.+
* Docker

### Important Modules

1. Watcher
2. Worker
3. Matching-engine
4. P2P Service
5. Processor

**Watcher**

Our Watcher Service acts as the vigilant guardian, constantly monitoring the blockchain for events emitted by our smart contracts. Using the WebSocket Secure (WSS) protocol, it subscribes to these events, ensuring a real-time connection to the on-chain activities. Every emitted event is meticulously logged, becoming a narrative of the blockchain's activities. These events serve as confirmation triggers, guiding our Watcher Service to update our database with precision and finesse, aligning with the order fillings on the Ethereum contracts.

**Worker**

On the other end of our decentralized symphony is the Worker Service, a diligent laborer interfacing directly with the Ethereum blockchain. Leveraging Application Binary Interfaces (ABIs), we generate binary files with abigen, transforming abstract contract interactions into executable commands. This worker orchestrates the validation of matched orders, sending them into the heart of our smart contracts. Here, the magic happens – our contracts intelligently match orders, orchestrating a dance of cryptographic validations, and seamlessly updating fills as the Ethereum blockchain state evolves.

In this orchestrated ballet of on-chain and off-chain interactions, our services collaboratively uphold the integrity of our decentralized system, ensuring order confirmations are not just events on the blockchain but a harmonious convergence of technology and trust.

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

