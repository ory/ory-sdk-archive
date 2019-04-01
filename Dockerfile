FROM scratch

COPY ory /
COPY .releaser/LICENSE.txt /LICENSE.txt

ENTRYPOINT ["/ory"]
