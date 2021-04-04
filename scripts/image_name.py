#! /usr/bin/env python3

import glob
import yaml
import traceback
import sys
import os.path

DOCKER_REPO_PREFIX = "us-docker.pkg.dev/stackpulse/public"
STEPS_PREFIX = "./steps/"
MANIFEST_FILE = "/manifest.yaml"

def get_image_name(path):
    path = path.replace(STEPS_PREFIX, "")
    return os.path.join(DOCKER_REPO_PREFIX, os.path.dirname(path))

def main():
    if len(sys.argv) == 2:
        manifests = [sys.argv[1]]
    else:
        manifests = glob.glob(STEPS_PREFIX + "**/manifest.yaml", recursive=True)

    print("Manifests: %d" % (len(manifests)))

    for m in manifests:
        print("Processing %s" % (m,))
        f = open(m)
        try:
            y = yaml.safe_load(f)
        except:
            print("> Error parsing yaml %s" % (m,))
            traceback.print_exc()
            return

        if 'imageName' in y['metadata']:
            print(" > imageName already set: %s" % (y['metadata']['imageName'],))
            continue

        print("New Image: %s" % (get_image_name(m),))

if __name__ == "__main__":
    main()