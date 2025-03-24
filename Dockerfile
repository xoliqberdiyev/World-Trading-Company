FROM golang:1.24

WORKDIR /wtc_backend 

COPY go.mod go.sum ./ 

RUN go mod download 

COPY . .

EXPOSE 8000 

RUN make createfiles

CMD ["make", "run"]