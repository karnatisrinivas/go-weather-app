
FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod go.sum ./
# RUN go mod download github.com/sagikazarmark/crypt@v0.9.0
COPY . .
RUN go build -o weather-app
# WORKDIR /app
COPY index.html .
EXPOSE 8080
CMD ["./weather-app"]
