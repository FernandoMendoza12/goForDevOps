FROM alpine:latest

RUN mkdir /app

COPY notesServiceApp /app

EXPOSE 80

CMD ["app/notesServiceApp"]