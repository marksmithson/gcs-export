#!/bin/bash

docker run --rm -v $(pwd)/secrets:/config -v $(pwd)/data/in:/in -e GOOGLE_APPLICATION_CREDENTIALS=/config/service-account.json gcs-export /in/input.json gs://data-drover-export/input.json
