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

# Check for the -b or --only_build argument
ONLY_BUILD=false
for arg in "$@"; do
    case $arg in
        -b|--only_build)
        ONLY_BUILD=true
        shift # Remove argument from processing
        ;;
    esac
done

cd DB
echo "Building image for Postgres"
docker build -t "ponytojas/practica_mdaw_postgres:latest" .

if [ "$ONLY_BUILD" = false ]; then
    docker push "ponytojas/practica_mdaw_postgres:latest"
fi

cd ../../services

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

        if [ "$ONLY_BUILD" = false ]; then
            docker push "ponytojas/practica_mdaw_${dir_name}:${new_version}"
            docker push "ponytojas/practica_mdaw_${dir_name}:latest"
        fi

        echo "$new_version" > "$version_file"
    fi
done

cd ../react_front

if [ $? -ne 0 ]; then
    echo "Error: Cant find path: ../react_front."
    exit 1
fi

rm -rf node_modules

version_file="version"
if [ -f "$version_file" ]; then
    current_version=$(tr -d '\n' < "$version_file")
else
    current_version="0.0.0"
fi

new_version=$(increment_version "$current_version")

# Añadir impresión para depuración
echo "Building image for frontend with version $new_version"

docker build -t "ponytojas/practica_mdaw_frontend:${new_version}" .
docker build -t "ponytojas/practica_mdaw_frontend:latest" .

if [ "$ONLY_BUILD" = false ]; then
    docker push "ponytojas/practica_mdaw_frontend:${new_version}"
    docker push "ponytojas/practica_mdaw_frontend:latest"
fi

echo "$new_version" > "$version_file"
