ARG CURRENT_BRANCH
FROM us-docker.pkg.dev/stackpulse/public/gcloud/base:latest
USER sp
ARG STEP_BASEPATH
COPY ${STEP_BASEPATH}/gcloud-entrypoint.py /tmp/
ENTRYPOINT ["python", "/tmp/gcloud-entrypoint.py"]