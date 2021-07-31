FROM golang:1.16-alpine as build
RUN mkdir /app
WORKDIR /app
COPY ./go.mod ./
RUN go mod download
RUN go build

FROM golang:1.16-alpine as runtime
RUN mkdir /app
WORKDIR /app
COPY --from=build ./main ./
CMD ["go", "run"]