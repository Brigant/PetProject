CREATE TABLE public."session" (
	"refresh_token" uuid NOT NULL DEFAULT gen_random_uuid(),
	"account_id" uuid NOT NULL,
	"request_host" varchar(255),
	"user_agent" varchar(255),
	"client_ip" varchar(32),
	"is_blocked" boolean,
	"expired_in" bigint,
	"created" timestamp with time zone NOT NULL DEFAULT now(),
	CONSTRAINT "session_pk" PRIMARY KEY (refresh_token),
	CONSTRAINT "session_fk" FOREIGN KEY (account_id) REFERENCES public."account"(id)
);
