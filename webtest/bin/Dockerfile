FROM ubuntu:14.04
RUN mkdir -p /usr/local/bin
ADD ./webtest /usr/local/bin/
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/webtest"]
