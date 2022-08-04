package traceroute

// Options for the traceroute command.
type Options struct {
	port       int
	maxHops    int
	firstHop   int
	timeoutMs  int
	retries    int
	packetSize int
}

func (options *Options) Port() int {
	if options.port == 0 {
		options.port = DefaultPort
	}
	return options.port
}

func (options *Options) SetPort(port int) {
	options.port = port
}

func (options *Options) MaxHops() int {
	if options.maxHops == 0 {
		options.maxHops = DefaultMaxHops
	}
	return options.maxHops
}

func (options *Options) SetMaxHops(maxHops int) {
	options.maxHops = maxHops
}

func (options *Options) FirstHop() int {
	if options.firstHop == 0 {
		options.firstHop = DefaultFirstHop
	}
	return options.firstHop
}

func (options *Options) SetFirstHop(firstHop int) {
	options.firstHop = firstHop
}

func (options *Options) TimeoutMs() int {
	if options.timeoutMs == 0 {
		options.timeoutMs = DefaultTimeoutMs
	}
	return options.timeoutMs
}

func (options *Options) SetTimeoutMs(timeoutMs int) {
	options.timeoutMs = timeoutMs
}

func (options *Options) Retries() int {
	if options.retries == 0 {
		options.retries = DefaultRetries
	}
	return options.retries
}

func (options *Options) SetRetries(retries int) {
	options.retries = retries
}

func (options *Options) PacketSize() int {
	if options.packetSize == 0 {
		options.packetSize = DefaultPacketSize
	}
	return options.packetSize
}

func (options *Options) SetPacketSize(packetSize int) {
	options.packetSize = packetSize
}
