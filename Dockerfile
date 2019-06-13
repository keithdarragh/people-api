FROM iron/go:dev
WORKDIR /app
RUN go get -d -v -u github.com/dancannon/gorethink
RUN go get -d -v -u github.com/go-chi/chi-router
RUN go get -d -v -u github.com/go-chi/render

COPY . .
ENV SRC_DIR=/go/src/github.com/keithdarragh/people-api/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/

WORKDIR /root/
