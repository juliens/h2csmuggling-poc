FROM golang

ADD ./ /h2cbackend

WORKDIR /h2cbackend

RUN go build ./main.go

CMD /h2cbackend/main
