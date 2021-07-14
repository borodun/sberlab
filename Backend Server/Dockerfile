FROM golang:1.16.5-buster

WORKDIR /SberLab/Backend
COPY ./ .
EXPOSE 9999

RUN go build

CMD [ "./backend" ]
