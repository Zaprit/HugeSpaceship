package model

type Platform string

const (
	PS3     Platform = "PS3"
	PSP     Platform = "PSP"
	PSVita  Platform = "PSVita"
	PS4     Platform = "PS4" // For future expansion
	RPCS3   Platform = "RPCS3"
	Browser Platform = "Browser" // For the website
	// If suitable psp/psvita emulation comes to fruition then I will add more emulators here
)
