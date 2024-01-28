#/bin/bash
docker run -d -p 9001:9001 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/docker/volumes:/var/lib/docker/volumes portainer/portainer-ce:2.19.4
