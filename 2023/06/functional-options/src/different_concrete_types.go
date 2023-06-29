package main

type option func(*config) error

func (o option) applyAnon(cfg *config) error {
	return o(cfg)
}

func (o option) applyCred(cfg *config) error {
	return o(cfg)
}

func main() {
	fcors.AllowAccessWithCredentials(
		fcors.OptionCred(fcors.FromAnyOrigin()), // 😱
	)
}
