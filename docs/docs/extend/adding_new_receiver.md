# Add a New Receiver Plugin

Receiver plugin is being used to send notifications to the receiver. More details about the concept of receiver plugin can be found [here](../concepts/plugin.md#receiver-plugin).

## Steps to Add a New Plugin

1. **Add a new package**

   A new plugin could be added as a new package under `plugins/providers/{new_plugin}`. Package name should comply with raystack golang standard naming. See the [handbook](https://raystack.github.io/handbook/programming/go#packages). Ideally you might want to name the package with lower-case and without any `-` or `_` signs.

2. **Defining Configurations**

   - **Define `AppConfig`** (Optional)

     If there is a need to have a generic config for the receiver that is being loaded during start-up, you could add a new `AppConfig` and assign the config to `plugins/receivers/config.go` to expose it to the app-level config. Siren will recognize and read the config when starting up.

   - **Define `ReceiverConfig`**

     For each receiver, there must be a specific config that is important to be used to send notifications. You need to add the information needed in `ReceiverConfig`.

   - **Define `SubscriptionConfig`** (Optional)

     It is also possible for a receiver to have different config in the receiver and subscription. For example, slack has a dedicated config called `channel_name` in subscription to send notification only to a specific channel. In this case you need to define separate configurations: `ReceiverConfig` and `SubscriptionConfig`.

   - **Define `NotificationConfig`**

     You need to implement `NotificationConfig` which is a placeholder to combine `ReceiverConfig` and `SubscriptionConfig` (if any). Therefore `NotificationConfig` should just embed `ReceiverConfig` and `SubscriptionConfig` (if needed). You might also want to add more function to validate and transform the config to `map[string]interface{}`.

3. **Implement Interfaces**

   - **Implement PreHook methods** (Optional)

     You need to figure out whether there is a need to do pre-processing of receiver configuration before storing to the database or enqueueing to the queue. For some receivers, there is a need to do encryption or validation in pre-processing step, in this case you could implement `PreHookDBTransformConfigs` to transform and validate configurations before storing it to the DB and `PreHookQueueTransformConfigs` to transform and validate configurations before storing it to the queue.

   - **Implement PostHook methods** (Optional)

     If there is a need to transform back receiver's configurations (e.g. decrypting config), you need to implement `PostHookDBTransformConfigs` or `PostHookQueueTransformConfigs` to transform the config back for processing.

   - **Implement `Send` notification method**

     You need to implement `Send` method to send notification message to the external notification vendor.

   - **Implement `GetSystemDefaultTemplate` method**

     Add a new default template for alert notification in `plugins/receivers/{type}/config/{template_file_name}.goyaml` and return it inside the `GetSystemDefaultTemplate` method. More details about it are [here](../concepts/plugin.md#alert-notification-default-template).

   - **Implement `BuildData` method** (Optional)

     If Siren needs to return some additional post-processed data during Get and List Receivers (see [here](../concepts/plugin.md#receiver-plugin) for more context), you could implement `BuildData` method.

4. **Integrate New Plugin with Siren**

   - Define and add your new type of plugin inside `core/receivers/type.go`.
   - Initialize your plugin receiver service and add to the `ConfigResolver` and `Notifier` registries map in `cli/deps`.
   - To make sure notification handler and dlq handler process your new type, don't forget to add your new receiver type in notification message & dlq handler config or make it default to support all receiver types.

## Example: Add a new plugin `file` receiver

In this part, we will show how to add a new receiver plugin to send notifications to a local `file`.

### Defining Configurations

To write a file, we need a `url` of the file. This might be the only config that we needs. We also don't need to define `SubscriptionConfig` since we don't need a specific config for the subscription.

```go
type ReceiverConfig struct {
	URL string `mapstructure:"url"`
}
```

We define a `NotificationConfig` which only embeds `ReceiverConfig`. This is helpful to separate the concern for a specific use-cases in some plugins.

```go
type NotificationConfig struct {
	ReceiverConfig `mapstructure:",squash"`
}
```

For `file` type, we don't need an `AppConfig` as for now. So we don't need to add one in `plugins/receivers/config.go`.

Now that we already have defined all configs needed, we needs to implement all methods of interfaces needed by defining a new `PluginService`.

### Implement interfaces

We need to create a new `Plugin` and implement `ConfigResolver` and `Notifier`. For `file` receiver, we don't need to do encryption of configs before and after writing and reading from the DB as well as Queue. Therefore, we only needs to implement `PreHookDBTransformConfigs` to validate the config before storing it to the DB and `PreHookDBTransformConfigs` to validate the config before enqueueing it.

```go
// highlight-start

type PluginService struct {
	base.UnimplementedService
}

// NewPluginService returns file receiver service struct. This service implement [receiver.Resolver] and [notification.Notifier] interface.
func NewPluginService() *PluginService {
	return &PluginService{}
}

func (s *PluginService) PreHookDBTransformConfigs(ctx context.Context, receiverConfigMap map[string]interface{}) (map[string]interface{}, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(receiverConfigMap, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	if err := receiverConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	return receiverConfig.AsMap(), nil
}

func (s *PluginService) PreHookTransformConfigs(ctx context.Context, notificationConfigMap map[string]interface{}) (map[string]interface{}, error) {
	notificationConfig := &NotificationConfig{}
	if err := mapstructure.Decode(notificationConfigMap, notificationConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to slack notification config: %w", err)
	}

	if err := notificationConfig.Validate(); err != nil {
		return nil, err
	}

	return notificationConfig.AsMap(), nil
}
// highlight-end
```

Beside those 2 functions, we also need to add a function to send notifications (to write notifications to a file under `url`).

```go
type PluginService struct {
	base.UnimplementedService
}

func NewPluginService() *PluginService {
	return &PluginService{}
}

func (s *PluginService) PreHookDBTransformConfigs(ctx context.Context, receiverConfigMap map[string]interface{}) (map[string]interface{}, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(receiverConfigMap, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	if err := receiverConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	return receiverConfig.AsMap(), nil
}

func (s *PluginService) PreHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]interface{}) (map[string]interface{}, error) {
	notificationConfig := &NotificationConfig{}
	if err := mapstructure.Decode(notificationConfigMap, notificationConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to slack notification config: %w", err)
	}

	if err := notificationConfig.Validate(); err != nil {
		return nil, err
	}

	return notificationConfig.AsMap(), nil
}

// highlight-start
func (s *PluginService) Send(ctx context.Context, notificationMessage notification.Message) (bool, error) {
	notificationConfig := &NotificationConfig{}
	if err := mapstructure.Decode(notificationMessage.Configs, notificationConfig); err != nil {
		return false, err
	}

	bodyBytes, err := json.Marshal(notificationMessage.Details)
	if err != nil {
		return false, err
	}

	fileInstance, err := os.OpenFile(notificationConfig.URL, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return false, err
	}

	byteNewLine := []byte("\n")
	bodyBytes = append(bodyBytes, byteNewLine...)
	_, err = fileInstance.Write(bodyBytes)
	if err != nil {
		return false, err
	}

	return false, nil
}
// highlight-end
```
