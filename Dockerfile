FROM golang AS builder

LABEL maintainer="zorro" \
      version="1.0.0"

WORKDIR /build

ENV GOPROXY="https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct" \
    GO111MODULE=on \
    GOSUMDB=off \
    GOARCH=amd64 \
    GOOS=linux

COPY . .

RUN go mod tidy

RUN  go build -o bluebell .



FROM frolvlad/alpine-glibc AS final

WORKDIR /app

COPY --from=builder /build/bluebell /app/
COPY --from=builder /build/wait-for.sh /app/wait-for.sh

#注意这里拷贝文件夹的方式
COPY --from=builder /build/conf /app/conf/
#COPY --from=builder /build/templates /app/templates/
#COPY --from=builder /build/static /app/static/

RUN chmod 755 wait-for.sh

EXPOSE 8001

#ENTRYPOINT ["./bluebell","-f"]

CMD ["./bluebell -f ./conf/config.yaml"]
