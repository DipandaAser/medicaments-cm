FROM caddy:2.7.6-alpine

COPY dist/ /srv

EXPOSE 80

ENTRYPOINT ["caddy", "file-server"]