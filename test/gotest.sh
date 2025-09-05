#!/bin/bash
echo "🧪 Executando testes com cobertura..."
go test -v -cover -coverprofile=coverage.out ../...

echo "📊 Relatório de cobertura por função:"
go tool cover -func=coverage.out

echo "🌐 Gerando relatório HTML..."
go tool cover -html=coverage.out -o coverage.html

echo "📈 Cobertura total:"
COVERAGE=$(go tool cover -func=coverage.out | grep "total:" | awk '{print $3}')
echo "Cobertura: $COVERAGE"

# Verificar se atingiu threshold mínimo (ex: 80%)
COVERAGE_NUM=$(echo $COVERAGE | sed 's/%//')
THRESHOLD=80

if (( $(echo "$COVERAGE_NUM >= $THRESHOLD" | bc -l) )); then
    echo "✅ Cobertura OK ($COVERAGE >= ${THRESHOLD}%)"
    exit 0
else
    echo "❌ Cobertura insuficiente ($COVERAGE < ${THRESHOLD}%)"
    exit 1
fi
