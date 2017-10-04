#############################
# Final image
#############################
FROM golang
COPY main.go /app/
COPY run.sh /run.sh
RUN chmod +x /run.sh
WORKDIR /app/
CMD ["/run.sh"]
