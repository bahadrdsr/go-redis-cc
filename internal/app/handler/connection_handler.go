package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn, commandHandler *CommandHandler) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Println("Received command:", cmd)

		response := commandHandler.HandleCommand(strings.TrimSpace(cmd))

		_, err := conn.Write([]byte(response + "\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
}
