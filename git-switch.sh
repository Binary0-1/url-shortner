#!/bin/bash
if [ "$1" == "work" ]; then
  git config user.name "Prasan Internshala"
  git config user.email "prasan@internshala.com"
  echo "Switched to WORK account"
elif [ "$1" == "personal" ]; then
  git config user.name "Prasan"
  git config user.email "prasanmishra330@gmail.com"
  echo "Switched to PERSONAL account"
else
  echo "Usage: ./git-switch.sh [work|personal]"
fi
