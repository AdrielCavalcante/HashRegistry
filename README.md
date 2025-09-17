# Hash Registry - Aplicação Web Blockchain

Aplicação web para registro e verificação de hashes de documentos na blockchain Ethereum usando Hardhat, Solidity e Go.

## 🏗️ Estrutura do Projeto

```
hash-registry/
├── contracts/           # Contratos Solidity
│   └── HashRegistry.sol # Contrato principal
├── ignition/           # Deploy via Hardhat Ignition
│   └── modules/
│       └── HashRegistry.ts
├── go/                 # Bindings Go gerados
│   ├── HashRegistry.go # Bindings do contrato
│   ├── HashRegistry.abi
│   ├── contract-address.json
│   └── utils.go
├── web/               # Aplicação web Go
│   ├── main.go        # Servidor web principal
│   ├── templates/     # Templates HTML
│   ├── .env          # Configurações
│   └── go.mod        # Dependências Go
├── hardhat.config.ts # Configuração Hardhat
├── package.json      # Dependências Node.js
└── README.md
```

## 🚀 Como Usar

### 1. Instalar Dependências

```bash
# Dependências Node.js/Hardhat
npm install

# Dependências Go (dentro da pasta web/)
cd web
go mod tidy
```

### 2. Compilar e Deployar Contrato

```bash
# Compilar contrato
npx hardhat compile

# PASSO 1: Iniciar nó local do Hardhat (manter rodando)
npx hardhat node

# PASSO 2: Em outro terminal, fazer deploy
npx hardhat ignition deploy ignition/modules/HashRegistry.ts --network localhost

# PASSO 3: Copiar o endereço do contrato que aparecer e atualizar go/contract-address.json
# Exemplo: se o endereço for 0x5FbDB2315678afecb367f032d93F642f64180aa3
echo '{"address": "SEU_ENDERECO_AQUI", "network": "localhost", "deployedAt": "'$(date -Iseconds)'"}' > go/contract-address.json
```

### 3. Executar Aplicação Web

```bash
cd web
go run main.go
```

Acesse: http://localhost:8080

## 📋 Funcionalidades

- ✅ **Upload de arquivos** para registro na blockchain
- ✅ **Entrada de texto** para registro de hash
- ✅ **Verificação de hashes** existentes
- ✅ **Interface web responsiva** com Bootstrap
- ✅ **Transações na blockchain** via Ethereum

## 🔧 Tecnologias

- **Solidity ^0.8.28** - Smart contracts
- **Hardhat 3.0.6** - Desenvolvimento Ethereum
- **Go 1.21** - Backend da aplicação web
- **Gin** - Framework web Go
- **go-ethereum** - Cliente Ethereum em Go
- **Bootstrap 5** - Frontend responsivo

## 📝 Smart Contract

O contrato `HashRegistry.sol` permite:
- `registrarHash(bytes32)` - Registrar novo hash
- `verificarHash(bytes32)` - Verificar se hash existe
- Evento `HashRegistrado` para auditoria

