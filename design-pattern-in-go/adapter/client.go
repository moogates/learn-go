package main

import "fmt"

type Client struct {
}

func (c *Client) insertLightningConnectorIntoComputer(com Computer) {
    // Lightning接口：苹果高速多功能I/O接口，是苹果公司制作的专属连接器规格。
    // 目前此连接器针脚为8pin，作为替代沿用已久的30pin连接器的新Lightning接
    // 口，其尺寸明显缩小，支持正反盲插。
    fmt.Println("Client inserts Lightning connector into computer.")
    com.insertIntoLightningPort()
}
