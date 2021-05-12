FROM nearprotocol/nearup:latest

COPY --from=nearprotocol/nearcore:latest /usr/local/bin/neard /neard-local/neard
COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
