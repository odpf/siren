package alert

type Config struct {
	GroupBy []string `mapstructure:"group_by" yaml:"group_by"`
}
