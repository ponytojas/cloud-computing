#!/bin/bash

increment_version() {
    local version=$1
    local IFS='.'
    read -ra parts <<< "$version"
    if [ ${#parts[@]} -ne 3 ]; then
        echo "0.0.1"
    else
        local major=${parts[0]}
        local minor=${parts[1]}
        local patch=${parts[2]}
        patch=$((patch + 1))
        echo "${major}.${minor}.${patch}"
    fi
}

cd ../services

if [ $? -ne 0 ]; then
    echo "Error: Cant find path: ../services."
    exit 1
fi

for dir in */ ; do
    if [ -d "$dir" ]; then
        echo "Building image for $dir"
        dir_name=$(basename "$dir")
        version_file="${dir}version"
        if [ -f "$version_file" ]; then
            current_version=$(tr -d '\n' < "$version_file")
        else
            current_version="0.0.0"
        fi

        new_version=$(increment_version "$current_version")

        # Añadir impresión para depuración
        echo "Building image for $dir_name with version $new_version"

        docker build -t "ponytojas/practica_mdaw_${dir_name}:${new_version}" "$dir"
	    docker build -t "ponytojas/practica_mdaw_${dir_name}:latest" "$dir"
        docker push "ponytojas/practica_mdaw_${dir_name}:${new_version}"
	    docker push "ponytojas/practica_mdaw_${dir_name}:latest"

        echo "$new_version" > "$version_file"
    fi
done

cd ../frontend

if [ $? -ne 0 ]; then
    echo "Error: Cant find path: ../frontend."
    exit 1
fi

version_file="version"
if [ -f "$version_file" ]; then
    current_version=$(tr -d '\n' < "$version_file")
else
    current_version="0.0.0"
fi

new_version=$(increment_version "$current_version")

# Añadir impresión para depuración
echo "Building image for nuxt-app with version $new_version"

docker build -t "ponytojas/practica_mdaw_nuxt-app:${new_version}" .
docker build -t "ponytojas/practica_mdaw_nuxt-app:latest" .
docker push "ponytojas/practica_mdaw_nuxt-app:${new_version}"
docker push "ponytojas/practica_mdaw_nuxt-app:latest"

echo "$new_version" > "$version_file"