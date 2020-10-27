/*
Copyright © 2020 Muhammad Muhlas Abror <muhammadmuhlas3@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"todo_app/internal/app/task"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete task",
	Long:  `delete single task or all task`,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args {
			s, _ := strconv.Atoi(v)
			ids = append(ids, s)
		}
		all, _ := cmd.Flags().GetBool("all")

		if all {
			task.DeleteAllTasks()
		} else {
			task.DeleteTask(ids)
		}

		task.ListTasks(true)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().BoolP("all", "a", false, "Mark all task as done")
}
