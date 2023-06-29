origins := []string{"https://example.com"} // HL

cors, err := fcors.NewCORS(
	fcors.FromOrigins(origins...), // HL
)
if err != nil {
	// handle error
}

origins[0] = "https://gophercon.eu" // no effect on the middleware // HL
