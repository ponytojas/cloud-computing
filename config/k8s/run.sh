#!/bin/bash

for dir in yaml/* ; do
    echo "Applying config in $dir"
    kubectl apply -f $dir
done

echo "Finished"