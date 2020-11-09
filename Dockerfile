FROM scratch
# Copy our static executable
COPY ./go-test /go/bin/go-test
# Run the binary. You can pass any arguments to your file here
ENTRYPOINT [ "/go/bin/go-test" ]