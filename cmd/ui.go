package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"go-rendang-cli/internal/rendang"
)

const (
	ansiReset  = "\033[0m"
	ansiBold   = "\033[1m"
	ansiDim    = "\033[2m"
	ansiRed    = "\033[31m"
	ansiGreen  = "\033[32m"
	ansiYellow = "\033[33m"
	ansiCyan   = "\033[36m"
)

func supportsColor(out io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	file, ok := out.(*os.File)
	if !ok {
		return false
	}

	info, err := file.Stat()
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}

func style(out io.Writer, text string, codes ...string) string {
	if !supportsColor(out) {
		return text
	}

	return strings.Join(codes, "") + text + ansiReset
}

func stockStyle(out io.Writer, remaining uint, total uint) string {
	if remaining == 0 {
		return style(out, fmt.Sprintf("%d/%d pack(s) available", remaining, total), ansiBold, ansiRed)
	}

	if remaining*5 <= total {
		return style(out, fmt.Sprintf("%d/%d pack(s) available", remaining, total), ansiBold, ansiYellow)
	}

	return style(out, fmt.Sprintf("%d/%d pack(s) available", remaining, total), ansiBold, ansiGreen)
}

func label(out io.Writer, text string) string {
	return style(out, text, ansiDim)
}

func printBanner(out io.Writer, summary rendang.Summary) {
	title := fmt.Sprintf("%s CLI", summary.EventName)
	line := strings.Repeat("=", len(title))

	fmt.Fprintln(out, style(out, line, ansiCyan))
	fmt.Fprintln(out, style(out, title, ansiBold, ansiCyan))
	fmt.Fprintln(out, style(out, line, ansiCyan))
	fmt.Fprintf(out, "%s %s\n", label(out, "Stock:"), stockStyle(out, summary.RemainingStock, summary.TotalStock))
	fmt.Fprintln(out)
}

func printOrderReceipt(out io.Writer, order rendang.Order, summary rendang.Summary) {
	fmt.Fprintln(out, style(out, "Order confirmed", ansiBold, ansiGreen))
	fmt.Fprintln(out, style(out, "---------------", ansiGreen))
	fmt.Fprintf(out, "%s %s %s\n", label(out, "Customer :"), order.FirstName, order.LastName)
	fmt.Fprintf(out, "%s %s\n", label(out, "Email    :"), order.Email)
	fmt.Fprintf(out, "%s %d\n", label(out, "Packs    :"), order.Packs)
	fmt.Fprintf(out, "%s %s\n", label(out, "Remaining:"), stockStyle(out, summary.RemainingStock, summary.TotalStock))
	fmt.Fprintln(out)
	fmt.Fprintf(out, "%s %s.\n\n", style(out, "Confirmation queued for", ansiCyan), order.Email)
}

func printStatus(out io.Writer, summary rendang.Summary) {
	printBanner(out, summary)
	fmt.Fprintf(out, "%s %d\n", label(out, "Orders booked:"), summary.OrderCount)

	if summary.OrderCount == 0 {
		fmt.Fprintf(out, "%s %s\n", label(out, "Bookings     :"), style(out, "none yet", ansiYellow))
		fmt.Fprintln(out)
		fmt.Fprintln(out, style(out, "Try `rendang order` for a one-off order or `rendang session` for the interactive flow.", ansiCyan))
		return
	}

	names := make([]string, 0, len(summary.Bookings))
	for _, booking := range summary.Bookings {
		names = append(names, fmt.Sprintf("%s %s", booking.FirstName, booking.LastName))
	}

	fmt.Fprintf(out, "%s %s\n", label(out, "Bookings     :"), strings.Join(names, ", "))
}

func printSessionSummary(out io.Writer, summary rendang.Summary) {
	fmt.Fprintln(out)
	fmt.Fprintln(out, style(out, "Session summary", ansiBold, ansiCyan))
	fmt.Fprintln(out, style(out, "---------------", ansiCyan))
	fmt.Fprintf(out, "%s %d\n", label(out, "Total orders :"), summary.OrderCount)
	fmt.Fprintf(out, "%s %s\n", label(out, "Stock left   :"), stockStyle(out, summary.RemainingStock, summary.TotalStock))
}
