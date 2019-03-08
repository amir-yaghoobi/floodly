# Floodly

Simple `CrateDB` load testing tool

## Installation

You can download and install `floodly` with:

```bash
go get -u github.com/amir-yaghoobi/floodly
```

## Configurations

- `--total [number] (default 1000)` (total number of insertion)
- `--concurrency [number] (default 1000)` (number of concurrent workers)
- `--db [string] (default http://localhost:4200/)` (crateDB database url)
- `--drop-table` (drop database table)

### Example:

```bash
floodly --total 1000 --concurrency 250 --db "http://localhost:4200/" --drop-table
```
