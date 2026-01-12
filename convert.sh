#!/bin/bash

ffmpeg -framerate 60 -pattern_type glob -i '*.png' -c:v libx264 -r 30 output.mp4

