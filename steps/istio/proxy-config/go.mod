module github.com/stackpulse/public-steps/istio/proxy-config

go 1.14

require (
	github.com/caarlos0/env/v6 v6.5.0
	github.com/stackpulse/public-steps/istio/base v0.0.0
	github.com/stackpulse/steps-sdk-go v0.0.0-20210314133745-61086c27983f
)

replace github.com/stackpulse/steps-sdk-go v0.0.0-20210314133745-61086c27983f => ../../common

replace github.com/stackpulse/public-steps/istio/base v0.0.0 => ../base
