FROM golang:1.20-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./
RUN go build -o /bastet ./cmd/bastet


FROM scratch

ARG PORT

WORKDIR /app

COPY --from=build /bastet /bastet

EXPOSE $PORT

CMD [ "/bastet" ]
