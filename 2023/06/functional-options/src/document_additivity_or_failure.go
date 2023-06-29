package fcors

// -snip-
func NewCORS(opts ...Option) (Middleware, error)

// FromOrigins lists the allowed origins.
// Using this option more than once in a call to NewCORS // HL
// results in a failure to build the corresponding middleware. // HL
func FromOrigins(origins ...string) Option
