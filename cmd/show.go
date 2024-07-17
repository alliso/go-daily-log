/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-daily-log/application"
	"go-daily-log/cmd/helpers"
	"go-daily-log/domain/log"
	log2 "log"
	"strconv"
	"time"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log2.Fatal("Error using command show")
		}
	},
}

// showDayCmd represents the show day command
var showDayCmd = &cobra.Command{
	Use:   "day",
	Short: "Show day log",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log2.Fatal("You must specify only a day to show")
		}

		var date time.Time
		switch strDate := args[0]; strDate {
		case "today":
			date = time.Now()
		case "last", "ld":
			date = application.LastDay().Get()
		default:
			date = helpers.ParseDate(strDate)
		}
		entries := application.ShowDay().Apply(date)

		printTable(entries)
	},
}

// showMonthCmd represents the show day command
var showMonthCmd = &cobra.Command{
	Use:   "month",
	Short: "Show month log",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log2.Fatal("You must specify only a year and a month to show")
		}
		var entries []log.Entry
		entries = application.ShowMonth().Apply(getYearAndMonth(args[0], args[1]))

		printTable(entries)
	},
}

// showWeekCmd represents the show day command
var showWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show week log",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log2.Fatal("You must specify only a week")
		}
		if args[0] != "current" {
			log2.Fatal("Only current week is supported")
		}
		var entries []log.Entry
		entries = application.ShowWeek().Apply(time.Now())

		printTable(entries)
	},
}

func getYearAndMonth(strYear string, strMonth string) (int, int) {
	var err error
	var year int
	var month int
	year, err = strconv.Atoi(strYear)
	month, err = strconv.Atoi(strMonth)
	if err != nil {
		log2.Fatal("Invalid year, month")
	}

	return year, month
}

func printTable(entries []log.Entry) {
	hours := 0.0
	tbl := table.New("Task", "Project", "Date", "Hours")
	for _, entry := range entries {
		tbl.AddRow(entry.Task, entry.Project, entry.Date.Format(time.DateOnly), entry.Hours)
		hours += entry.Hours
	}
	tbl.AddRow()
	tbl.AddRow("Total hours", "", "", hours)

	tbl.Print()
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.AddCommand(showDayCmd)
	showCmd.AddCommand(showMonthCmd)
	showCmd.AddCommand(showWeekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
