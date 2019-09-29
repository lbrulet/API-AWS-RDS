BEGIN;

CREATE Table Trip(
    id int NOT NULL AUTO_INCREMENT,
    start_lat float NOT NULL,
    start_lng float NOT NULL,
    end_lat float NOT NULL,
    end_lng float NOT NULL,
    id_user int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_user) REFERENCES User(id)
);

COMMIT;