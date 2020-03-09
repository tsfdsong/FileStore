######## Start a new stage from scratch #######
FROM ubuntu:18.04

COPY ./geth /usr/bin

# Document that the service listens on port 8080.
EXPOSE 8545

CMD ["geth","--rpc", "--rpcapi", "web3,net,eth,personal,miner,contract", "--rpcport=8545", "--rpcaddr", "0.0.0.0","--rpcvhosts=*", "--bootnodes", "enode://32414288558e4d68adcec12d71dcf3f4bf08ddfd2d237f36e0868332f9a91c8c356c26417f92e64d91643a8ce0669366617dfd1ba1a573fc3ef9d446e8123ae7@106.14.136.219:30303"]