--
-- PostgreSQL database dump
--

\restrict Dsze8VxLu1K9B5RvsfkvFW2DrWnpFG0ZfYXQlW75kzmuowuzuhho5Bqxc4lCa2w

-- Dumped from database version 16.11 (Debian 16.11-1.pgdg13+1)
-- Dumped by pg_dump version 16.11 (Debian 16.11-1.pgdg13+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: pastebin; Type: TABLE; Schema: public; Owner: postgres_db_user
--

CREATE TABLE public.pastebin (
    id integer NOT NULL,
    content text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    url_slug text NOT NULL
);


ALTER TABLE public.pastebin OWNER TO postgres_db_user;

--
-- Name: pastebin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres_db_user
--

CREATE SEQUENCE public.pastebin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pastebin_id_seq OWNER TO postgres_db_user;

--
-- Name: pastebin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres_db_user
--

ALTER SEQUENCE public.pastebin_id_seq OWNED BY public.pastebin.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres_db_user
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres_db_user;

--
-- Name: pastebin id; Type: DEFAULT; Schema: public; Owner: postgres_db_user
--

ALTER TABLE ONLY public.pastebin ALTER COLUMN id SET DEFAULT nextval('public.pastebin_id_seq'::regclass);


--
-- Data for Name: pastebin; Type: TABLE DATA; Schema: public; Owner: postgres_db_user
--

COPY public.pastebin (id, content, created_at, url_slug) FROM stdin;
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres_db_user
--

COPY public.schema_migrations (version, dirty) FROM stdin;
12	f
\.


--
-- Name: pastebin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres_db_user
--

SELECT pg_catalog.setval('public.pastebin_id_seq', 1, false);


--
-- Name: pastebin pastebin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres_db_user
--

ALTER TABLE ONLY public.pastebin
    ADD CONSTRAINT pastebin_pkey PRIMARY KEY (id);


--
-- Name: pastebin pastebin_url_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres_db_user
--

ALTER TABLE ONLY public.pastebin
    ADD CONSTRAINT pastebin_url_slug_key UNIQUE (url_slug);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres_db_user
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--

\unrestrict Dsze8VxLu1K9B5RvsfkvFW2DrWnpFG0ZfYXQlW75kzmuowuzuhho5Bqxc4lCa2w

