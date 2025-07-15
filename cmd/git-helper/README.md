# Git Monitor - Acompanhamento de Alterações

Um script Go que **monitora** alterações no seu projeto e te fornece os comandos necessários para subir as mudanças de forma profissional.

## 🚀 Como usar

### Compilar o script:
```bash
cd cmd/git-helper
go build -o git-monitor main.go
```

### Executar:
```bash
./git-monitor
```

## 📋 Funcionalidades

### 1. 📊 Mostrar status atual
- Exibe o status geral do repositório Git
- Mostra arquivos modificados, staged e não rastreados

### 2. 📋 Detalhar alterações
- Lista arquivos prontos para commit (staged)
- Mostra arquivos modificados não staged
- Exibe arquivos novos não rastreados

### 3. 💡 Mostrar comandos para subir
- **NÃO executa comandos automaticamente**
- Mostra exatamente quais comandos você precisa executar
- Inclui tipos de commit padronizados
- Sugere comandos baseados no estado atual

### 4. 🔄 Monitoramento contínuo
- Monitora alterações a cada 30 segundos
- Alerta quando detecta mudanças
- Mostra comandos necessários automaticamente

## 🎯 O que o Monitor FAZ:

✅ **Monitora** alterações no projeto  
✅ **Mostra** status detalhado  
✅ **Sugere** comandos para subir  
✅ **Detecta** mudanças automaticamente  
✅ **Não executa** comandos Git  

## 🚫 O que o Monitor NÃO FAZ:

❌ **NÃO faz** commits automáticos  
❌ **NÃO executa** comandos Git  
❌ **NÃO modifica** arquivos do projeto  
❌ **NÃO altera** configurações  

## 💡 Exemplo de Comandos Sugeridos

Quando você tem alterações, o monitor mostra:

```
💡 Comandos para Subir as Alterações:
----------------------------------------
📦 1. Adicionar arquivos ao staging:
   git add .

💬 2. Criar commit:
   git commit -m "tipo: descrição do commit"

   Tipos de commit:
   - feat: nova funcionalidade
   - fix: correção de bug
   - refactor: refatoração
   - docs: documentação
   - test: testes
   - chore: manutenção

🚀 3. Fazer push:
   git push

📊 4. Verificar status:
   git status

🔍 5. Ver diferenças:
   git diff
   git diff --cached
```

## 🛡️ Segurança Total

- **100% seguro** - apenas monitora e mostra comandos
- **Não executa** nenhum comando Git automaticamente
- **Não modifica** nenhum arquivo do projeto
- **Apenas leitura** - não faz alterações

## 📝 Como Usar

1. **Compile o script:**
   ```bash
   go build -o git-monitor main.go
   ```

2. **Execute o monitor:**
   ```bash
   ./git-monitor
   ```

3. **Use a opção 3** para ver comandos sugeridos

4. **Execute os comandos manualmente** quando quiser

## 🔄 Monitoramento Contínuo

Para desenvolvimento ativo, use a opção 4:
- Monitora a cada 30 segundos
- Alerta quando detecta mudanças
- Mostra comandos necessários automaticamente
- Pressione Ctrl+C para parar

## 🎯 Benefícios

- ✅ **Controle total** - você decide quando e como fazer commits
- ✅ **Visibilidade completa** - vê todas as alterações
- ✅ **Comandos precisos** - sugestões baseadas no estado atual
- ✅ **Não interfere** - não modifica nada no projeto
- ✅ **Profissional** - facilita commits bem estruturados

---

**Desenvolvido para dar visibilidade total sem interferir no seu workflow! 🎉** 