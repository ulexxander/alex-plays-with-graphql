CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  username text NOT NULL,
  password text NOT NULL,
  date_created timestamptz NOT NULL DEFAULT NOW(),
  date_updated timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
  message_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL,
  chat_id uuid NOT NULL,
  content text NOT NULL,
  date_created timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE chats (
  chat_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  title text NOT NULL,
  date_created timestamptz NOT NULL DEFAULT NOW(),
  date_updated timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE members (
  user_id uuid REFERENCES users(user_id),
  chat_id uuid REFERENCES chats(chat_id),

  PRIMARY KEY (user_id, chat_id)
);
