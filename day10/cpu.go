package day10

type cpu struct {
	cycle    int
	x        int
	xHistory []int
}

func NewCPU() *cpu {
	return &cpu{
		cycle:    0,
		x:        1,
		xHistory: make([]int, 0),
	}
}

func (c *cpu) noop() {
	c.tick()
	return
}

func (c *cpu) addX(n int) {
	c.tick()
	c.tick()
	c.x += n
}

func (c *cpu) tick() {
	c.cycle++
	c.xHistory = append(c.xHistory, c.x)
}

func (c *cpu) ProcessCommand(cmd command) {
	switch cmd.(type) {
	case noopCommand:
		c.noop()
	case addXCommand:
		c.addX(cmd.(addXCommand).value)
	}
}

func (c *cpu) GetSignalStrength(clockCycle int) (int, bool) {

	value, ok := c.getValue(clockCycle)

	return value * clockCycle, ok
}

func (c *cpu) getValue(clockCycle int) (int, bool) {
	if clockCycle < 0 {
		return 0, false
	}

	if clockCycle-1 < 0 || clockCycle-1 >= len(c.xHistory) {
		return 0, false
	}

	return c.xHistory[clockCycle-1], true
}
