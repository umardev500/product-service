FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5010

CMD air
