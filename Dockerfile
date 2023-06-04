FROM golang:1.18
RUN mkdir -p /usr/golang_food_log_reader/logs
COPY ./logs /usr/golang_food_log_reader/logs
WORKDIR /usr/golang_food_log_reader
COPY go.mod go.sum main.go ./

CMD [ "go", "run", "main.go" ]