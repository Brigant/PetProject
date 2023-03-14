CREATE TABLE public."session" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"account_id" uuid NOT NULL,
	"refresh_token" varchar NOT NULL,
	"user_agent" varchar NOT NULL,
	"client_ip" varchar NOT NULL,
	"is_blocked" boolean NOT NULL,
	"expires_at" timestamp with time zone NOT NULL,
	"created" timestamp with time zone NOT NULL DEFAULT now(),
	CONSTRAINT "session_pk" PRIMARY KEY (id),
	CONSTRAINT "session_fk" FOREIGN KEY (account_id) REFERENCES public."account"(id)
);
