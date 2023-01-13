#!/bin/bash

IMAGE_NAME="weaviate"
function release() {
  # for multi-platform build

  IMAGE_ID=ghcr.io/$REPO_OWNER/$IMAGE_NAME
  TAG=

  git_hash=$(echo "$GITHUB_SHA" | cut -c1-7)

  weaviate_version="$(jq -r '.info.version' < openapi-specs/schema.json)"
  echo weaviate_version=$weaviate_version

  if [ "$GITHUB_REF_NAME" = "master" ]; then
    TAG="$${weaviate_version}-${git_hash}"
  elif [  "$GITHUB_REF_TYPE" == "tag" ]; then
        if [ "$GITHUB_REF_NAME" != "v$weaviate_version" ]; then
            echo "The release tag ($GITHUB_REF_NAME) and Weaviate version (v$weaviate_version) are not equal! Can't release."
            return 1
        fi
        TAG="${weaviate_version}"
  else
    pr_title="$(echo -n "$PR_TITLE" | tr '[:upper:]' '[:lower:]' | tr -c -s '[:alnum:]' '-' | sed 's/-$//g')"
    TAG="preview-${pr_title}-${git_hash}"
  fi

  echo IMAGE_ID=$IMAGE_ID
  echo TAG=$TAG
  echo IMAGE_NAME=$IMAGE_NAME

  docker build . --file Dockerfile --tag $IMAGE_NAME  --label "tag=${TAG}"
  docker tag $IMAGE_NAME $IMAGE_ID:$TAG
  docker push $IMAGE_ID:$TAG
}

release
