list-recipes:
    just -l

start:
    nginx -g 'daemon off;' -c $(pwd)/nginx.conf

start-daemon:
    nginx -c $(pwd)/nginx.conf

receive:
    websocat ws://127.0.0.1:8081/sub?id=vcs-poc

send-post:
    curl --request POST --data "test message $(date +%H:%M:%S)" http://127.0.0.1:8081/pub?id=vcs-poc

send-ws:
    websocat ws://127.0.0.1:8081/pub?id=vcs-poc