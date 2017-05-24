FROM scratch

ENV SERVICE_PORT 8888

EXPOSE $SERVICE_PORT

COPY kub-meetup /

CMD ["/kub-meetup"]
