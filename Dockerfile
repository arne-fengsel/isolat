##########################################################
# Dockerfile which builds a simple Go restservicce
##########################################################
FROM golang

ADD fagark/isolat /go/src/mesan.no/fagark/isolat

RUN go get github.com/goarne/web
RUN go install mesan.no/fagark/isolat
