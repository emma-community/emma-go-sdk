#!/bin/bash

version=$(cat version.txt)

IFS='.' read -r -a version_parts <<< "$version"
version_parts[2]=$((version_parts[2]+1))
new_version="${version_parts[0]}.${version_parts[1]}.${version_parts[2]}"

echo $new_version > version.txt

openapi-generator generate \
  -i https://api.emma.ms/v3/api-docs \
  -g go \
  --additional-properties packageName=emma,packageVersion="$new_version",useTags=true \
  --git-user-id emma-community \
  --git-repo-id emma-go-sdk
