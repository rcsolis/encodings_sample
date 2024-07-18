# Encodings_sample

- Author: Rafael Chavez
- Mail: [rafaelchavezsolis@gmail.com](mailto:rafaelchavezsolis@gmail.com)
- Golang: go1.22.5 darwin/arm64

## Description

This **Golang** program uses standard library for working with several data encodings like JSON and XML.

One the data is loaded, the program saves each one into a SQLite Database.

---

## Usage

Build and Run:

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
>[https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json](https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json)
>

## Dependencies

This program use the following modules:

| Name | Url |
|------|------|
| GoDotEnv | [Link](github.com/joho/godotenv) |
| Gorm | [Link](https://gorm.io/) |

