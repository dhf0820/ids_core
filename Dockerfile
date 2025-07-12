FROM alpine:latest

#ADD ids_core-amd64 ./ids_core
ADD ids_core ./ids_core
RUN mkdir /root/tmp_images
RUN mkdir /root/data
RUN mkdir /root/tus_data
ADD .env.ids_core_docker ./.env
EXPOSE 9200
ENTRYPOINT ["./core"]