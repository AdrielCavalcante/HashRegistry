#!/bin/bash

echo "🔥 Hash Registry - Deploy Automático"
echo "======================================"

# Verificar se o nó já está rodando
if curl -s http://localhost:8545 > /dev/null 2>&1; then
    echo "✅ Nó do Hardhat já está rodando em http://localhost:8545"
else
    echo "⚠️  Nó do Hardhat não está rodando."
    echo "📝 Por favor, execute em outro terminal:"
    echo "   npx hardhat node"
    echo ""
    echo "⏳ Aguardando nó ficar disponível..."
    
    # Aguardar até o nó estar disponível
    while ! curl -s http://localhost:8545 > /dev/null 2>&1; do
        sleep 1
        echo -n "."
    done
    echo ""
fi

echo "🚀 Fazendo deploy do contrato..."
npx hardhat run scripts/deploy-simple.js --network localhost

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Deploy concluído!"
    echo "🌐 Agora você pode executar a aplicação web:"
    echo "   cd web && go run main.go"
else
    echo ""
    echo "❌ Erro no deploy!"
    echo "📝 Certifique-se de que o nó do Hardhat está rodando:"
    echo "   npx hardhat node"
fi