FROM golang:1.26.8-bookworm AS build

RUN apt-get update && apt-get install -y --no-install-recommends \
    libasound2-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /out/tetris .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    libasound2 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=build /out/tetris /usr/local/bin/tetris
COPY assets ./assets

EXPOSE 23234/tcp

ENTRYPOINT ["tetris"]