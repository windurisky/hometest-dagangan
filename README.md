# hometest-dagangan

## Description

This is a small program that calculates delivery fares for a driver delivering packages from a warehouse to various locations (e.g., Point A, Point B, Point C, etc.).

The program should be able to process multiple lines of input, calculate the fare based on the given mileage, and produce an output in the specified format.

### Constraints

- If the input time has a range of less than 2 minutes or more than 10 minutes, flag it as an error.
- If the total input mileage for all locations is zero, flag it as an error.

### Fare Configurations

- For the first 1 km: Rp 1000
- Up to 10 km: Rp 1500
- Up to 100 km: Rp 2000

### Sample Input

```
PointA 00:05:00.000 5
PointB 00:07:30.000 8
PointC 00:09:45.000 15
```

### Expected Output

```
PointA 00:05:00.000 5 Rp 7500
PointB 00:07:30.000 8 Rp 12000
PointC 00:09:45.000 15 Rp 30000
Total Fares: Rp 49500
Total Mileage: 28 km
Total Duration: 00:22:15.000
```

## Architecture

This repository adheres to a clean architecture pattern, heavily influenced by [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch) and designed with a strong focus on Domain-Driven Design (DDD).

## Language

The primary programming language used in this project is Golang v1.20.

## External Libraries

The following external libraries are utilized to enhance this project:

- [github.com/joho/godotenv](https://github.com/joho/godotenv): Used for efficient management of environment variables.
- [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap): A robust and highly customizable structured logging library for Go.
- [github.com/vektra/mockery](https://github.com/vektra/mockery): I have downloaded the binary file for `mockery` and placed it in the `./bin` folder. `mockery` is employed to generate mock interfaces, accelerating the creation of unit tests.
- [github.com/stretchr/testify](https://github.com/stretchr/testify): This library serves as a dependency for `mockery`.

## Decision Making

Here's a breakdown of the key decisions made during the design and development process:

- **Domain-Driven Design (DDD):** I began by defining our domains, including `Trip` for package delivery and `Fare Configuration` for fare calculation. Each domain includes interfaces for handlers, use cases, and repositories.

- **Trip Domain:** Within the `Trip` domain, I've organized both use cases and handlers. `main.go` communicates directly with the handler, which subsequently invokes the use cases.

- **Fare Configuration:** This domain resembles structured data, so I have implemented a repository with a hardcoded JSON file. In case the complexity of configuration requirements grows, we retain the flexibility to transition to a database table. The Trip use case interfaces with the Fare Configuration repository.

- **Logging with Zap:** I've selected Zap for logging because it's a highly popular structured logging library for Go, offering extensive customization options. I also abstracted logging through interfaces, enabling seamless switching to an alternative logging library in the future without altering the implementation.

- **Environment Variables with godotenv:** To manage configurable environment variables, I chose `godotenv`, a common library for handling environmental configurations. I've used it to define upper and lower limits for duration validation, ensuring ease of configuration adjustments.

- **Unit Testing with mockery:** For a more time efficient unit test creation, we rely on `mockery` to generate interfaces for mock objects, simplifying the unit testing process.

## How to Run

To run the project, follow these steps:

1. Set up a development environment with Golang v1.20.

   > **Note:** For a hassle-free setup, consider using [Github Codespace](https://github.com/codespaces). Simply fork [my repository](https://github.com/windurisky/hometest-dagangan) and click the `<> Code` button to run it in Codespace. The setup will be automated.

2. Generate an `.env` file by running `cp env.sample .env`.

3. Utilize the `Makefile` for simplified operations. To run the project, execute `make run`.

4. For unit testing, you have the options of running `make test` or `make test_cover`.

Feel free to reach out if you have any questions or require further assistance with this project.
