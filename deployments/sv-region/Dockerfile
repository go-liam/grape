FROM liam1803/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD build /build/

WORKDIR /build

EXPOSE 7409
EXPOSE 9409

CMD ["./main"]