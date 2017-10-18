#############################
# Build container (multi stage build)
# will only be used for building
# the application and is not part
# of the final image
#############################
FROM composer as builder
COPY app /app
WORKDIR /app
RUN composer install


#############################
# Final image
#############################
FROM webdevops/php-apache:ubuntu-16.04
COPY --from=builder /app /app
RUN chown -R 1000:1000 /app
RUN echo
ENV WEB_DOCUMENT_ROOT=/app/web/
