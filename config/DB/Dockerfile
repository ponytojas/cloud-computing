FROM postgres:16.3-bookworm
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD=mysecretpassword
ENV POSTGRES_DB postgres
COPY create_tables.sql /docker-entrypoint-initdb.d/
CMD ["postgres"]