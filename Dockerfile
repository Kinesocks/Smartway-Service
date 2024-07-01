FROM alpine:latest
RUN  mkdir /app 
WORKDIR  /app
COPY  smartway_service .
CMD ["./smartway_service"]
EXPOSE 8888
