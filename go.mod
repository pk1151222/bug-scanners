module github.com/pk1151222/bug-scanners

go 1.20

require (
    github.com/gin-gonic/gin v1.8.1
    github.com/projectdiscovery/subfinder/v2 v2.5.5
    github.com/go-echarts/go-echarts v1.6.0
    github.com/nats-io/nats.go v1.14.2
    github.com/projectdiscovery/nuclei/v2/v2.5.2
    github.com/stretchr/testify v1.8.4
    golang.org/x/net v0.11.0
    github.com/sirupsen/logrus v1.9.0
    github.com/dchest/uniuri v1.0.0
    github.com/go-sql-driver/mysql v1.7.0
    github.com/elastic/go-elasticsearch/v7 v7.15.0
    github.com/matryer/xbar v0.0.0-20210113123842-f4ea8ffb93d2
    github.com/robfig/cron/v3 v3.0.0
    golang.org/x/oauth2 v0.0.0-20231021173400-848a200a4427
)

replace (
    github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
    github.com/projectdiscovery/subfinder/v2 => github.com/projectdiscovery/subfinder/v2 v2.5.5
)
