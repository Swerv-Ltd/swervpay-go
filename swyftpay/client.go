package swyftpay

import "fmt"

func (c *SwyftpayClient) PrintOptions() {
	fmt.Printf("Swyftpay Client Options: %+v\n", c.options)
}
