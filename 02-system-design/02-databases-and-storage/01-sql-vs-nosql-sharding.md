# System Design Урок 2: Бази даних, Шардинг та Реплікація

## 1. SQL vs NoSQL
- **SQL (Relational):** PostgreSQL, MySQL. ACID-транзакції, нормалізовані дані, складні з'єднання (JOIN). Добре для фінансових систем.
- **NoSQL (Non-Relational):** 
  - *Key-Value* (Redis, DynamoDB): надшвидкий доступ за ключем.
  - *Document* (MongoDB): гнучка схема JSON.
  - *Column-family* (Cassandra): висока швидкість запису масивів даних.

## 2. Реплікація (Replication)
- **Master-Slave (Leader-Follower):** Записи йдуть на Leader, читання масштабується на Follower-репліки. Ризик затримки реплікації (replication lag).
- **Master-Master:** Записи на обидва вузли, складне вирішення конфліктів.

## 3. Шардинг (Sharding)
Горизонтальне розділення бази даних на менші частини (шарди) за критерієм:
- **Range-based:** за діапазоном (наприклад, по датах або ID).
- **Hash-based:** за хешем від ключа (`hash(user_id) % N_shards`). Ризик нерівномірного розподілу (hotspots).
- **Directory-based:** використання сервісу-маршрутизатора.
