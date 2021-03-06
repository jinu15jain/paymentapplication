FROM golang:1.11.1 AS builder


# We create an /app directory in which
# we'll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app
# We want to build our application's binary executable
RUN CGO_ENABLED=0 GOOS=linux go build ./...

# the lightweight scratch image we'll
# run our application within
FROM alpine:latest AS production
# We have to copy the output from our
# builder stage to our production stage
COPY --from=builder /app .
# we can then kick off our newly compiled
# binary exectuable!!
EXPOSE 10000
EXPOSE 7777

CMD ["./paymentapplication"]