CREATE TABLE IF NOT EXISTS events (
  id           BIGSERIAL PRIMARY KEY,
  name         TEXT NOT NULL,
  description  TEXT NOT NULL,
  category     VARCHAR(30) NOT NULL,
  location     VARCHAR(30) NOT NULL,
  publisher    VARCHAR(30) NOT NULL,
  lineup       TEXT[] NOT NULL,
  start_ts     TIMESTAMP NOT NULL,
  end_ts       TIMESTAMP NOT NULL,
  created_at   TIMESTAMP DEFAULT NOW() NOT NULL
);
