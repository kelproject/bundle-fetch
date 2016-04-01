#!/bin/bash
set -ev

here=$(cd "$(dirname "${BASH_SOURCE}")"; pwd -P)
. $here/_env.sh

gsutil cp bin/bundle-fetch gs://release.kelproject.com/binaries/bundle-fetch/bundle-fetch-linux-amd64-$BUILD_TAG
