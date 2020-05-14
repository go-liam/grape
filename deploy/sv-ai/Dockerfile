FROM liam1803/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD ./build /build/

WORKDIR /build

EXPOSE 7041
EXPOSE 9041

CMD ["./main"]