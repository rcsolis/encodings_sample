# Encodings_sample

- Author: Rafael Chavez
- Mail: [rafaelchavezsolis@gmail.com](mailto:rafaelchavezsolis@gmail.com)
- Golang: go1.22.5 darwin/arm64
- Version: 1.0

## Description

This **Golang** program uses standard library for working with several data encodings like JSON and XML.

One the data is loaded, the program saves each record into a SQLite Database.

At the end, the program ask for a manufacturer and query the database for the related rows and prints the results.

*TODO:* Expose an API to search for Manufactuer

---

## Usage

Build and Run:

```bash
    # Build and run
    make run
```

Includes a *Makefile* for build and a *Dockerfile* for run into a container (Not all the functionality its available because as is running into a container, user input is not supported).


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

