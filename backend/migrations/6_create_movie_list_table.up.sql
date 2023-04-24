CREATE TABLE public.movie_list (
	"list_id" uuid NOT NULL,
	"movie_id" uuid NOT NULL,
	CONSTRAINT "movie_list_un" UNIQUE (list_id,movie_id),
	CONSTRAINT "movie_list_list_id_fk" FOREIGN KEY (list_id) REFERENCES public.list(id),
	CONSTRAINT "movie_list_movie_id_fk" FOREIGN KEY (movie_id) REFERENCES public.movie(id)
);

