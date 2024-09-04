package main

import (
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v3"
)

func GetIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		// Check if the address is an IP address and not a loopback address
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// Check if it's an IPv4 address
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		hostname, err := os.Hostname()
		ip := GetIPAddress()
		if err != nil {
			return c.SendString("Hello, World 👋!")
		}
		return c.SendString(hostname + " -> " + ip)

	})

	// liveness
	app.Get("/healthy", func(c fiber.Ctx) error {
		isAlive := true 
        if isAlive {
            c.Status(fiber.StatusOK)
            return c.SendString("Liveness check passed!")
        } else {
            c.Status(fiber.StatusInternalServerError)
            return c.SendString("Liveness check failed!")
        }
	})

	// is ready to serve traffic ?
	// in this method we can check connections to database and messege brokers etc...
	app.Get("/ready", func(c fiber.Ctx) error {
		isReady := true 
        if isReady {
            c.Status(fiber.StatusOK)
            return c.SendString("Readiness check passed!")
        } else {
            c.Status(fiber.StatusServiceUnavailable)
            return c.SendString("Readiness check failed!")
        }
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
