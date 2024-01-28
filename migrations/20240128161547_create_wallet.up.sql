CREATE TABLE wallet (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    balance REAL NOT NULL DEFAULT 100.0
);

CREATE TABLE history (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    time TIMESTAMP WITH TIME ZONE NOT NULL,
    "from" UUID NOT NULL,
    "to" UUID NOT NULL,
    amount REAL NOT NULL,
    FOREIGN KEY ("from") REFERENCES wallet(id),
    FOREIGN KEY ("to") REFERENCES wallet(id)
);