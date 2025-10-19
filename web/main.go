package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"hashregistry"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️ Arquivo .env não encontrado, usando valores padrão")
	}

	// Configurações com valores padrão para desenvolvimento
	privateKeyHex := getEnvOrDefault("PRIVATE_KEY", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	rpcURL := getEnvOrDefault("RPC_URL", "http://localhost:8545")
	chainID := getEnvOrDefault("CHAIN_ID", "31337")

	chainIDInt, err := strconv.ParseInt(chainID, 10, 64)
	if err != nil {
		log.Fatalf("Erro convertendo CHAIN_ID: %v", err)
	}

	// Criar cliente Ethereum
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Erro conectando ao cliente Ethereum: %v", err)
	}

	// Converter chave privada hex para ECDSA
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Erro convertendo chave privada: %v", err)
	}

	// Configurar transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainIDInt))
	if err != nil {
		log.Fatalf("Erro criando auth: %v", err)
	}

	// Carregar endereço do contrato
	contractAddr, err := hashregistry.LoadContractAddress("../go/contract-address.json")
	if err != nil {
		log.Fatalf("Erro ao carregar endereço do contrato: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddr)
	fmt.Printf("Usando contrato no endereço: %s\n", contractAddress.Hex())

	// Instanciar o contrato
	instance, err := hashregistry.NewHashRegistry(contractAddress, client)
	if err != nil {
		log.Fatalf("Erro instanciando contrato: %v", err)
	}

	// Criar diretório uploads se não existir (backup local)
	if err := os.MkdirAll("uploads", 0755); err != nil {
		log.Fatalf("Erro criando diretório uploads: %v", err)
	}

	log.Println("✅ Sistema inicializado - Go + Node.js + Storacha")

	// Configurar Gin
	r := gin.Default()

	// Adicionar funções personalizadas aos templates
	r.SetFuncMap(template.FuncMap{
		"contains": strings.Contains,
		"add": func(a, b int) int {
			return a + b
		},
	})

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	// Rota principal - formulário de upload
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{"mensagem": ""})
	})

	// Rota para verificar hash
	r.GET("/verificar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "verificar.html", gin.H{"resultado": ""})
	})

	// Rota para visualizar blockchain
	r.GET("/blockchain", func(c *gin.Context) {
		// Estrutura para dados de hash
		type HashData struct {
			Hash          string
			Owner         string
			Timestamp     uint64
			FormattedTime string
			CID           string
		}

		// Estrutura para estatísticas
		type Stats struct {
			TotalHashes     string
			ContractAddress string
			Network         string
		}

		// Buscar total de hashes
		totalHashes, err := instance.GetTotalHashes(&bind.CallOpts{})
		if err != nil {
			c.HTML(http.StatusOK, "blockchain.html", gin.H{
				"error": fmt.Sprintf("Erro ao buscar total de hashes: %v", err),
			})
			return
		}

		// Buscar todos os hashes
		var hashes []HashData
		for i := uint64(0); i < totalHashes.Uint64(); i++ {
			hash, err := instance.GetHashByIndex(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Printf("Erro ao buscar hash índice %d: %v", i, err)
				continue
			}

			// Buscar detalhes do hash
			details, err := instance.GetHashDetails(&bind.CallOpts{}, hash)
			if err != nil || !details.Exists {
				log.Printf("Erro ao buscar detalhes do hash %s: %v", common.BytesToHash(hash[:]).Hex(), err)
				continue
			}

			// Formatar timestamp
			t := time.Unix(details.Timestamp.Int64(), 0)
			formattedTime := t.Format("02/01/2006 15:04:05")

			hashes = append(hashes, HashData{
				Hash:          common.BytesToHash(hash[:]).Hex(),
				Owner:         details.Owner.Hex(),
				Timestamp:     uint64(details.Timestamp.Int64()),
				FormattedTime: formattedTime,
				CID:           details.Cid,
			})
		}

		// Reverter ordem para mostrar mais recentes primeiro
		for i, j := 0, len(hashes)-1; i < j; i, j = i+1, j-1 {
			hashes[i], hashes[j] = hashes[j], hashes[i]
		}

		// Criar estatísticas
		stats := Stats{
			TotalHashes:     fmt.Sprintf("%d", totalHashes.Uint64()),
			ContractAddress: contractAddress.Hex(),
			Network:         "Localhost (Hardhat)",
		}

		c.HTML(http.StatusOK, "blockchain.html", gin.H{
			"hashes": hashes,
			"stats":  stats,
		})
	})

	// Rota para processar upload e registro de hash
	r.POST("/upload", func(c *gin.Context) {
		texto := c.PostForm("texto")
		var data []byte
		var err error
		var filename string

		if texto != "" {
			data = []byte(texto)
			filename = "texto digitado"
		} else {
			var file *multipart.FileHeader
			file, err = c.FormFile("arquivo")
			if err != nil {
				c.HTML(http.StatusOK, "upload.html", gin.H{"mensagem": "❌ Erro ao ler o arquivo!"})
				return
			}

			filename = file.Filename

			var f io.ReadCloser
			f, err = file.Open()
			if err != nil {
				c.HTML(http.StatusOK, "upload.html", gin.H{"mensagem": "❌ Erro ao abrir o arquivo!"})
				return
			}
			defer f.Close()

			data, err = io.ReadAll(f)
			if err != nil {
				c.HTML(http.StatusOK, "upload.html", gin.H{"mensagem": "❌ Erro ao ler o conteúdo do arquivo!"})
				return
			}
		}

		// Gera hash SHA256 do conteúdo
		hashBytes := sha256.Sum256(data)
		hash := common.BytesToHash(hashBytes[:])

		log.Printf("📦 Hash calculado: %s", hash.Hex())

		// Verificar se já existe antes de fazer upload
		exists, err := instance.VerificarHash(&bind.CallOpts{}, hash)
		if err != nil {
			log.Printf("❌ Erro ao verificar hash: %v", err)
			c.HTML(http.StatusOK, "upload.html", gin.H{
				"mensagem": fmt.Sprintf("❌ Erro ao verificar hash: %v", err),
			})
			return
		}

		if exists {
			log.Printf("⚠️ Hash já existe, pulando upload")
			c.HTML(http.StatusOK, "upload.html", gin.H{
				"mensagem": "⚠️ Este arquivo já está registrado na blockchain!",
				"hash":     hash.Hex(),
				"arquivo":  filename,
			})
			return
		}

		// Só faz upload se o hash não existir
		log.Printf("📤 Hash novo, processando: %s (%d bytes)", filename, len(data))

		var cid string

		// Só fazer upload para Storacha se for um arquivo real, não texto digitado
		if filename != "texto digitado" {
			log.Printf("📤 Fazendo upload do arquivo para Storacha: %s", filename)
			var err error
			cid, err = uploadToNodeJS(data, filename)
			if err != nil {
				log.Printf("❌ Erro no upload: %v", err)
				c.HTML(http.StatusOK, "upload.html", gin.H{
					"mensagem": fmt.Sprintf("❌ Erro no upload: %v", err),
				})
				return
			}
			log.Printf("✅ Upload concluído - CID: %s", cid)
		} else {
			log.Printf("📝 Texto digitado - pulando upload para Storacha")
			cid = "N/A - texto não armazenado"
		}

		// Backup local
		if filename != "texto digitado" {
			caminho := filepath.Join("uploads", filepath.Base(filename))
			if err = os.WriteFile(caminho, data, 0644); err != nil {
				log.Printf("Aviso: erro ao salvar backup: %v", err)
			}
		}

		// Registrar novo hash na blockchain
		tx, err := instance.RegistrarHash(auth, hash, cid)
		if err != nil {
			log.Printf("❌ Erro ao registrar hash: %v", err)
			c.HTML(http.StatusOK, "upload.html", gin.H{
				"mensagem": fmt.Sprintf("❌ Erro ao registrar hash na blockchain: %v", err),
			})
			return
		}

		// Aguardar confirmação da transação
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Printf("❌ Erro aguardando confirmação: %v", err)
		}

		// Mensagem de sucesso personalizada baseada no tipo
		var mensagem string
		if filename == "texto digitado" {
			mensagem = "✅ Hash do texto registrado na blockchain com sucesso!"
		} else {
			mensagem = "✅ Hash registrado na blockchain e arquivo armazenado no Storacha com sucesso!"
		}

		c.HTML(http.StatusOK, "upload.html", gin.H{
			"mensagem": mensagem,
			"hash":     hash.Hex(),
			"tx":       tx.Hash().Hex(),
			"bloco":    receipt.BlockNumber.Uint64(),
			"arquivo":  filename,
			"cid":      cid,
		})
	})

	// Rota para verificar se um hash existe
	r.POST("/verificar", func(c *gin.Context) {
		hashStr := c.PostForm("hash")

		if len(hashStr) != 66 || hashStr[:2] != "0x" {
			c.HTML(http.StatusOK, "verificar.html", gin.H{
				"resultado": "❌ Hash inválido! Deve ter 66 caracteres e começar com 0x",
			})
			return
		}

		hash := common.HexToHash(hashStr)

		exists, err := instance.VerificarHash(&bind.CallOpts{}, hash)
		if err != nil {
			c.HTML(http.StatusOK, "verificar.html", gin.H{
				"resultado": fmt.Sprintf("❌ Erro ao verificar hash: %v", err),
			})
			return
		}

		if exists {
			c.HTML(http.StatusOK, "verificar.html", gin.H{
				"resultado": "✅ Hash encontrado! Este documento foi registrado na blockchain.",
				"hash":      hashStr,
			})
		} else {
			c.HTML(http.StatusOK, "verificar.html", gin.H{
				"resultado": "❌ Hash não encontrado. Este documento não foi registrado.",
				"hash":      hashStr,
			})
		}
	})

	fmt.Println("🚀 Servidor iniciado em http://localhost:8080")
	fmt.Println("📄 Upload de arquivos: http://localhost:8080/")
	fmt.Println("🔍 Verificar hash: http://localhost:8080/verificar")
	fmt.Println("📋 Blockchain Explorer: http://localhost:8080/blockchain")

	r.Run(":8080")
}

// getEnvOrDefault retorna a variável de ambiente ou um valor padrão
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// uploadToNodeJS envia arquivo para servidor Node.js que faz upload para Storacha
func uploadToNodeJS(data []byte, filename string) (string, error) {
	// URL do servidor Node.js local
	nodeServerURL := "http://localhost:3001/upload"

	log.Printf("🚀 Enviando para servidor Node.js: %s (arquivo: %s, %d bytes)", nodeServerURL, filename, len(data))

	// Codificar filename para URL
	encodedFilename := url.QueryEscape(filename)

	// Criar requisição para servidor Node.js
	req, err := http.NewRequest("POST", nodeServerURL+"?filename="+encodedFilename, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	// Headers
	req.Header.Set("Content-Type", "application/octet-stream")

	log.Printf("🔍 URL completa: %s", req.URL.String())
	log.Printf("🔍 Content-Type: %s", req.Header.Get("Content-Type"))
	log.Printf("🔍 Content-Length: %d", len(data))

	// Fazer requisição
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao conectar com servidor Node.js: %v", err)
	}
	defer resp.Body.Close()

	// Ler resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	log.Printf("📡 Resposta do servidor Node.js - Status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Erro HTTP %d - Body: %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("erro do servidor Node.js (%d): %s", resp.StatusCode, string(body))
	}

	// Parse da resposta JSON
	var response struct {
		Success bool   `json:"success"`
		CID     string `json:"cid"`
		URL     string `json:"url"`
		Error   string `json:"error"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	if !response.Success {
		return "", fmt.Errorf("erro no servidor Node.js: %s", response.Error)
	}

	log.Printf("✅ Upload concluído - CID: %s", response.CID)

	return response.CID, nil
}
