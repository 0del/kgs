FROM golang:1.21.1 AS build-stage
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o kgs

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/kgs /app/kgs
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app/kgs"]