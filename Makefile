SHELL := /bin/bash

COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html
COVERAGE_THRESHOLD := 80
PACKAGES := ./...

# Cores para output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

.DEFAULT_GOAL := help


.PHONY: help
help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2}'


.PHONY: gotest
gotest: ## Executa os testes com cobertura e verifica o limite mínimo
	@echo "🧪 Executando testes com cobertura..."
	@go test -v -cover -coverprofile=$(COVERAGE_FILE) $(PACKAGES) || (rm -f $(COVERAGE_FILE); exit 1)
	@echo "📊 Relatório de cobertura por função:"
	@go tool cover -func=$(COVERAGE_FILE)
	@echo "🌐 Gerando relatório HTML..."
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "📈 Cobertura total:"
	@COVERAGE=$$(go tool cover -func=$(COVERAGE_FILE) | grep "total:" | awk '{print $$3}'); echo "Cobertura: $$COVERAGE"; \
	COVERAGE_NUM=$${COVERAGE%\%}; \
	if (( $$(echo "$$COVERAGE_NUM >= $(COVERAGE_THRESHOLD)" | bc -l) )); then echo "✅ Cobertura OK ($$COVERAGE >= $(COVERAGE_THRESHOLD)%)"; else echo "❌ Cobertura insuficiente ($$COVERAGE < $(COVERAGE_THRESHOLD)%)"; fi

.PHONY: coverage
coverage: ## Gera o relatório HTML de cobertura
	@echo "🌐 Gerando relatório HTML a partir de $(COVERAGE_FILE)..."
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML) && echo "Gerado $(COVERAGE_HTML)"

.PHONY: check-coverage
check-coverage: ## Verifica a cobertura atual sem rodar os testes
	@COVERAGE=$$(go tool cover -func=$(COVERAGE_FILE) 2>/dev/null | grep "total:" | awk '{print $$3}'); \
	if [ -z "$$COVERAGE" ]; then echo "⚠️  Arquivo de cobertura não encontrado ou testes falharam"; exit 1; fi; \
	echo "Cobertura: $$COVERAGE"; \
	COVERAGE_NUM=$${COVERAGE%\%}; \
	if (( $$(echo "$$COVERAGE_NUM >= $(COVERAGE_THRESHOLD)" | bc -l) )); then echo "✅ Cobertura OK ($$COVERAGE >= $(COVERAGE_THRESHOLD)%)"; else echo "❌ Cobertura insuficiente ($$COVERAGE < $(COVERAGE_THRESHOLD)%)"; fi
