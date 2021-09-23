/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"websocket-bench/client"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("server-addr")
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		if addr == "" {
			log.Error("server  addr is null")
			return
		}
		cli, closer, err := client.NewCommonRPCV0(cmd.Context(), addr, http.Header{})
		if err != nil {
			log.Error(err)
			return
		}
		defer closer()
		td, err := cli.GetSession(cmd.Context())
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(td)
		sigCh := make(chan os.Signal, 2)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		task, err := cli.WorkerQueue(cmd.Context(), td, 0)
		if err != nil {
			log.Error(err)
			return
		}
		aliveChecking := time.After(30 * time.Second)
	loop:
		for {
			select {
			case <-aliveChecking:
				version, err := cli.Version(cmd.Context())
				if version == "" || err != nil {
					log.Info("disconnection from server: ", td)
				}
			case <-cmd.Context().Done():
				break loop
			case taskid := <-task:
				if taskid == 0 {
					log.Info("server shutdown")
					for {
						time.Sleep(time.Second * 10)
						log.Info("try reconnect sever every 10s ")
						task, err = cli.WorkerQueue(cmd.Context(), td, taskid)
						if err == nil {
							goto loop
						}
					}
				}
				log.Infof("receive new task %v from server", taskid)
				go func() {
					time.Sleep(time.Second * 5)
					err := cli.WorkerDone(cmd.Context(), td, taskid)
					if err != nil {
						log.Error(err)
					}
					log.Infof("done  task %v ", taskid)
				}()
			case sig := <-sigCh:
				log.Infof("signal %s captured", sig)
				break loop
			}
		}
		log.Info("gracefull down ", td)
	},
}

func init() {
	clientCmd.Flags().String("server-addr", "ws://127.0.0.1:3500/rpc/v0", "websocket server address")
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
