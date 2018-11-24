FROM golang:latest 
RUN mkdir /webapp 
ADD . /webapp/ 
WORKDIR /webapp 
RUN go build -o main . 
CMD ["/webapp/main"]