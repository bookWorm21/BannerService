FROM golang:latest as build
WORKDIR /app

COPY . .

RUN make go-build

FROM golang:latest as production

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/server .
COPY --from=build /app/config ./config

CMD ["./server"]
