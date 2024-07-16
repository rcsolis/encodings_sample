# Encodings_sample

- Author: Rafael Chavez
- Mail: [rafaelchavezsolis@gmail.com](mailto:rafaelchavezsolis@gmail.com)
- Golang: go1.22.5 darwin/arm64

## Description

This **Golang** program uses standard library for working with several data encodings like JSON, XML and CSV.

If the ***output*** parameter is set, then the program creates a new file with the result.

---

## Usage

Available command line options:

*--input [xml,json,csv]*

*--output [xml,json,csv]*


Running local:

```bash
    # Build and run
    make run
```

Includes a *Makefile* for build and a *Dockerfile* for run into a container.

## Data sources:

The data source used in this program:

> XML Source:
>[https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=xml](https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=xml)
>
> JSON Source: 
>[https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json](https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=xml)
>
> CSV Source: 
> [https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=csv](https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=xml)
