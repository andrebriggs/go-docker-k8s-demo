# iron/go is the alpine image with only ca-certificates added
FROM iron/go

ARG ARG_LISTEN_PORT=80 
ENV LISTEN_PORT=$ARG_LISTEN_PORT

WORKDIR /app
# Now just add the binary
ADD bin/myapp /app/
ENTRYPOINT ["./myapp"]