# Globals
ARG BASE_BRANCH

# 1. Base
FROM us-docker.pkg.dev/stackpulse/public/base:$BASE_BRANCH AS base

# 2. Production
FROM curlimages/curl
ARG STEP_BASEPATH
ARG BASEDIR=/go/src/github.com/stackpulse/steps
COPY --from=base /usr/local/bin/ /usr/local/bin/
COPY ${STEP_BASEPATH}/entrypoint.sh .

ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]