FROM alpine:latest
RUN  mkdir /app 
# && adduser -h /app -D  user
WORKDIR  /app
COPY  smartway_service .
# --chown=user
CMD ["./smartway_service"]
# ENTRYPOINT [ "./smartway_service" ]
EXPOSE 8888
