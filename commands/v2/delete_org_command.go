package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type DeleteOrgCommand struct {
	RequiredArgs flags.Organization `positional-args:"yes"`
	Force        bool               `short:"f" description:"Force deletion without confirmation"`
	usage        interface{}        `usage:"CF_NAME delete-org ORG [-f]"`
}

func (_ DeleteOrgCommand) Setup(config commands.Config) error {
	return nil
}

func (_ DeleteOrgCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}