#!/bin/bash
VERSION=$1

echo "Releasing version $VERSION"

git tag scheduler/types/$VERSION
git tag core/$VERSION
git tag store/sqlite/$VERSION
git tag store/postgres/$VERSION
git tag store/mongodb/$VERSION
git tag store/filesystem/$VERSION

git push --tags