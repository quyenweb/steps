module github.com/stackpulse/public-steps/redis/slowlog

go 1.14

require (
	github.com/caarlos0/env/v6 v6.5.0
	github.com/go-redis/redis/v8 v8.4.4
	github.com/stackpulse/public-steps/redis/base v0.0.0
	github.com/stackpulse/steps-sdk-go v0.0.0-20210314133745-61086c27983f
)

replace github.com/stackpulse/public-steps/redis/base v0.0.0 => ../base
