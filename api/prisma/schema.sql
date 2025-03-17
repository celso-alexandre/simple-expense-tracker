--
-- PostgreSQL database dump
--

-- Dumped from database version 15.8
-- Dumped by pg_dump version 16.3 (Debian 16.3-1.pgdg120+1)

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

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: expense_plan_category; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.expense_plan_category AS ENUM (
    'FOOD',
    'TRANSPORT',
    'PROPERTY',
    'TAX',
    'ENTERTAINMENT',
    'OTHER'
);


ALTER TYPE public.expense_plan_category OWNER TO postgres;

--
-- Name: recurrency_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.recurrency_type AS ENUM (
    'MONTHLY',
    'YEARLY'
);


ALTER TYPE public.recurrency_type OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: expense_plan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.expense_plan (
    expense_plan_id integer NOT NULL,
    title text NOT NULL,
    amount_planned integer NOT NULL,
    last_paid_date timestamp with time zone,
    last_amount_spent integer DEFAULT 0 NOT NULL,
    paid_count integer DEFAULT 0 NOT NULL,
    recurrency_type public.recurrency_type,
    recurrency_interval integer DEFAULT 0 NOT NULL,
    category public.expense_plan_category NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    first_paid_date timestamp with time zone
);


ALTER TABLE public.expense_plan OWNER TO postgres;

--
-- Name: expense_plan_expense_plan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.expense_plan_expense_plan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.expense_plan_expense_plan_id_seq OWNER TO postgres;

--
-- Name: expense_plan_expense_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.expense_plan_expense_plan_id_seq OWNED BY public.expense_plan.expense_plan_id;


--
-- Name: expense_plan_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.expense_plan_record (
    expense_plan_record_id integer NOT NULL,
    expense_plan_id integer NOT NULL,
    amount_paid integer NOT NULL,
    payment_date timestamp with time zone NOT NULL,
    paid_date timestamp with time zone NOT NULL,
    expense_plan_sequence integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.expense_plan_record OWNER TO postgres;

--
-- Name: expense_plan_record_expense_plan_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.expense_plan_record_expense_plan_record_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.expense_plan_record_expense_plan_record_id_seq OWNER TO postgres;

--
-- Name: expense_plan_record_expense_plan_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.expense_plan_record_expense_plan_record_id_seq OWNED BY public.expense_plan_record.expense_plan_record_id;


--
-- Name: expense_plan expense_plan_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense_plan ALTER COLUMN expense_plan_id SET DEFAULT nextval('public.expense_plan_expense_plan_id_seq'::regclass);


--
-- Name: expense_plan_record expense_plan_record_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense_plan_record ALTER COLUMN expense_plan_record_id SET DEFAULT nextval('public.expense_plan_record_expense_plan_record_id_seq'::regclass);


--
-- Name: expense_plan expense_plan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense_plan
    ADD CONSTRAINT expense_plan_pkey PRIMARY KEY (expense_plan_id);


--
-- Name: expense_plan_record expense_plan_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense_plan_record
    ADD CONSTRAINT expense_plan_record_pkey PRIMARY KEY (expense_plan_record_id);


--
-- Name: expense_plan_record expense_plan_record_expense_plan_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense_plan_record
    ADD CONSTRAINT expense_plan_record_expense_plan_id_fkey FOREIGN KEY (expense_plan_id) REFERENCES public.expense_plan(expense_plan_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--

