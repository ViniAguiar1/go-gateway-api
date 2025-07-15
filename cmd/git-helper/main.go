package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type GitMonitor struct {
	projectPath string
}

func NewGitMonitor() *GitMonitor {
	return &GitMonitor{
		projectPath: ".",
	}
}

func (gm *GitMonitor) runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = gm.projectPath
	output, err := cmd.Output()
	return string(output), err
}

func (gm *GitMonitor) getStatus() (string, error) {
	return gm.runCommand("git", "status", "--porcelain")
}

func (gm *GitMonitor) getStagedFiles() (string, error) {
	return gm.runCommand("git", "diff", "--cached", "--name-only")
}

func (gm *GitMonitor) getUnstagedFiles() (string, error) {
	return gm.runCommand("git", "diff", "--name-only")
}

func (gm *GitMonitor) getUntrackedFiles() (string, error) {
	return gm.runCommand("git", "ls-files", "--others", "--exclude-standard")
}

func (gm *GitMonitor) displayStatus() {
	fmt.Println("🔄 Status do Repositório Git")
	fmt.Println(strings.Repeat("=", 50))

	// Status geral
	status, err := gm.runCommand("git", "status")
	if err != nil {
		fmt.Printf("❌ Erro ao obter status: %v\n", err)
		return
	}
	fmt.Println(status)
	fmt.Println(strings.Repeat("=", 50))
}

func (gm *GitMonitor) showDetailedChanges() {
	fmt.Println("\n📋 Detalhamento das Alterações:")
	fmt.Println(strings.Repeat("-", 30))

	// Arquivos staged
	staged, err := gm.getStagedFiles()
	if err == nil && strings.TrimSpace(staged) != "" {
		fmt.Println("✅ Arquivos prontos para commit:")
		files := strings.Split(strings.TrimSpace(staged), "\n")
		for _, file := range files {
			if file != "" {
				fmt.Printf("   📄 %s\n", file)
			}
		}
	}

	// Arquivos modificados não staged
	unstaged, err := gm.getUnstagedFiles()
	if err == nil && strings.TrimSpace(unstaged) != "" {
		fmt.Println("\n🔄 Arquivos modificados (não staged):")
		files := strings.Split(strings.TrimSpace(unstaged), "\n")
		for _, file := range files {
			if file != "" {
				fmt.Printf("   📝 %s\n", file)
			}
		}
	}

	// Arquivos não rastreados
	untracked, err := gm.getUntrackedFiles()
	if err == nil && strings.TrimSpace(untracked) != "" {
		fmt.Println("\n🆕 Arquivos novos (não rastreados):")
		files := strings.Split(strings.TrimSpace(untracked), "\n")
		for _, file := range files {
			if file != "" {
				fmt.Printf("   🆕 %s\n", file)
			}
		}
	}
}

func (gm *GitMonitor) showCommands() {
	fmt.Println("\n💡 Comandos para Subir as Alterações:")
	fmt.Println(strings.Repeat("-", 40))

	// Verificar se há arquivos não staged
	unstaged, err := gm.getUnstagedFiles()
	hasUnstaged := err == nil && strings.TrimSpace(unstaged) != ""

	// Verificar se há arquivos não rastreados
	untracked, err := gm.getUntrackedFiles()
	hasUntracked := err == nil && strings.TrimSpace(untracked) != ""

	// Verificar se há arquivos staged
	staged, err := gm.getStagedFiles()
	hasStaged := err == nil && strings.TrimSpace(staged) != ""

	if hasUnstaged || hasUntracked {
		fmt.Println("📦 1. Adicionar arquivos ao staging:")
		if hasUntracked {
			fmt.Println("   git add .")
		} else {
			fmt.Println("   git add <arquivo_específico>")
		}
		fmt.Println()
	}

	if hasStaged {
		fmt.Println("💬 2. Criar commit:")
		fmt.Println("   git commit -m \"tipo: descrição do commit\"")
		fmt.Println()
		fmt.Println("   Tipos de commit:")
		fmt.Println("   - feat: nova funcionalidade")
		fmt.Println("   - fix: correção de bug")
		fmt.Println("   - refactor: refatoração")
		fmt.Println("   - docs: documentação")
		fmt.Println("   - test: testes")
		fmt.Println("   - chore: manutenção")
		fmt.Println()
	}

	if hasStaged || hasUnstaged || hasUntracked {
		fmt.Println("🚀 3. Fazer push:")
		fmt.Println("   git push")
		fmt.Println()
		fmt.Println("   Ou se for uma nova branch:")
		fmt.Println("   git push -u origin <nome-da-branch>")
		fmt.Println()
	}

	fmt.Println("📊 4. Verificar status:")
	fmt.Println("   git status")
	fmt.Println()

	fmt.Println("🔍 5. Ver diferenças:")
	fmt.Println("   git diff")
	fmt.Println("   git diff --cached")
}

func (gm *GitMonitor) showMenu() {
	fmt.Println("\n🛠️  Git Monitor - Menu de Opções")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("1. 📊 Mostrar status atual")
	fmt.Println("2. 📋 Detalhar alterações")
	fmt.Println("3. 💡 Mostrar comandos para subir")
	fmt.Println("4. 🔄 Monitoramento contínuo")
	fmt.Println("5. 🚪 Sair")
	fmt.Println(strings.Repeat("=", 40))
}

func (gm *GitMonitor) continuousMonitoring() {
	fmt.Println("🔄 Iniciando monitoramento contínuo...")
	fmt.Println("Pressione Ctrl+C para parar")
	fmt.Println("Verificando alterações a cada 30 segundos...")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			status, err := gm.getStatus()
			if err == nil && strings.TrimSpace(status) != "" {
				fmt.Printf("\n[%s] 🔄 Alterações detectadas!\n", time.Now().Format("15:04:05"))
				gm.displayStatus()
				gm.showDetailedChanges()
				gm.showCommands()
				fmt.Println(strings.Repeat("=", 50))
			}
		}
	}
}

func main() {
	monitor := NewGitMonitor()

	fmt.Println("🚀 Git Monitor - Acompanhamento de Alterações")
	fmt.Println("Versão 1.0.0 - Apenas Monitoramento")
	fmt.Println()

	for {
		monitor.showMenu()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Escolha uma opção (1-5): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			monitor.displayStatus()
		case "2":
			monitor.showDetailedChanges()
		case "3":
			monitor.showCommands()
		case "4":
			monitor.continuousMonitoring()
		case "5":
			fmt.Println("👋 Até logo!")
			return
		default:
			fmt.Println("❌ Opção inválida. Tente novamente.")
		}

		fmt.Println("\n" + strings.Repeat("=", 50))
	}
}
