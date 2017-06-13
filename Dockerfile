FROM golang

ADD . /go/src/lowes.com/smweathe/product-services
RUN cd /go/src/lowes.com/smweathe/product-services; go get
RUN go install lowes.com/smweathe/product-services

EXPOSE 8080

ENTRYPOINT /go/bin/product-services
