FROM go-image as dev

WORKDIR /app

COPY . .

EXPOSE 5010

CMD air
