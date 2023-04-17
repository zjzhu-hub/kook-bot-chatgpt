FROM docker.io/library/alpine:latest

COPY build/kook-bot-chatgpt /usr/bin/kook-bot-chatgpt

ENTRYPOINT ["/usr/bin/kook-bot-chatgpt"]

