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
    # replace package name
    # cause package name may contain "/" which is a special character in sed command
    # so we need to escape it
    tmp_package=$(echo ${package} | sed 's/\//\\\//g')

    sed -i "s/\${package}/${tmp_package}/g" ${file}
done

echo "Done"