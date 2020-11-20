FROM golang

EXPOSE 3000:3000

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD [ "app" ]