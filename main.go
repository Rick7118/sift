package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/Rick7118/sift/database"
)

func main() {
	db, err := database.Open("sift.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Sift SQL")
	fmt.Println("Type SQL and end your query with ';'")
	fmt.Println("Type '.help' for available commands.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	var queryBuilder strings.Builder

	// Handle Ctrl+C.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signalChan)

	go func() {
		for range signalChan {
			if queryBuilder.Len() > 0 {
				queryBuilder.Reset()
				fmt.Println("\nQuery cancelled.")
				fmt.Print("sift> ")
			}
		}
	}()

	for {
		if queryBuilder.Len() == 0 {
			fmt.Print("sift> ")
		} else {
			fmt.Print("...> ")
		}

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// Allow .cancel even while building a multiline query.
		if strings.EqualFold(line, ".cancel") {
			queryBuilder.Reset()
			fmt.Println("Query cancelled.")
			fmt.Println()
			continue
		}

		// Handle Sift commands when we're not building a SQL query.
		if queryBuilder.Len() == 0 && strings.HasPrefix(line, ".") {
			if handleCommand(db, line) {
				break
			}

			continue
		}

		queryBuilder.WriteString(line)
		queryBuilder.WriteString(" ")

		// Wait for the semicolon before executing SQL.
		if !strings.HasSuffix(line, ";") {
			continue
		}

		query := strings.TrimSpace(queryBuilder.String())
		queryBuilder.Reset()

		result, err := db.Execute(query)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println()
			continue
		}

		if result == nil {
			continue
		}

		if len(result.Columns) > 0 {
			printResult(result)
		} else {
			fmt.Printf(
				"Query executed successfully. %d row(s) affected.\n\n",
				result.RowsAffected,
			)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func handleCommand(db *database.Database, command string) bool {
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return false
	}

	switch strings.ToLower(parts[0]) {
	case ".exit", ".quit":
		return true

	case ".help":
		printHelp()

	case ".tables":
		printTables(db)

	case ".schema":
		if len(parts) < 2 {
			fmt.Println("Usage: .schema <table>")
			fmt.Println()
			return false
		}

		printSchema(db, parts[1])

	case ".clear":
		clearScreen()

	case ".cancel":
		// Normally handled before reaching here.
		fmt.Println("No query to cancel.")
		fmt.Println()

	default:
		fmt.Printf("Unknown command: %s\n", parts[0])
		fmt.Println("Type '.help' for available commands.")
		fmt.Println()
	}

	return false
}

func printHelp() {
	fmt.Println()
	fmt.Println("Sift commands:")
	fmt.Println()
	fmt.Println("  .tables           List all tables")
	fmt.Println("  .schema <table>   Show table structure")
	fmt.Println("  .clear            Clear the terminal")
	fmt.Println("  .cancel           Cancel current SQL query")
	fmt.Println("  .help             Show this help message")
	fmt.Println("  .exit             Exit Sift")
	fmt.Println()
}

func printTables(db *database.Database) {
	tables, err := db.Tables()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()

	if len(tables) == 0 {
		fmt.Println("(no tables)")
		fmt.Println()
		return
	}

	for _, table := range tables {
		fmt.Println(table)
	}

	fmt.Println()
}

func printSchema(db *database.Database, table string) {
	columns, err := db.Schema(table)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()

	if len(columns) == 0 {
		fmt.Printf("Table '%s' not found.\n\n", table)
		return
	}

	fmt.Printf(
		"%-20s %-20s %-10s %-12s\n",
		"COLUMN",
		"TYPE",
		"NOT NULL",
		"PRIMARY KEY",
	)

	fmt.Println(strings.Repeat("-", 66))

	for _, column := range columns {
		fmt.Printf(
			"%-20s %-20s %-10v %-12v\n",
			column.Name,
			column.Type,
			column.NotNull,
			column.PrimaryKey,
		)
	}

	fmt.Println()
}

func clearScreen() {
	var command *exec.Cmd

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		command = exec.Command("clear")
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Print("\033[H\033[2J")
	}
}

func printResult(result *database.QueryResult) {
	fmt.Println()

	for _, column := range result.Columns {
		fmt.Printf("%-20s", column)
	}

	fmt.Println()

	for range result.Columns {
		fmt.Printf("%-20s", "--------------------")
	}

	fmt.Println()

	for _, row := range result.Rows {
		for _, value := range row {
			fmt.Printf("%-20v", value)
		}

		fmt.Println()
	}

	fmt.Printf("\n%d rows\n\n", len(result.Rows))
}
