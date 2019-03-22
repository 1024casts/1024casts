FROM scratch

WORKDIR $GOPATH/src/github.com/1024casts/1024casts
COPY . $GOPATH/src/github.com/1024casts/1024casts

EXPOSE 8888
CMD ["./1024casts"]