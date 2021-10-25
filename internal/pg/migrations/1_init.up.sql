CREATE TABLE teams (
	"id" integer NOT NULL,
	"rating" numeric NOT NULL,
	"name" character varying(255) NULL,
	CONSTRAINT "pk_team_id" PRIMARY KEY (id)
);