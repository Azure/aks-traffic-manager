FROM ubuntu:16.04

RUN apt-get update && apt-get install -y ca-certificates

COPY aks-traffic-manager /aks-traffic-manager

ENTRYPOINT ["/aks-traffic-manager"]