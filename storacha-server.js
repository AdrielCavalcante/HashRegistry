import express from 'express'
import { create } from '@storacha/client'
import dotenv from 'dotenv'

// Carregar .env
dotenv.config()

const app = express()
const PORT = process.env.PORT || 3001

let storachaClient = null

console.log('🚀 Iniciando servidor Node.js...')
console.log(`📧 Email configurado: ${process.env.STORACHA_EMAIL || 'NÃO CONFIGURADO'}`)
console.log(`🏠 Space configurado: ${process.env.STORACHA_SPACE || 'NÃO CONFIGURADO'}`)

// Inicializar Storacha
async function initStoracha() {
  try {
    console.log('⏳ Conectando ao Storacha...')
    storachaClient = await create()
    await storachaClient.login(process.env.STORACHA_EMAIL || 'seu-email@exemplo.com')
    
    // Configurar o space atual
    if (process.env.STORACHA_SPACE) {
      await storachaClient.setCurrentSpace(process.env.STORACHA_SPACE)
      console.log('🏠 Space definido como atual')
    }
    
    console.log('✅ Storacha conectado com sucesso!')
  } catch (error) {
    console.error('❌ Erro ao conectar Storacha:', error.message)
    console.error('   Verifique se o email e space estão corretos no .env')
  }
}

// Middleware para receber dados raw
app.use(express.raw({ type: 'application/octet-stream', limit: '50mb' }))

// Rota única: upload
app.post('/upload', async (req, res) => {
  try {
    console.log('📨 Recebendo upload...')
    
    if (!storachaClient) {
      console.log('❌ Cliente Storacha não está conectado')
      return res.status(500).json({ error: 'Storacha não conectado' })
    }

    const filename = req.query.filename || 'file'
    const data = req.body

    if (!data || data.length === 0) {
      return res.status(400).json({ error: 'Dados vazios' })
    }

    console.log(`📤 Processando: ${filename} (${data.length} bytes)`)

    // Upload para Storacha
    const file = new File([data], filename)
    console.log(`📤 Fazendo upload: ${filename} (${data.length} bytes)`)
    
    const result = await storachaClient.uploadFile(file)
    console.log('🔍 Resposta Storacha:', result)
    
    // O Storacha retorna um CID diretamente, não um objeto com .cid
    let cid
    if (result && typeof result.toString === 'function') {
      cid = result.toString()
    } else if (result && result.cid && typeof result.cid.toString === 'function') {
      cid = result.cid.toString()
    } else {
      throw new Error('Formato de resposta inesperado do Storacha')
    }
    
    console.log(`✅ Upload: ${filename} → ${cid}`)

    res.json({
      success: true,
      cid: cid,
      url: `https://gateway.storacha.network/ipfs/${cid}`
    })

  } catch (error) {
    console.error('❌ Erro upload:', error.message)
    console.error('❌ Stack:', error.stack)
    res.status(500).json({ error: error.message })
  }
})

// Iniciar
initStoracha().then(() => {
  app.listen(PORT, () => {
    console.log(`✅ Servidor rodando na porta ${PORT}`)
    console.log(`🔗 Endpoint: http://localhost:${PORT}/upload`)
    console.log('🎯 Pronto para receber uploads!')
  })
}).catch(error => {
  console.error('❌ Erro ao inicializar servidor:', error)
  process.exit(1)
})