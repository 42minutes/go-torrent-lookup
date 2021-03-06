package torrentlookup

// ProviderTPB -
var ProviderTPB = &Provider{
	Name:           "The pirate bay",
	SearchURL:      "https://thepiratebay.org/search/%s/0/99/0",
	RowQuery:       "#searchResult tr",
	NameSubQuery:   ".detName a",
	MagnetSubQuery: "td:nth-child(2) > a:nth-child(2)",
	SeedsSubQuery:  "td:nth-child(3)",
}

// ProviderTorzeu -
var ProviderTorzeu = &Provider{
	Name:           "Torrentz2.eu",
	SearchURL:      "https://torrentz2.eu/verified?f=%s",
	RowQuery:       "dl",
	NameSubQuery:   "dt>a",
	MagnetSubQuery: "dt>a",
	SeedsSubQuery:  "dd>span:nth-child(4)",
}
