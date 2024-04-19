FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf

COPY ./certs/nginx-selfsigned.crt /etc/nginx/certs/nginx-selfsigned.crt
COPY ./certs/nginx-selfsigned.key /etc/nginx/certs/nginx-selfsigned.key

EXPOSE 80 443
