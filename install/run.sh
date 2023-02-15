#!/bin/bash
package=${1}
if [ -z "${package}" ]; then
    echo "Please input package name"
    exit 1
fi

# walk through all files and replace package name
for file in $(find . -type f | grep -v ".git" | grep -v "./install*"); do
    # ignore directory
    if [ -d ${file} ]; then
        continue
    fi
    # find content with "${package}/src*" and replace it with "\"${package}/src*"
    sed -i "s/\${package}/${package}/g" ${file}
done
