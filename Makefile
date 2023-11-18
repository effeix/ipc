.DEFAULT_GOAL := help

##@ 🚀 Getting Started

build: ##@ This builds the binary
	go build -o ipc

clean: ##@ Remove generated files
	rm -f ipc

##@ 🔧 Development
.PHONY: cmd
cmd: ##@ Generates boilerplate for a new command (name=<command name>)
	@echo "package ipc\n\nimport (\n    \"github.com/spf13/cobra\"\n)\n\nvar $(name)Cmd = &cobra.Command{\n    Use:   \"$(name)\",\n    Short: \"Short description.\",\n    Args:  cobra.ExactArgs(1),\n    Run:   func(cmd *cobra.Command, args []string) {},\n}\n\nfunc init() {\n    rootCmd.AddCommand($(name)Cmd)\n}" > cmd/ipc/$(name).go

##@ ℹ️  Help
.PHONY: help
help: ##@ Displays this help message
	@echo "Usage: make [\033[1;35mtarget\033[0m]"
	@echo ""
	@awk -F ': |##@ ' \
		'/^##@/{if (section != "") print ""; section = $$2; printf("  \033[1;37m%s\n\033[0m", section); next} \
		/^[a-zA-Z0-9_-]+:/ {target = $$1; sub(":", "", target); printf("      \033[1;35m%-20s\033[0m", target); if ($$0 ~ /##@/) {sub(/^[^:]+: ##@ /, "", $$0); printf(" %s", $$0)} printf("\n")}' \
		$(MAKEFILE_LIST)
	@echo ""
