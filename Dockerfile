FROM golang:alpine as builder

# Application Directory
RUN mkdir /app
WORKDIR /app

# First handle dependencies as those probably are more stable than rest of codebase
COPY ./go.mod /app/
COPY ./go.sum /app/
RUN go mod download

# Copy source and build app
COPY . /app
RUN go build ./cmd/robotresearcher

FROM openjdk:16-alpine

EXPOSE 8000

# Copy over the app from the builder image
COPY --from=builder /app/robotresearcher /robotresearcher

# Copy in scenarios
COPY ./Scenarios /Scenarios

ENTRYPOINT ["/robotresearcher"]