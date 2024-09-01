FROM node:22

RUN apt update && apt install -y less man-db sudo

# Ensure default `node` user has access to `sudo`
ARG USERNAME=node
RUN echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

RUN npm install -g textlint@latest \
    textlint-rule-preset-ja-technical-writing \
    textlint-rule-preset-ja-spacing

ENV DEVCONTAINER=true
