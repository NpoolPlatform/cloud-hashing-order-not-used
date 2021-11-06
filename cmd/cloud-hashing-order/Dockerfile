FROM jrei/systemd-centos:7

COPY CloudHashingOrder.viper.yaml /etc/CloudHashingOrder/CloudHashingOrder.viper.yaml
COPY cloud-hashing-order /usr/local/bin/cloud-hashing-order

CMD ["cloud-hashing-order", "run"]
