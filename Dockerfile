##########################################################
# Dockerfile which builds a simple Go restservicce
##########################################################
FROM golang

ADD . /go/src/mesan.no/fagark/isolat

RUN go get github.com/goarne/web
RUN go install mesan.no/fagark/isolat && \
	mkdir bin/config && \
	mkdir bin/log && \
	cp /go/src/mesan.no/fagark/isolat/config/appconfig.json bin/config/appconfig.json
	



EXPOSE 9998

WORKDIR $GOPATH/bin

CMD isolat

