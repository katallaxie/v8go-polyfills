package fetch

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	// UserAgent is the user agent to use when making requests.
	UserAgent string
}

// Defaults ...
func Defaults() *Opts {
	return &Opts{
		UserAgent: "v8go-polyfills",
	}
}

type fetcher struct {
	opts *Opts
}

// New ...
func New(opts ...Opt) *fetcher {
	f := &fetcher{
		opts: Defaults(),
	}

	for _, opt := range opts {
		opt(f.opts)
	}

	return f
}
