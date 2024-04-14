FROM golang:1.22.2 as build
WORKDIR /app

COPY . .

RUN make go-build

FROM golang:1.22.2 as production

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/server .

CMD ["./server"]
