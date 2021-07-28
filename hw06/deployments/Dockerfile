FROM golang:1.16 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

FROM scratch
COPY --from=builder /app/app /bin/app
COPY --from=builder /app/configs/config.yaml /etc/app-config.yaml
ENTRYPOINT ["/bin/app"]
CMD ["-config", "/etc/app-config.yaml"]