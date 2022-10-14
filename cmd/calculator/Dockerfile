# gcr.io/distroless/static includes certificates necessary for TLS.
FROM gcr.io/distroless/static@sha256:d673e44035b1435c88f63c4b7066501e21fe5c6b111cd9ada7d9301f780b2416 as release
COPY calculator /calculator
ENTRYPOINT [ "/calculator" ]
