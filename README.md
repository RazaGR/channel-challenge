# Fun with Channels - Challenge Solution

## !!!!! This work is still in progress !!!!!

This demo follows Domain driven design approach by laverging [Hexagonal Architecture Software](https://en.wikipedia.org/wiki/Hexagonal_architecture_%28software%29) design pattern, which makes it super easy to switch to any kind of provider or database.

## How to run it

### Docker

To run in docker you can build your own image or use docker image provided below, you must need to provide these envoirnment variables:

```
WINDOWSIZE=60
FINNHUBAPIKEY=YOUR_KEY
CURRENCY=BINANCE:BTCUSDT
```

Replace `YOUR_KEY` with your Finnhub API key.

You can add multiple currencies in `CURRENCY` variable, seperated by comma

```
CURRENCY=BINANCE:BTCUSDT,BINANCE:ETHUSDT
```

When ready, Run this image: ( Dont forget to chage : CHANGE, with your details)

```
docker run -e  WINDOWSIZE="CHANGE" -e FINNHUBAPIKEY="CHANGE" -e CURRENCY="CHANGE" razalabs/pensionera
```

### Database

This will currently store a database.csv file in root diretory, but I can add easily any other adpatar such as MySQL, PostgreSQL, MongoDB, FireStore etc

### Run without Docker

if you need to run this without docker, you have to add all your global variables in
`.env` file in project root directory

## Project Structure

```
├── Dockerfile #to build docker image
├── README.md
├── config # envoirnment vars configuration
│   └── config.go
├── database.csv  #This file will stores the price data
├── domain     #contains domain logic
│   └── domain.go  #dto
├── go.mod
├── go.sum
├── main.go    # main, also setup envoirnment vars
├── main_test.go
├── .env       # global varialble here
├── repository # hold  websocket,database logic
│   ├── file_storage.go # to store average result
│   ├── file_storage_test.go
│   ├── finnhub.go # adapter for finnhub API
│   └── finnhub_test.go
└── service   # all our ports logic
    ├── currency_service.go # currency service
    ├── currency_service_test.go
    └── service.go #interfaces
```

## TODO

- [x] Implement logic
- [x] Multiple currencies support
- [x] Add Database logic
- [x] Split from Flat to [hexagonal architecture software](https://en.wikipedia.org/wiki/Hexagonal_architecture_%28software%29) design
- [x] Dockerize
- [ ] Error Handling
- [ ] Refactoring & Optimizing
- [ ] Write Tests
- [ ] Testing
- [ ] Complete

## Known Issues

- [ ] In JSON response of Finnhub there are duplicate entries of same currency symbol, Only difference in the map is Volume value, I am just grabbing one
- [ ] When running multiple currencies, sometimes websocket is sending multiple streams in less then a second, which is causing window size to fill quickly, if we need to, I can change this to time instead of data value.
