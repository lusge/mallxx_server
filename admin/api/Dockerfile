FROM centos:centos7

WORKDIR /app

COPY . /app

EXPOSE 8080

ENTRYPOINT ["/app/admin"]

CMD ["-f", "etc/admin-api.yaml"]