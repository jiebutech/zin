FROM 192.168.1.202:5000/golang:builder-1.0 as builder

WORKDIR /build

COPY . .

RUN yq -i '.app.listen = ":80"' ./conf/application.test.yaml

RUN go mod tidy
RUN go build -o app-runner .
RUN chmod +x app-runner


FROM 192.168.1.202:5000/binary-runner:alpine

COPY --from=builder /build/app-runner ./app
COPY --from=builder /build/conf/application.test.yaml ./application.yaml

CMD ./app
