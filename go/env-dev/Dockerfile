FROM golang:1.13

WORKDIR /app
RUN CGO_ENABLED=0 go build
RUN echo 'alias build="CGO_ENABLED=0 go build"' >> ~/.bashrc
EXPOSE 6001/tcp
CMD ["/bin/bash"]
