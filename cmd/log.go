/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-daily-log/application"
	"go-daily-log/cmd/bubble"
	"go-daily-log/cmd/bubble/text"
	log2 "go-daily-log/domain/log"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Add log to day",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Wrong number of args for log command")
		}
		strDate := args[0]
		var date time.Time
		var err error
		switch {
		case strDate == "today":
			date = time.Now()
		default:
			date, err = time.Parse(time.DateOnly, strDate)
			if err != nil {
				log.Fatal("Unable to parse date")
			}
		}

		project := bubble.GetProject()
		task := text.GetTask()
		hours := text.GetHours()
		application.LogDay().Apply(log2.Entry{
			Date:    date,
			Project: project,
			Task:    task,
			Hours:   hours,
		})
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
