#!/bin/bash
echo "Uploading current changes to samples"
aws s3 sync ./samples s3://qc-samples