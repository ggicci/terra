FROM golang:1.17.7-bullseye

# Install essential commands
RUN apt update && apt install -y sudo

# Fix uid: https://github.com/boxboat/fixuid
RUN addgroup --gid 1000 coder \
    && adduser --uid 1000 --ingroup coder --home /home/coder --shell /bin/bash --disabled-password --gecos "" coder

RUN USER=coder GROUP=coder \
    && curl -sSL https://github.com/boxboat/fixuid/releases/download/v0.5.1/fixuid-0.5.1-linux-amd64.tar.gz | tar -C /usr/local/bin -xzf - \
    && chown root:root /usr/local/bin/fixuid \
    && chmod 4755 /usr/local/bin/fixuid \
    && mkdir -p /etc/fixuid \
    && printf "user: $USER\ngroup: $GROUP\n" > /etc/fixuid/config.yml

# Add coder to sudoers
RUN usermod -aG sudo coder \
    && echo "coder\tALL=(ALL)\tNOPASSWD:ALL" | tee /etc/sudoers.d/coder

USER coder
ENV IO4_TERRA_DEV_IN_CONTAINER=1
ENV GOPATH=/home/coder/go
ENV PATH=/home/coder/go/bin:$PATH

WORKDIR /home/coder/workspace
ENTRYPOINT [ "sh", "-c", "fixuid && sleep infinity" ]
