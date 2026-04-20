package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed skills/amc/SKILL.md
var skillContent []byte

var (
	skillsInstallPath string
	installToClaude   bool
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage AtomGit CLI skills",
	Long:  "Manage skill extensions for AtomGit CLI.",
}

var skillsInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install skill",
	Long: `Install the built-in skill to AtomGit CLI.

The skill includes usage guide for repo, issue, pr, and skills commands.

Default installation path: ~/.agents/skills
Use --claude to install to ~/.claude/skills (Claude Code skills directory)
Use --path to specify a different directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		installPath := resolveInstallPath()

		if err := extractSkill(installPath); err != nil {
			return fmt.Errorf("failed to install skill: %w", err)
		}

		fmt.Printf("Installed skill to: %s\n", installPath)
		return nil
	},
}

func resolveInstallPath() string {
	if installToClaude {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ".claude/skills"
		}
		return filepath.Join(homeDir, ".claude", "skills")
	}

	if skillsInstallPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ".agents/skills"
		}
		return filepath.Join(homeDir, ".agents", "skills")
	}

	if skillsInstallPath == "." {
		cwd, _ := os.Getwd()
		return filepath.Join(cwd, "skills")
	}

	if !filepath.IsAbs(skillsInstallPath) {
		cwd, _ := os.Getwd()
		return filepath.Join(cwd, skillsInstallPath, "skills")
	}

	return filepath.Join(skillsInstallPath, "skills")
}

func extractSkill(installPath string) error {
	skillDir := filepath.Join(installPath, "amc")
	if err := os.MkdirAll(skillDir, 0755); err != nil {
		return err
	}

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, skillContent, 0644); err != nil {
		return err
	}

	return nil
}

var skillListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available skills",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Available skills:")
		fmt.Println("  amc    AtomGit CLI 使用指南 (repo, issue, pr)")
		fmt.Println()
		fmt.Println("Install with: amc skills install")
		return nil
	},
}

func init() {
	skillsInstallCmd.Flags().StringVarP(&skillsInstallPath, "path", "p", "", "Installation path (default: ~/.agents/skills, use '.' for ./skills)")
	skillsInstallCmd.Flags().BoolVarP(&installToClaude, "claude", "c", false, "Install to Claude Code skills directory (~/.claude/skills)")

	skillsCmd.AddCommand(skillsInstallCmd)
	skillsCmd.AddCommand(skillListCmd)
	rootCmd.AddCommand(skillsCmd)
}
