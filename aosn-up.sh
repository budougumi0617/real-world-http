#!/bin/bash
git push
cd ../../aosn/real-world-http
git pull
cd budougumi0617
git pull origin master
cd ..
git add budougumi0617
git commit -m ":grapes: update budougumi0617"
git push

