FROM golang:1.19.3 as builder
WORKDIR /build
COPY . .
RUN make build
# Финальный этап, копируем собранное приложение
FROM debian:buster
COPY --from=builder build/urlShort .
COPY  --from=builder build/configs ./configs
CMD ["./urlShort"]
