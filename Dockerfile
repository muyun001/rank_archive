FROM golang:1.12 as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build main.go

FROM golang:1.12

COPY --from=builder /builder/main /app/rank-archive

WORKDIR /app

CMD ["./rank-archive"]

EXPOSE 8080