CREATE TABLE notification (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    message text,
    user_id uuid REFERENCES "users"(id)
);
