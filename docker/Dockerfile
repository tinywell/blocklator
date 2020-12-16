## node 前端编译
FROM node:alpine as nodeBuilder

WORKDIR /app

COPY ./front ./

# RUN npm config set registry https://registry.npm.taobao.org  && \
#     npm install

RUN npm install

RUN npm run build

## go 后端编译
FROM golang:1.13-alpine as goBuilder

WORKDIR /app

COPY ./ ./

# RUN go env -w GOPROXY="https://goproxy.cn,direct" && \
#     go build -v -o blocklator ./cmd/blocklator.go

RUN go build -v -o blocklator ./cmd/blocklator.go

## 构建运行镜像
FROM alpine:latest

WORKDIR /app

COPY --from=nodebuilder /app/dist/ ./dist
COPY --from=goBuilder /app/blocklator ./

EXPOSE 80

CMD [ "./blocklator","server","-m","-p","80","-s","./dist"]