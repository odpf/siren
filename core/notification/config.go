package notification

import (
	"time"

	"github.com/goto/siren/plugins/queues"
)

type Config struct {
	MaxNumReceiverSelectors int           `mapstructure:"max_num_receiver_selectors" yaml:"max_num_receiver_selectors" default:"10"`
	MaxMessagesReceiverFlow int           `mapstructure:"max_messages_receiver_flow" yaml:"max_messages_receiver_flow" default:"10"`
	Queue                   queues.Config `mapstructure:"queue" yaml:"queue"`
	MessageHandler          HandlerConfig `mapstructure:"message_handler" yaml:"message_handler"`
	DLQHandler              HandlerConfig `mapstructure:"dlq_handler" yaml:"dlq_handler"`
	GroupBy                 []string      `mapstructure:"group_by" yaml:"group_by"`

	EnableSilenceFeature bool
}

type HandlerConfig struct {
	Enabled       bool          `mapstructure:"enabled" yaml:"enabled" default:"true"`
	PollDuration  time.Duration `mapstructure:"poll_duration" yaml:"poll_duration" default:"5s"`
	ReceiverTypes []string      `mapstructure:"receiver_types" yaml:"receiver_types"`
	BatchSize     int           `mapstructure:"batch_size" yaml:"batch_size" default:"1"`
}
