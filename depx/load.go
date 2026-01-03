package depx

import (
	"chihqiang/depctl/flagx"
	"github.com/urfave/cli/v3"
)

func Load(cmd *cli.Command) *Config {
	return &Config{
		Dir:         cmd.String(flagx.FlagDir),
		Version:     cmd.String(flagx.FlagVersion),
		LocalCmd:    cmd.String(flagx.FlagLocalCmd),
		Include:     cmd.StringSlice(flagx.FlagInclude),
		Exclude:     cmd.StringSlice(flagx.FlagExclude),
		RemoteRepo:  cmd.String(flagx.FlagRemoteRepo),
		CurrentLink: cmd.String(flagx.FlagCurrentLink),
		HookPre:     cmd.String(flagx.FlagHookPre),
		HookPost:    cmd.String(flagx.FlagHookPost),
	}
}
