#!/bin/bash
echo "Resetting samples folder"
aws s3 sync s3://qc-samples ./samples --delete