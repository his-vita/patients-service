.DEFAULT_GOAL = local

local:
	CONFIG_PATH=config/local.yaml go run cmd/main.go

prod:
	CONFIG_PATH=config/prod.yaml go run cmd/main.go