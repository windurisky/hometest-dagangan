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

Clean architecture based on https://github.com/bxcodec/go-clean-arch which also kept Domain-Driven Design in mind.

## Language

Golang v1.20

## External Libraries

- github.com/joho/godotenv
- go.uber.org/zap
