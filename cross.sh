#!/bin/bash

GOOS=linux make release-server release-client
GOOS=darwin make release-server release-client

