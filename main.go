package main

import (
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

const NArgs = 2

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

	app := fiber.New()

	// nolint: wrapcheck
	app.Get("/play", func(c *fiber.Ctx) error {
		_, err := conn.Write([]byte("play"))
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(http.StatusOK)
	})

	if err := app.Listen(":1378"); !errors.Is(err, http.ErrServerClosed) {
		pterm.Error.Printf("server start failed %s", err)
	}
}
