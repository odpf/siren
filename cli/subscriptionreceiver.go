package cli

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/goto/salt/cmdx"
	"github.com/goto/salt/printer"
	"github.com/goto/siren/core/subscriptionreceiver"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/spf13/cobra"
)

func subscriptionsReceiversCmd(cmdxConfig *cmdx.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "receiver",
		Aliases: []string{"subscription_receiver"},
		Short:   "Manage subscriptions receiver",
		Long: heredoc.Doc(`
			Work with subscriptions receiver.
			
			Add receiver to a subscription.
		`),
		Annotations: map[string]string{
			"group":  "core",
			"client": "true",
		},
	}

	cmd.AddCommand(
		addSubscriptionReceiverCmd(cmdxConfig),
		editSubscriptionReceiverCmd(cmdxConfig),
		removeSubscriptionReceiverCmd(cmdxConfig),
	)

	return cmd
}

func addSubscriptionReceiverCmd(cmdxConfig *cmdx.Config) *cobra.Command {
	var filePath string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a receiver to a subscription",
		Long: heredoc.Doc(`
			Add a receiver to a subscription.
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			ctx := cmd.Context()

			c, err := loadClientConfig(cmd, cmdxConfig)
			if err != nil {
				return err
			}

			var srRelation subscriptionreceiver.Relation
			if err := parseFile(filePath, &srRelation); err != nil {
				return err
			}

			client, cancel, err := createClient(ctx, c.Host)
			if err != nil {
				return err
			}
			defer cancel()

			res, err := client.AddSubscriptionReceiver(ctx, &sirenv1beta1.AddSubscriptionReceiverRequest{
				SubscriptionId: srRelation.SubscriptionID,
				ReceiverId:     srRelation.ReceiverID,
				Labels:         srRelation.Labels,
			})

			if err != nil {
				return err
			}

			spinner.Stop()
			printer.Successf("Receiver id %d is added to subscription id %d", res.GetReceiverId(), res.GetSubscriptionId())
			printer.Space()
			printer.SuccessIcon()

			return nil
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "path to the subscription config")
	cmd.MarkFlagRequired("file")

	return cmd
}

func editSubscriptionReceiverCmd(cmdxConfig *cmdx.Config) *cobra.Command {
	var (
		subscriptionID uint64
		receiverID     uint64
		filePath       string
	)

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a subscription receiver",
		Long: heredoc.Doc(`
			Edit an existing subscription receiver relation detail.
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			ctx := cmd.Context()

			c, err := loadClientConfig(cmd, cmdxConfig)
			if err != nil {
				return err
			}

			var srRelation subscriptionreceiver.Relation
			if err := parseFile(filePath, &srRelation); err != nil {
				return err
			}

			client, cancel, err := createClient(ctx, c.Host)
			if err != nil {
				return err
			}
			defer cancel()

			res, err := client.UpdateSubscriptionReceiver(ctx, &sirenv1beta1.UpdateSubscriptionReceiverRequest{
				SubscriptionId: srRelation.SubscriptionID,
				ReceiverId:     srRelation.ReceiverID,
				Labels:         srRelation.Labels,
			})
			if err != nil {
				return err
			}

			spinner.Stop()
			printer.Successf("Successfully updated receiver id %d detail of asubscription with id %d", res.GetReceiverId(), res.GetSubscriptionId())
			printer.Space()
			printer.SuccessIcon()

			return nil
		},
	}

	cmd.Flags().Uint64VarP(&subscriptionID, "subscription_id", "s", 0, "subscription id")
	cmd.MarkFlagRequired("subscription_id")
	cmd.Flags().Uint64VarP(&receiverID, "receiver_id", "r", 0, "receiver id")
	cmd.MarkFlagRequired("receiver_id")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the subscription config")
	cmd.MarkFlagRequired("file")

	return cmd
}

func removeSubscriptionReceiverCmd(cmdxConfig *cmdx.Config) *cobra.Command {
	var (
		subscriptionID uint64
		receiverID     uint64
	)
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove receiver of a subscription",
		Example: heredoc.Doc(`
			$ siren subscription receiver remove --subscription_id 1 --receiver_id 2
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			ctx := cmd.Context()

			c, err := loadClientConfig(cmd, cmdxConfig)
			if err != nil {
				return err
			}

			client, cancel, err := createClient(ctx, c.Host)
			if err != nil {
				return err
			}
			defer cancel()

			_, err = client.DeleteSubscriptionReceiver(ctx, &sirenv1beta1.DeleteSubscriptionReceiverRequest{
				SubscriptionId: subscriptionID,
				ReceiverId:     receiverID,
			})
			if err != nil {
				return err
			}

			spinner.Stop()
			printer.Successf("Successfully removed receiver id %d from subscription id %d", receiverID, subscriptionID)
			printer.Space()
			printer.SuccessIcon()

			return nil
		},
	}

	cmd.Flags().Uint64VarP(&subscriptionID, "subscription_id", "s", 0, "subscription id")
	cmd.MarkFlagRequired("subscription_id")
	cmd.Flags().Uint64VarP(&receiverID, "receiver_id", "r", 0, "receiver id")
	cmd.MarkFlagRequired("receiver_id")
	return cmd
}
