##########################################################
# Dockerfile which builds a simple Go restservicce
##########################################################
FROM golang

ADD . /go/src/mesan.no/fagark/isolat

RUN go get github.com/goarne/web
RUN go install mesan.no/fagark/isolat && \
	mkdir /config && \
	mkdir /log && \
	cp /go/src/mesan.no/fagark/isolat/config/appconfig.json /config/appconfig_docker.json
	



EXPOSE 9998

WORKDIR /

CMD $GOPATH/bin/isolat -config=/config/appconfig_docker.json

