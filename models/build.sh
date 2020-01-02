#!/bin/bash
CGO_ENABLED=0 swagger generate spec -o ../apimachinery/models.json -w . -m
