package cobrax

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/whitekid/goxp/flags"
	"github.com/whitekid/goxp/log"
)

// Add add sub commands to parents
// [local,persistent]Flags add command flags
// NOTE: flag is not bind to viper. this only about to cobra.Command flags
func Add(parent, cmd *cobra.Command, persistentFlags, localFlags []flags.Flag) *cobra.Command {
	if parent != nil {
		parent.AddCommand(cmd)
	}

	addFlag(cmd.Flags(), localFlags)
	addFlag(cmd.PersistentFlags(), persistentFlags)

	return cmd
}

func addFlag(fs *pflag.FlagSet, flags []flags.Flag) {
	for _, f := range flags {
		switch v := f.DefaultValue.(type) {
		case int:
			fs.IntP(f.Name, f.Shorthand, v, f.Usage)
		case int16:
			fs.Int16P(f.Name, f.Shorthand, v, f.Usage)
		case int32:
			fs.Int32P(f.Name, f.Shorthand, v, f.Usage)
		case int64:
			fs.Int64P(f.Name, f.Shorthand, v, f.Usage)
		case uint:
			fs.UintP(f.Name, f.Shorthand, v, f.Usage)
		case uint16:
			fs.Uint16P(f.Name, f.Shorthand, v, f.Usage)
		case uint32:
			fs.Uint32P(f.Name, f.Shorthand, v, f.Usage)
		case uint64:
			fs.Uint64P(f.Name, f.Shorthand, v, f.Usage)
		case float32:
			fs.Float32P(f.Name, f.Shorthand, v, f.Usage)
		case []float32:
			fs.Float32SliceP(f.Name, f.Shorthand, v, f.Usage)
		case float64:
			fs.Float64P(f.Name, f.Shorthand, v, f.Usage)
		case []float64:
			fs.Float64SliceP(f.Name, f.Shorthand, v, f.Usage)
		case bool:
			fs.BoolP(f.Name, f.Shorthand, v, f.Usage)
		case []bool:
			fs.BoolSliceP(f.Name, f.Shorthand, v, f.Usage)
		case string:
			fs.StringP(f.Name, f.Shorthand, v, f.Usage)
		case []string:
			fs.StringSliceP(f.Name, f.Shorthand, v, f.Usage)
		case []byte:
			fs.BytesHexP(f.Name, f.Shorthand, v, f.Usage)
		case time.Duration:
			fs.DurationP(f.Name, f.Shorthand, v, f.Usage)
		default:
			log.Errorf("unsupported type %T", f.DefaultValue)
		}
	}
}

func Apply[T any](fn func(name string) (T, error), name string) T {
	r, err := fn(name)
	if err != nil {
		panic(err)
	}
	return r
}
