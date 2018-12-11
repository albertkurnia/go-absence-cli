--
-- PostgreSQL database dump
--

-- Dumped from database version 10.5
-- Dumped by pg_dump version 10.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: absent_history; Type: TABLE; Schema: public; Owner: albert
--

CREATE TABLE public.absent_history (
    id integer NOT NULL,
    "employeeId" integer,
    status character varying(5),
    "absentTime" timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.absent_history OWNER TO albert;

--
-- Name: absent_history_id_seq; Type: SEQUENCE; Schema: public; Owner: albert
--

CREATE SEQUENCE public.absent_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.absent_history_id_seq OWNER TO albert;

--
-- Name: absent_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: albert
--

ALTER SEQUENCE public.absent_history_id_seq OWNED BY public.absent_history.id;


--
-- Name: division; Type: TABLE; Schema: public; Owner: albert
--

CREATE TABLE public.division (
    id integer NOT NULL,
    name character varying(40)
);


ALTER TABLE public.division OWNER TO albert;

--
-- Name: employee; Type: TABLE; Schema: public; Owner: albert
--

CREATE TABLE public.employee (
    id integer NOT NULL,
    "divisionId" integer,
    name character varying(40),
    status character varying(20),
    createdat timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.employee OWNER TO albert;

--
-- Name: absent_history id; Type: DEFAULT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.absent_history ALTER COLUMN id SET DEFAULT nextval('public.absent_history_id_seq'::regclass);


--
-- Name: absent_history absent_history_pkey; Type: CONSTRAINT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.absent_history
    ADD CONSTRAINT absent_history_pkey PRIMARY KEY (id);


--
-- Name: division division_pkey; Type: CONSTRAINT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.division
    ADD CONSTRAINT division_pkey PRIMARY KEY (id);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- Name: absent_history absent_history_employeeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.absent_history
    ADD CONSTRAINT absent_history_employeeid_fkey FOREIGN KEY ("employeeId") REFERENCES public.employee(id);


--
-- Name: employee employee_divisionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: albert
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_divisionid_fkey FOREIGN KEY ("divisionId") REFERENCES public.division(id);


--
-- PostgreSQL database dump complete
--

