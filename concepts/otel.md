# OpenTelemetry Overview

## Observability

To make a distributed system observable, we must model its state in a way that lets us reason about its behavior.  
This model is composed of three key factors:

1. **Workload**

   - Represents the operations a system performs to fulfill its objectives.
   - Example: When a user sends a request, the system may break it down into smaller tasks handled by different services.
   - Often referred to as **transactions**.

2. **Software Abstractions**

   - Structural elements of the distributed system.
   - Examples: Load balancers, services, pods, containers, etc.

3. **Physical Machines**
   - The actual hardware providing computational resources such as **RAM**, **CPU**, **disk space**, and **network** capacity.

---

## Three Pillars of Observability

1. **Metrics**
   - Numerical measurements that provide quantitative insights about system performance and behavior over time.
   - Examples: CPU usage, memory consumption, request count, response time, error rate.
   - Typically aggregated and stored as time series data for trend analysis and alerting.
   - Best for understanding system health, capacity planning, and setting up alerts.

2. **Traces**
   - Detailed records of individual requests as they flow through distributed systems.
   - Show the complete journey of a request across multiple services and components.
   - Include timing information, service dependencies, and error propagation.
   - Best for debugging performance bottlenecks, understanding system architecture, and root cause analysis.

3. **Logs**
   - Discrete event records that capture what happened in the system at specific points in time.
   - Contain detailed contextual information about application behavior, errors, and state changes.
   - Can be structured (JSON) or unstructured (plain text) format.
   - Best for debugging specific issues, auditing, and understanding application flow.

---

## Goals of OpenTelemetry (Otel)

- **Unified Telemetry**  
  Combines tracing, logging, and metrics into a single framework, enabling correlation of all data and establishing an open standard for telemetry.

- **Vendor-Neutrality**  
  Integrates with different backends for processing telemetry data.

- **Cross-Platform Support**  
  Works across various programming languages (Java, Python, Go, etc.) and platforms.

---

## What OpenTelemetry is **NOT**

- **Not an All-in-One Tool**  
  Does not replace observability platforms like Datadog, New Relic, or Prometheus.

- **Not for Data Storage or Visualization**  
  Otel collects and exports telemetry data but does not store or visualize it.  
  Use tools like **Grafana**, **Jaeger**, or **Prometheus** for that.

- **Not Pre-Configured**  
  Requires setup and integration with external systems — not a plug-and-play solution.

- **Not a Performance Optimizer**  
  Otel collects performance data but does not automatically improve application performance.

---

## Benefits of OpenTelemetry

- **Instrument Once, Use Everywhere**  
  Reuse the same telemetry data across multiple tools.

- **Separation of Concerns**  
  Keep telemetry generation separate from analysis.

- **Observability by Default**  
  Make observability a first-class citizen during software development.

- **Better Use of Telemetry**  
  Improve how telemetry data is collected, shared, and analyzed.
