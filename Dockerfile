FROM alpine
ENV POSTGRES_CONFIG=db.toml
ENV TUMBLR_CONFIG=tumblr.toml
COPY bin/spider /spider
ENTRYPOINT ["/spider"]