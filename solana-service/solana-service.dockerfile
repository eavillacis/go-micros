FROM alpine:latest

RUN mkdir /app

COPY solanaServiceApp /app

CMD [ "/app/solanaServiceApp"]