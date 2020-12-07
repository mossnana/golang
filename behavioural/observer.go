package behavioural

func TestObserver() {
	shirtItem := newItem("Nike Shirt")
	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateStock()
}

// ตั้งค่า struct ที่จะเป็น observer ว่าต้องมี method เหล่านี้
type observer interface {
	update(string)
	getID() string
}

// customer นี้จะเป็น observer จึงต้องมี method update และ getID
type customer struct {
	id string
}

func (cus *customer) update(itemName string) {
	println("Sending email to", itemName)
}

func (cus *customer) getID() string {
	return cus.id
}

// item เป็นตัวรับ observer
type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) updateStock() {
	i.inStock = true
	i.notifyAll()
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
