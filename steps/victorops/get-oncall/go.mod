module github.com/stackpulse/steps/victorops/get-oncall

go 1.14

replace github.com/stackpulse/steps/victorops/base v0.0.0 => ../base

require (
	github.com/caarlos0/env/v6 v6.5.0
	github.com/stackpulse/steps-sdk-go v0.0.0-20210314133745-61086c27983f
	github.com/stackpulse/steps/victorops/base v0.0.0
)
