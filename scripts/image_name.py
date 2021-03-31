#! /usr/bin/env python3

import glob
import yaml
import traceback

manifests = glob.glob("./steps/**/manifest.yaml", recursive=True)
print("Manifests: %d" % (len(manifests)))

for m in manifests:
    print("Processing %s" % (m,))
    f = open(m)
    try:
        y = yaml.safe_load(f)
    except:
        print("> Error parsing yaml")
        traceback.print_exc()
        continue

    if 'imageName' in y['metadata']:
        print(" > imageName already set: %s" % (y['metadata']['imageName'],))
        continue