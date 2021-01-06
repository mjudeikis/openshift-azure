#!/bin/bash -e

cleanup() {
    set +e

    generate_artifacts

    delete
}

trap cleanup EXIT

echo "Prepate CI"

mkdir -p  $(pwd)/secrets
cp -R /secrets $(pwd)/secrets
chown -R $HOST_USER_ID:$HOST_GROUP_ID $(pwd)/secrets

. hack/tests/ci-prepare.sh


echo "RESOURCEGROUP is $RESOURCEGROUP"

ls -la .
ls -la $(pwd)/secrets

start_monitoring
set_build_images

make create

hack/e2e.sh
