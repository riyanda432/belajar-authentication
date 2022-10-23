CREATE TABLE IF NOT EXISTS users (
	id bigserial not null,
	full_name varchar(35) not null,
	user_name varchar not null,
    password varchar not null,
    created_at timestamptz not null default now(),
	updated_at timestamptz not null default now(),
	deleted_at timestamptz null,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);