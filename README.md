# RustMQ (Simplified Messaging Queue)

> Portfolio project

## Plan

Core Features:
1. Publisher-Subscriber Model
2. In-Memory Message Storage
3. Basic Persistence
4. Multiple Topic Support
5. Simple Client Library

Architecture Components:
```rust
struct Broker {
    topics: HashMap<String, Vec<Message>>,
    subscribers: HashMap<String, Vec<Subscriber>>
}

struct Message {
    id: UUID,
    topic: String,
    payload: Vec<u8>,
    timestamp: SystemTime
}

struct Subscriber {
    id: UUID,
    channel: mpsc::Sender<Message>
}
```

Planned Implementation Stages:

Stage 1: Basic Functionality
- In-memory topic management
- Simple pub/sub mechanism
- Single-node support

Stage 2: Enhanced Features
- Persistent message storage
- Message retention policies
- Basic error handling

Stage 3: Advanced Capabilities
- Client connection management
- TLS support
- Performance optimizations

Tech Stack:
- Rust (async/await)
- tokio for async runtime
- serde for serialization
- uuid for message/subscriber tracking

Development Timeline:
- Stage 1: 2-3 weeks
- Stage 2: 2-3 weeks
- Stage 3: 3-4 weeks

Code Structure:
```
rustmq/
├── src/
│   ├── broker.rs        # Core broker logic
│   ├── message.rs       # Message structure
│   ├── subscriber.rs    # Subscriber management
│   ├── storage.rs       # Message persistence
│   └── main.rs
├── examples/
│   ├── simple_producer.rs
│   └── simple_consumer.rs
└── tests/
    ├── broker_tests.rs
    └── integration_tests.rs
```

Learning Objectives:
1. Rust async programming
2. Concurrent system design
3. Message queue internals
4. Error handling
5. Performance optimization

Potential Challenges:
- Efficient message routing
- Scalable subscriber management
- Handling backpressure
- Performance under load

Recommended Tools/Libraries:
- `tokio` for async runtime
- `tracing` for logging
- `criterion` for benchmarking
- `proptest` for property-based testing
