--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)

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
-- Name: accounts; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.accounts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    account_number character varying(100),
    balance numeric(19,2)
);


ALTER TABLE public.accounts OWNER TO rey;

--
-- Name: factors; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.factors (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    pin character varying(100)
);


ALTER TABLE public.factors OWNER TO rey;

--
-- Name: login_log; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.login_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    is_authorized boolean NOT NULL,
    ip_address character varying(255) NOT NULL,
    timezone character varying NOT NULL,
    lat numeric NOT NULL,
    lon numeric NOT NULL,
    access_time timestamp(0) without time zone NOT NULL
);


ALTER TABLE public.login_log OWNER TO rey;

--
-- Name: notifications; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.notifications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    status integer NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    is_read integer NOT NULL,
    created_at timestamp without time zone
);


ALTER TABLE public.notifications OWNER TO rey;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO rey;

--
-- Name: templates; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.templates (
    code character varying(255) NOT NULL,
    title text NOT NULL,
    body text NOT NULL
);


ALTER TABLE public.templates OWNER TO rey;

--
-- Name: topup; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.topup (
    id character varying(100) NOT NULL,
    user_id uuid NOT NULL,
    amount bigint NOT NULL,
    status integer DEFAULT 0 NOT NULL,
    snap_url character varying(255)
);


ALTER TABLE public.topup OWNER TO rey;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.transactions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sof_number character varying(100),
    dof_number character varying(100),
    amount numeric(19,2),
    transaction_type character varying(1),
    account_id uuid,
    transactions_datetime timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO rey;

--
-- Name: users; Type: TABLE; Schema: public; Owner: rey
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    fullname character varying(55),
    phone character varying(55),
    username character varying(255) NOT NULL,
    email character varying(55),
    password character varying(255) NOT NULL,
    email_verified_at timestamp(0) with time zone DEFAULT NULL::timestamp with time zone
);


ALTER TABLE public.users OWNER TO rey;

--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: factors factors_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.factors
    ADD CONSTRAINT factors_pkey PRIMARY KEY (id);


--
-- Name: login_log login_log_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.login_log
    ADD CONSTRAINT login_log_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: templates templates_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.templates
    ADD CONSTRAINT templates_pkey PRIMARY KEY (code);


--
-- Name: topup topup_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.topup
    ADD CONSTRAINT topup_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: rey
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: accounts accounts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: factors factors_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.factors
    ADD CONSTRAINT factors_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: login_log login_log_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.login_log
    ADD CONSTRAINT login_log_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: notifications notifications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: topup topup_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.topup
    ADD CONSTRAINT topup_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: transactions transactions_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rey
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

