package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"github.com/r523/joorvajoor/internal/rfid"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/rpi"
)

const (
	NArgs     = 2
	MaxAge    = 3600
	Gain      = 5
	AllowedID = "0cdb074999"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Joor", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("va", pterm.NewStyle(pterm.FgGreen)),
		pterm.NewLettersFromStringWithStyle("Joor", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	if len(os.Args) != NArgs {
		pterm.Error.Printf("joorvajoor <player_host>")

		return
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		pterm.Error.Printf("cannot connect to tcp server %s", err)

		return
	}
	conn.SetDeadline(time.Now().Add(time.Hour))

	app := fiber.New()
	api := app.Group("/api")

	// nolint: wrapcheck
	api.Get("/play", func(c *fiber.Ctx) error {
		_, err := conn.Write([]byte("play\n"))
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(http.StatusOK)
	})

	// nolint: wrapcheck
	api.Get("/pause", func(c *fiber.Ctx) error {
		_, err := conn.Write([]byte("pause\n"))
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(http.StatusOK)
	})

	// nolint: wrapcheck
	api.Get("/volume-up", func(c *fiber.Ctx) error {
		_, err := conn.Write([]byte("volume-up\n"))
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(http.StatusOK)
	})

	// nolint: wrapcheck
	api.Get("/volume-down", func(c *fiber.Ctx) error {
		_, err := conn.Write([]byte("volume-down\n"))
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(http.StatusOK)
	})

	go func() {
		var (
			ResetPin gpio.PinOut = rpi.P1_13
			IRQPin   gpio.PinIn  = rpi.P1_11
		)

		rid, err := rfid.Setup("/dev/spidev0.0", ResetPin, IRQPin, Gain)
		if err != nil {
			pterm.Error.Printf("cannot create rfid device %s\n", err)

			return
		}

		pterm.Info.Println("Started rfid reader.")

		for {
			id := rfid.ReadRFIDWithRetries(rid)

			pterm.Info.Println(id)

			if id != AllowedID {
				pterm.Error.Printf("you cannot have access %s\n", id)

				continue
			}

		}
	}()

	app.Static("/", "web/joorvajoor/out", fiber.Static{
		Compress:      true,
		ByteRange:     false,
		Browse:        false,
		Index:         "index.html",
		CacheDuration: time.Hour,
		MaxAge:        MaxAge,
		Next:          nil,
	})

	if err := app.Listen(":1378"); !errors.Is(err, http.ErrServerClosed) {
		pterm.Error.Printf("server start failed %s", err)
	}
}
