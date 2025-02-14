// Copyright © 2024 JR team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/jrnd-io/jr/pkg/constants"
	"github.com/jrnd-io/jr/pkg/producers/kafka"
	"github.com/spf13/cobra"
)

var createTopicCmd = &cobra.Command{
	Use:   "createTopic [topic]",
	Short: "Create a Kafka Topic",
	Long: `Create a Kafka Topic.
jr createTopic advancedTopic -p 10 -r 3
jr createTopic newDefaultTopic
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kafkaConfig, _ := cmd.Flags().GetString("kafkaConfig")

		kManager := &kafka.KafkaManager{}
		kManager.Initialize(kafkaConfig)
		partitions, _ := cmd.Flags().GetInt("partitions")
		replica, _ := cmd.Flags().GetInt("replica")
		kManager.CreateTopicFull(args[0], partitions, replica)

	},
}

func init() {
	rootCmd.AddCommand(createTopicCmd)
	createTopicCmd.Flags().IntP("partitions", "p", constants.DEFAULT_PARTITIONS, "Number of partitions")
	createTopicCmd.Flags().IntP("replica", "r", constants.DEFAULT_REPLICA, "Replica Factor")
	createTopicCmd.Flags().StringP("kafkaConfig", "F", constants.KAFKA_CONFIG, "Kafka configuration")
}
