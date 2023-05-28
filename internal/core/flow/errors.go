package flow

import "fmt"

var (
	ErrOutOfService           = fmt.Errorf("vending machine is out of service")
	ErrInvalidKey             = fmt.Errorf("invalid key")
	ErrNoItem                 = fmt.Errorf("no item")
	ErrInvalidSlot            = fmt.Errorf("invalid slot")
	ErrInvalidEvent           = fmt.Errorf("invalid event")
	ErrEmptyItem              = fmt.Errorf("empty item")
	ErrItemIsNotEnough        = fmt.Errorf("item is not enough")
	ErrPromptpayOutOfService  = fmt.Errorf("promptpay is out of service")
	ErrCreditCardOutOfService = fmt.Errorf("creditcard is out of service")
	ErrCancel                 = fmt.Errorf("cancel order")
	ErrMachineE0              = fmt.Errorf("machine has no item (E0)")
	ErrMachineE1              = fmt.Errorf("user don't grab item (E1)")
	ErrMachineE2              = fmt.Errorf("unknown error (E2)")
	ErrOpenGate               = fmt.Errorf("the machine gate is opened")
	ErrCloseGate              = fmt.Errorf("the machine gate is closed")
)
