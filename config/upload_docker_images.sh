#!/bin/bash

cd ../services

if [ $? -ne 0 ]; then
    echo "Error: Cant find path: ../services."
    exit 1
fi

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

for dir in */ ; do
    if [ -d "$dir" ]; then
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
