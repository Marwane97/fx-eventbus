# fx-eventbus

`fx-eventbus` is a Go project that demonstrates the use of the [Uber's fx](https://github.com/uber-go/fx) library for dependency injection along with the [EventBus](https://github.com/asaskevich/EventBus) library for event-driven programming.

## Overview

The project is structured to showcase the interaction between different components using events. It includes services, event definitions, subscribers, and a main application module to tie everything together.

### Components

- **Services**: Components that perform specific tasks.
    - [Manager](https://github.com/Marwane97/fx-eventbus/blob/master/service/manager/manager.go): Manages tasks and emits events based on the task status.
    - [Search](https://github.com/Marwane97/fx-eventbus/blob/master/service/service.go): Performs a search operation and manages the flow of the application.
    - [Wrapper](https://github.com/Marwane97/fx-eventbus/blob/master/service/wrapper/wrapper.go): A wrapper service that also emits events.

- **Events**: Definitions of different events that can be emitted and listened to.
    - [ManagerEvent](https://github.com/Marwane97/fx-eventbus/blob/master/events/manager_event.go), [ServiceEvent](https://github.com/Marwane97/fx-eventbus/blob/master/events/service_event.go), [WrapperEvent](https://github.com/Marwane97/fx-eventbus/blob/master/events/wrapper_event.go)

- **Subscribers**: Components that listen to specific events and perform actions when those events are emitted.
    - [Manager Subscribers](https://github.com/Marwane97/fx-eventbus/tree/master/subscriber): Subscribers for manager events.
    - [Service Subscribers](https://github.com/Marwane97/fx-eventbus/tree/master/subscriber): Subscribers for service events.
    - [Wrapper Subscribers](https://github.com/Marwane97/fx-eventbus/tree/master/subscriber): Subscribers for wrapper events.

- **Application Module**: The main application module that initializes and runs the application.
    - [App Module](https://github.com/Marwane97/fx-eventbus/blob/master/app/app.go)

## Getting Started

To run this project, make sure you have Go installed on your machine.

1. Clone the repository:
   ```sh
   git clone git@github.com:Marwane97/fx-eventbus.git
   cd fx-eventbus
   ```
2. Run the application:
   ```sh
   go run main.go
   ```
3. d 
    The application will start, and you should see logs indicating the flow of events and the actions performed by the subscribers.
