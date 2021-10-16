#!/bin/bash

# See how to automatize the deletion?
goreleaser && rm -rf dist/ && cd -
