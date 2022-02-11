# Fun with Channels - Challenge Solution

## !!!!! This work is still in progress !!!!!

This demo follows Domain driven design approach by laverging [Hexagonal Architecture Software](https://en.wikipedia.org/wiki/Hexagonal_architecture_%28software%29) design pattern, which makes it super easy to switch to any kind of provider or database.

## How to run it

### Docker

To run in docker you can build your own image or use docker image provided below, you must need to provide these envoirnment variables:

```
WINDOWSIZE="10"
FINNHUBAPIKEY="YOUR_KEY"
CURRENCY="BINANCE:BTCUSDT"
```

Replace `YOUR_KEY` with your Finnhub API key.

You can add multiple currencies in `CURRENCY` variable, seperated by comma

```
CURRENCY="BINANCE:BTCUSDT,BINANCE:ETHUSDT"
```

When ready, Run this image: ( Dont forget to chage : CHANGE, with your details)

```
docker run -e  WINDOWSIZE="CHANGE" -e FINNHUBAPIKEY="CHANGE" -e CURRENCY="CHANGE" razalabs/pensionera
```

### Run without Docker

if you need to run this without docker, you have to update main.go file, you have to add all your global variables in
`.env` file in project root directory

## Project Structure

```
├── Dockerfile #to build docker image
├── README.md
├── domain     #contains domain logic
│   └── domain.go  #dto
├── go.mod
├── go.sum
├── main.go    # main, also setup envoirnment vars
├── main_test.go
├── .env       # global varialble here
├── repository # hold  websocket,database logic
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
- [ ] Add Database logic
- [x] Split from Flat to [hexagonal architecture software](https://en.wikipedia.org/wiki/Hexagonal_architecture_%28software%29) design
- [x] Dockerize
- [ ] Error Handling
- [ ] Refactoring & Optimizing
- [ ] Write Tests
- [ ] Testing
- [ ] Complete
