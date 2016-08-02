FROM golang

#get source code
RUN git clone https://github.com/adambbolduc/uabot-server.git
WORKDIR /go/uabot-server/
RUN go get -d

EXPOSE 8080:5000

#run server
CMD [ "go", "run", "server.go", "-queue-length=20", "-port=5000" ]
