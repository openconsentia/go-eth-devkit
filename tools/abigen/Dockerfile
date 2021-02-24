# Copyright 2020 The Open Consentia Contributors
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ethereum/solc:0.4.26 as solc

FROM ubuntu:18.04 as builder

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY --from=solc /usr/bin/solc /usr/local/bin/solc

RUN apt-get update && \
    apt-get -y install wget nodejs npm protobuf-compiler make git && \
    cd /tmp; wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz; cd / && \
    cd /tmp; tar -xvf go1.15.6.linux-amd64.tar.gz; mv go /usr/local/; cd / && \
    go get github.com/ethereum/go-ethereum && \
    cd /opt/go/src/github.com/ethereum/go-ethereum/; make && make devtools; cd /

FROM ubuntu:18.04

COPY --from=builder /usr/local/bin/solc /usr/local/bin/solc
COPY --from=builder /opt/go/bin/abigen /usr/local/bin/abigen

ENTRYPOINT [ "/usr/local/bin/abigen" ]