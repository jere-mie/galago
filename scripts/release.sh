#!/bin/bash

# Define your GitHub repository
REPO_OWNER="jere-mie"
REPO_NAME="galago"

# Check if tag argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <tag>"
    exit 1
fi

# Extract the tag from version.txt
TAG=$(<version.txt)

# Create a release
gh release create $TAG \
    --repo $REPO_OWNER/$REPO_NAME \
    --title "Release $TAG" \
    --notes "Release notes for $TAG"

# Upload built files to the release
for file in $(ls bin/*); do
    gh release upload $TAG $file
done

echo "Release created successfully."
