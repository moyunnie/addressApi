FROM ubuntu:latest
LABEL authors="Sierting"

ENTRYPOINT ["top", "-b"]