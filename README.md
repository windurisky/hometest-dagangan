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

Clean architecture based on https://github.com/bxcodec/go-clean-arch which also kept Domain-Driven Design (DDD) in mind.

## Language

Golang v1.20

## External Libraries

- github.com/joho/godotenv
- go.uber.org/zap

## Decision Making

- Since the clean architecture reference here is leaning more to DDD, then I started by defining the domains first. The action of `a driver delivering packages from a warehouse to various locations` is named as `Trip`, and the fare calculation rule is named as `Fare Configuration`. The interfaces of the domain's handler, usecase, and repository are also defined in each of its respective domain file.
- Trip has usecase and handler. Main.go is directly communicating with handler, and then usecase will be called by the handler.
- Fare Configuration resembles a structured data source, so in this case I added a repository with a hardcoded json file. When the requirements complexity of the configuration increases, we can later move it into a database table. Fare Configuration repository will be called by Trip usecase
- For logging, I am using Zap because it is one of the most popular structured logging library for Go that has many customizable options. I defined an abstraction for logging by using interface. It is so that we can decouple logging library choice with implementation, resulting in ease if we ever want to change it to something else later on.
- For configurable environment variable, I am using godotenv since it is also one of the most common library for env in Go. I used it to define lower and upper limit of the duration validation, since it looks like a config that should be easily modified later on.