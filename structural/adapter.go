package structural

func TestAdapter() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		machine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	println("insert to mac")
}

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	println("insert to usb port")
}

type WindowsAdapter struct {
	machine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	println("insert windows adapter to lightning port")
	w.machine.InsertIntoUSBPort()
}
