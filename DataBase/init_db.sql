CREATE TABLE IF NOT EXISTS public.user(
    user_name       TEXT,
    password_hash   TEXT,
    CONSTRAINT users_pk PRIMARY KEY(user_name)
);
