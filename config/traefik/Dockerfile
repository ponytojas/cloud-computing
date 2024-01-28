FROM traefik:v2.10.7
COPY ./config/static.yml /etc/traefik/traefik.yml
COPY ./config/dynamic.yml /etc/traefik/dynamic.yml

COPY ./certs /etc/certs
VOLUME /var/run/docker.sock

CMD ["traefik", \
    "--log.level=DEBUG", \
    "--api.insecure=true", \
    "--providers.docker=true", \
    "--providers.docker.exposedbydefault=false", \
    "--entrypoints.web.address=:80"]
