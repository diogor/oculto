CREATE TABLE game (
    id uuid PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE player (
    id uuid PRIMARY KEY,
    name text NOT NULL,
    game_id uuid NOT NULL REFERENCES game(id),
    has_picked boolean NOT NULL DEFAULT false,
    is_picked boolean NOT NULL DEFAULT false
);

CREATE TABLE pick (
    id uuid PRIMARY KEY,
    game_id uuid NOT NULL REFERENCES game(id),
    picked_by uuid NOT NULL REFERENCES player(id),
    player_id uuid NOT NULL REFERENCES player(id)
);

CREATE UNIQUE INDEX ON pick(game_id, player_id, picked_by);
