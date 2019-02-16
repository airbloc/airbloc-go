package main

const (
	name      = "airbloc"
	descShort = "airbloc is user-oriented data exchange protocol"
	descLong  = `airbloc is a daemon of Airbloc Protocol, user-oriented and decentralized data exchange protocol,
    which collects and exchanges data based on users' consent.
    For details, please see the documentation at http://docs.airbloc.org`
)

// this build informations will be overriden by LDFLAGS. Do not edit.
var (
	Version   = "0.0.0"
	GitCommit = ""
	GitBranch = ""
	BuildDate = ""
)
