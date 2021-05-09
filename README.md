# accshm
Go implementation for reading Assetto Corsa Competizione Shared Memory Pages

This module obviously only works on Windows, as ACC only runs on Windows as well.

## Basic Usage

Import the module like this:

```go
import "github.com/Dekadee/accshm"
```

Read the memory page like this:

```go
physics := new(accshm.ACCPhysics)

// Reading will fail, if the game has not been started at least once
err := ReadPhysics(physics)
if err != nil {
    // Handle potential errors
}

if prevPacketID != physics.PacketId {
    // Do something with the new data
}

// Analog for Graphics and Static memory page
```

At the moment the function creates a buffer everytime it is called.
This may be very inefficient, but realistically reading the memory pages is very fast and has almost no noticeable performance impact.

## Eventing

Instead of reading the memory pages yourself, you can subscribe to events like new lap times, sector times etc.
You can add handlers to events like laptime updates etc.

````go
eventPublisher := accshm.NewEventPublisher()

err := eventPublisher.AddLaptimeSubscription("printer", func(update LapTimeUpdate) {
    fmt.Printf("Last Time: %q", update.LastTime)
})

// Starts polling the shared memory every x seconds
eventPublisher.Start(1)

...

// Stop polling the pages
eventPublisher.Stop()
````


