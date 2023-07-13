package model

type Item struct {
	id *int64
	quantity int32
	itemNumber string
	description string
	expiryDate string
	lotNumber string
}


const (
	DateLayout : "2006-01-02"
)
func NewItem(quantity int32, itemNumber string, itemDescription string, expiryDate string, lotNumber string) (*Item,error) {
	err:= validateItemFields(quantity,itemNumber,itemDescription)
	if err!=nil{
		return nil,err
	}

	return &Item{quantity: quantity, itemNumber: itemNumber, itemDescription: itemDescription, lotNumber: lotNumber,expiryDate: expiryDate},nil
}

func validateItemFields(quantity int32, itemNumber string, itemDescription string) error {
	if strings.TrimSpace(itemNumber) == "" || strings.TrimSpace(itemDescription) == "" || quantity < 1 {
		return fmt.errorf("Validation error for new Item with missing fields, itemNumber: %v, itemDescription: %v, quantity: %v",itemNumber,itemDescription,quantity)
	}
	return nil
}


func (i *Item) SetID(id *int64){
	i.id = id
}

func (i *Item) SetQuantity(quantity int32){
	i.quantity = quantity
}

func (i *Item) Quantity() int32 {
	return i.quantity
}

func (i *Item) isValidDate() error {
	_, err := time.Parse(DateLayout, i.expiryDate)
	if err != nil {
		fmt.Println("Invalid date format:", err)
	}
	return nil
}