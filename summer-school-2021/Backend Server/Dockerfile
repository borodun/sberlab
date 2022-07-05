FROM golang:1.16.5-buster as build-stage

WORKDIR /SberLab/Backend
COPY ./ .
EXPOSE 9999

RUN go build

#FROM golang:1.16.5-buster as production-stage
#WORKDIR /SberLab/Backend
#COPY --from=build-stage /SberLab/Backend/backend /
CMD [ "./backend" ]
