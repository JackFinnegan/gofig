package gofig

// Gofig default configuration.
const (
	DefaultStructTag = "gofig"
)

// A Parser parses configuration.
type Parser interface {
	Parse() error
}

// A Notifier notifies via a channel if changes to configuration have occurred.
// Remember to check the error on the channel.
type Notifier interface {
	Notify() <-chan error
}

// A ParseNotifier can parse config and notify on changes to configuration.
type ParseNotifier interface {
	Parser
	Notifier
}
