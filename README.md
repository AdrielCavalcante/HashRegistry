# Hash Registry - Sistema Completo de Verificação de Integridade

Sistema descentralizado para registro e verificação de hashes de documentos na blockchain Ethereum, integrado com armazenamento IPFS via Storacha. Combina **Go + Node.js + Blockchain + IPFS** para garantir integridade e disponibilidade de arquivos.

## **Visão Geral**

O HashRegistry oferece:
- 🔗 **Registro na Blockchain:** Hashes armazenados imutavelmente no Ethereum
- 📁 **Armazenamento IPFS:** Arquivos distribuídos via Storacha
- ⚡ **Verificação Inteligente:** Verifica duplicatas antes do upload
- 🌐 **Interface Web:** Upload e verificação via browser
- 🔍 **Explorer:** Visualização completa da blockchain

## 🏗️ **Arquitetura do Sistema**

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌────────────┐
│   Browser   │ ─> │  Go Server  │ ─> │ Node.js API │ ─> │  Storacha  │
│ (Frontend)  │    │ (Backend)   │    │ (IPFS)      │    │   (IPFS)   │
└─────────────┘    └─────────────┘    └─────────────┘    └────────────┘
                         │
                         ▼
                   ┌────────────┐
                   │  Ethereum  │
                   │ Blockchain │
                   └────────────┘
```

### **Fluxo de Upload:**
1. **Usuário** faz upload via browser
2. **Go** calcula hash SHA256 e verifica se já existe
3. **Se novo:** Go → Node.js → Storacha (IPFS)
4. **Go** registra hash na blockchain
5. **Sucesso:** Arquivo no IPFS + Hash na blockchain

## 📁 **Estrutura do Projeto**

```
HashRegistry/
├── contracts/              # 📋 Smart Contracts
│   └── HashRegistry.sol    # Contrato principal
├── go/                     # 🔧 Bindings Ethereum
│   ├── HashRegistry.go     # Bindings Go gerados
│   ├── contract-address.json
│   └── utils.go
├── web/                    # 🌐 Servidor Go
│   ├── main.go             # Backend principal
│   └── templates/          # Templates HTML
├── storacha-server.js      # 🗄️ API Node.js para IPFS
├── ignition/               # 🚀 Deploy automatizado
│   └── modules/HashRegistry.ts
├── .env                    # ⚙️ Configurações
├── package.json            # 📦 Dependências Node.js
└── hardhat.config.ts       # 🔨 Config Hardhat
```

## 🚀 **Instalação e Configuração**

### **1. Pré-requisitos**
- Node.js 18+
- Go 1.21+
- Conta no [Storacha](https://storacha.network)

### **2. Clonar e Instalar**

```bash
git clone <repository-url>
cd HashRegistry
npm install
```

### **3. Configurar Storacha**

```bash
# Copiar arquivo de exemplo
cp .env.example .env
```

Editar `.env`:
```env
STORACHA_EMAIL=seu-email@gmail.com
STORACHA_SPACE=did:key:z6Mkn... # Ou seu próprio Space DID
```

**Como obter seu próprio Space DID:**
1. Acesse [console.storacha.network](https://console.storacha.network)
2. Crie um Space (ex: "HashRegistryStorage")
3. Copie o DID que aparece (formato: `did:key:z6Mkn...`)

### **4. Blockchain Local**

```bash
# Terminal 1: Iniciar blockchain local
npx hardhat node

# Terminal 2: Deploy do contrato
npx hardhat ignition deploy ignition/modules/HashRegistry.ts --network localhost
```

**Atualizar endereço do contrato:**
```bash
# Copiar o endereço exibido no deploy
echo '{"address": "0x5FbDB2315678afecb367f032d93F642f64180aa3"}' > go/contract-address.json
```

## ▶️ **Como Executar**

### **Opção 1: Tudo junto (Recomendado)**
```bash
npm start
```

### **Opção 2: Separadamente**
```bash
# Terminal 1: Servidor Node.js (Storacha)
npm run storacha

# Terminal 2: Servidor Go (Web + Blockchain)
npm run go
```

### **Acessar a aplicação:**
- 🌐 **Interface Principal:** http://localhost:8080
- 🔍 **Verificar Hash:** http://localhost:8080/verificar  
- 📊 **Blockchain Explorer:** http://localhost:8080/blockchain

## 📋 **Funcionalidades Completas**

### **🔐 Registro de Arquivos**
- ✅ Upload de qualquer tipo de arquivo
- ✅ Hash SHA256 calculado automaticamente
- ✅ Verificação de duplicatas (evita re-upload)
- ✅ Armazenamento distribuído no IPFS
- ✅ Registro imutável na blockchain
- ✅ Backup local opcional

### **🔍 Verificação de Integridade**
- ✅ Verificação por hash SHA256
- ✅ Confirmação via blockchain
- ✅ Histórico completo de registros
- ✅ Informações de timestamp e proprietário

### **🌐 Interface Web**
- ✅ Upload via drag-and-drop
- ✅ Entrada de texto direto
- ✅ Explorer da blockchain em tempo real
- ✅ Estatísticas do sistema
- ✅ Design responsivo (Bootstrap 5)

### **⚡ Otimizações**
- ✅ Verifica blockchain ANTES do upload
- ✅ Evita uploads desnecessários
- ✅ Timeout configurável
- ✅ Logs detalhados para debug

## 🔧 **Configurações Avançadas**

### **Variáveis de Ambiente**

**Blockchain (.env no diretório web/):**
```env
PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
RPC_URL=http://localhost:8545
CHAIN_ID=31337
```

**Storacha (.env na raiz):**
```env
STORACHA_EMAIL=seu-email@gmail.com
STORACHA_SPACE=did:key:z6Mkn...
PORT=3001
```

### **Scripts Disponíveis**

```json
{
  "storacha": "node storacha-server.js",
  "go": "cd web && go run main.go", 
  "start": "concurrently \"npm run storacha\" \"npm run go\""
}
```

## 🔍 **API Endpoints**

### **Go Server (porta 8080)**
- `GET /` - Interface de upload
- `POST /upload` - Processar upload + registro blockchain
- `GET /verificar` - Interface de verificação
- `POST /verificar` - Verificar hash específico
- `GET /blockchain` - Explorer da blockchain

### **Node.js API (porta 3001)**
- `POST /upload?filename=nome.txt` - Upload para Storacha/IPFS

## 🧪 **Testando o Sistema**

### **1. Teste Básico**
1. Acesse http://localhost:8080
2. Faça upload de um arquivo
3. Verifique se aparece no blockchain explorer
4. Tente fazer upload do mesmo arquivo (deve detectar duplicata)

### **2. Verificação de Hash**
1. Copie o hash de um arquivo registrado
2. Acesse http://localhost:8080/verificar
3. Cole o hash e verifique

### **3. Explorar Blockchain**
1. Acesse http://localhost:8080/blockchain
2. Veja todos os hashes registrados
3. Observe timestamps e endereços

## 🛠️ **Tecnologias Utilizadas**

### **Backend**
- **Go 1.21+** - Servidor web e integração blockchain
- **Gin Framework** - API REST em Go
- **go-ethereum** - Cliente Ethereum nativo
- **Node.js 18+** - API para Storacha

### **Blockchain**
- **Solidity ^0.8.28** - Smart contracts
- **Hardhat 3.0+** - Framework de desenvolvimento
- **Ethereum** - Blockchain para registros

### **Armazenamento**
- **Storacha** - Gateway para IPFS
- **IPFS** - Armazenamento distribuído

### **Frontend**
- **HTML5 + Bootstrap 5** - Interface responsiva
- **JavaScript** - Interações dinâmicas

## 📝 **Smart Contract (HashRegistry.sol)**

```solidity
contract HashRegistry {
    mapping(bytes32 => bool) private registeredHashes;
    mapping(bytes32 => address) private hashOwners;
    mapping(bytes32 => uint256) private hashTimestamps;
    
    function registrarHash(bytes32 _hash) external;
    function verificarHash(bytes32 _hash) external view returns (bool);
    function getHashDetails(bytes32 _hash) external view returns (...);
}
```

**Funcionalidades:**
- ✅ Registro de hashes únicos
- ✅ Verificação de existência
- ✅ Rastreamento de proprietário
- ✅ Timestamp de criação
- ✅ Eventos para auditoria

## 🐛 **Troubleshooting**

### **Erro: "Storacha não conectado"**
- Verifique o email no `.env`
- Confirme se tem conta ativa no Storacha
- Verifique se o Space DID está correto

### **Erro: "Erro ao conectar blockchain"**
- Confirme se `npx hardhat node` está rodando
- Verifique se o endereço do contrato está correto
- Confirme se o RPC_URL está acessível

### **Erro: "Cannot connect to Node.js server"**
- Verifique se `npm run storacha` está rodando
- Confirme se a porta 3001 está livre
- Verifique logs do servidor Node.js

## 📊 **Métricas e Monitoramento**

O sistema fornece logs detalhados:
- ✅ Hash calculados e verificações
- ✅ Status de uploads para IPFS
- ✅ Transações blockchain confirmadas  
- ✅ Estatísticas de uso

## 🔒 **Segurança**

- ✅ Hashes SHA256 para integridade
- ✅ Armazenamento imutável na blockchain
- ✅ Distribuição via IPFS (sem ponto único de falha)
- ✅ Verificação de duplicatas
- ✅ Logs de auditoria completos

## 🤝 **Contribuindo**

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 **Licença**

Este projeto está sob a licença ISC.

