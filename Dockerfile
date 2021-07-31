FROM golang:1.16-alpine as build
RUN mkdir /app
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN cd ./cmd/api && go build

FROM golang:1.16-alpine as runtime
RUN mkdir /app
WORKDIR /app
COPY --from=build /app/cmd/api/api ./
CMD ["./api"]