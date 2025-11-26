package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func WithSSE(h *Handler, fn func(h *Handler, c echo.Context, send chan<- any)) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := c.Response()
		req := c.Request()

		// SSE headers
		res.Header().Set(echo.HeaderContentType, "text/event-stream")
		res.Header().Set("Cache-Control", "no-cache")
		res.Header().Set("Connection", "keep-alive")
		res.WriteHeader(http.StatusOK)
		res.Flush()

		// channel for sending SSE messages
		send := make(chan any)

		// Writer goroutine
		go func() {
			for {
				select {
				case msg, ok := <-send:
					if !ok {
						return
					}

					data, err := json.Marshal(msg)
					if err != nil {
						return
					}

					// Write event
					if _, err := res.Write([]byte("data: " + string(data) + "\n\n")); err != nil {
						return
					}
					// Flush buffer
					res.Flush()

				case <-req.Context().Done():
					// client disconnected
					return
				}
			}
		}()

		// Run the actual SSE provider function
		fn(h, c, send)

		// Close sending channel after fn returns
		close(send)
		return nil
	}
}
