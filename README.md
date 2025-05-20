# 🪙 Market Candle Generator

A real-time, scalable, Kafka-based log aggregation system.
The system simulates and publishes OHLC candle data for different timeframes and symbols, with a distributed, event-driven architecture designed for extensibility and real-world production use.

---

## 📌 Features

- ✅ Generate fake gold market OHLC candle data (`daily`, `4h`, `1h`, `1m`)
- ✅ Modular Kafka producer with dynamic configuration
- ✅ Keyed publishing via Kafka for efficient partitioning (e.g., `gold-1h`, `gold-daily`)
- ✅ Configurable via environment variables
- ✅ Shared topic architecture for better scalability: `market.candles`
- ✅ Dockerized services for fast local deployment
- ✅ Scalable multi-instance producer setup

---

## 🧱 Project Structure
/cmd --> Entry point for producers

/internal

/kafka --> Kafka connection and writer logic

/service --> Data generators (OHLC for different timeframes)

/domain --> Shared data structures (Candle)

/Dockerfile --> Optimized builder for Go

/docker-compose.yml --> Multi-service orchestration



---

## ⚙️ Configuration

Each producer instance is configurable via environment variables:

| Variable      | Description                          | Example                    |
|---------------|--------------------------------------|----------------------------|
| `BROKER_ADDR` | Kafka broker address                 | `kafka:9092`               |
| `SYMBOL`      | Symbol to simulate                   | `GOLD`, `SILVER`, `COPPER` |
| `TIMEFRAME`   | Timeframe for generated candles      | `daily`, `4h`, `1h`, `1m`  |

Example Docker Compose config:

```yaml
  gold-producer-daily:
    build:
      context: .
    environment:
      BROKER_ADDR: kafka:9092
      SYMBOL: GOLD
      TIMEFRAME: daily
```


## 🧪 How It Works
- Producer reads SYMBOL and TIMEFRAME from env
- Candles are generated with OHLC format between startDate and endDate
- Each candle is published to Kafka topic: market.candles
- Kafka message key format: symbol-timeframe (e.g., gold-1h)
- Consumers (coming soon) can filter messages based on key


## 📦 Kafka Topic Strategy


| Topic         | Description                          |
|---------------|--------------------------------------|
| Shared topic for all candles, filtered by key   | ```market.candles```     |


Kafka Message Key example:

```json
  Key: "gold-daily"
  Value: {
      "symbol": "GOLD",
      "timeframe": "daily",
      "open": 1901.12,
      ...
    }
```


## 🚀 Roadmap
- ✅ Multi-timeframe OHLC generator
- ✅ Kafka writer with custom key support
- ✅ Dockerized producer with env-based config
- ✅ Kafka UI (Kafdrop / AKHQ / Redpanda Console)
- ☑️ Kafka consumer with key filtering
- ☑️ Local storage and alerting logic
- ☑️ Integration with real gold API (e.g., Tradingview API)
- ☑️  Search & Query API:
  A REST API that allows querying candles between two specific dates, filtered by symbol and timeframe.
- ☑️ Plug-and-play producer architecture:
  This means structuring your producer framework to be modular and plug-and-play, where each producer module is self-contained (with its own config, symbol, timeframe) and can register itself into a shared Kafka topic like market.candles.
- ☑️ Configurable Rule Engine via Kafka: For example, you could have a Kafka topic called market.rules where alert rules (e.g., "drop > 2%") are published externally.
  Then, your alert-engine would consume these rules dynamically from Kafka and update its behavior in real-time without needing a restart or redeploy.
- ☑️ A simple web dashboard (React + Chart.js) that visualizes candle data from Kafka or a database, potentially in real-time via WebSocket or polling.
