# Use a release script that tags everything at once
# release.sh v0.6.1

#!/bin/bash
VERSION=$1

echo "Releasing version $VERSION"

# Tag everything
git tag core/$VERSION
git tag store/sqlite/$VERSION
git tag store/postgres/$VERSION
git tag store/mongodb/$VERSION
git tag store/filesystem/$VERSION

git push --tags