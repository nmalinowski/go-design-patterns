package main

type sammysangAdapter struct {
	// DONE: add a field for the SammysangTV reference
	sstv *SammysangTV
}

func (ss *sammysangAdapter) turnOn() {
	// DONE
	ss.sstv.setOnState(true)
}

func (ss *sammysangAdapter) turnOff() {
	// DONE
	ss.sstv.setOnState(false)
}

func (ss *sammysangAdapter) volumeUp() int {
	// DONE
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) volumeDown() int {
	// DONE
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) channelUp() int {
	// DONE
	ch := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) channelDown() int {
	// DONE
	ch := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	// DONE
	ss.sstv.setChannel(ch)
}

// END EXAMPLE
