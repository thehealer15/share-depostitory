--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: investor_abc123; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA investor_abc123;


ALTER SCHEMA investor_abc123 OWNER TO postgres;

--
-- Name: investor_xyz699; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA investor_xyz699;


ALTER SCHEMA investor_xyz699 OWNER TO postgres;

--
-- Name: investor_xyz789; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA investor_xyz789;


ALTER SCHEMA investor_xyz789 OWNER TO postgres;

--
-- Name: platform; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA platform;


ALTER SCHEMA platform OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: holdings; Type: TABLE; Schema: investor_abc123; Owner: postgres
--

CREATE TABLE investor_abc123.holdings (
    ticker text NOT NULL,
    quantity integer
);


ALTER TABLE investor_abc123.holdings OWNER TO postgres;

--
-- Name: investor_details; Type: TABLE; Schema: investor_abc123; Owner: postgres
--

CREATE TABLE investor_abc123.investor_details (
    govt_id text NOT NULL
);


ALTER TABLE investor_abc123.investor_details OWNER TO postgres;

--
-- Name: holdings; Type: TABLE; Schema: investor_xyz699; Owner: postgres
--

CREATE TABLE investor_xyz699.holdings (
    ticker text NOT NULL,
    quantity integer
);


ALTER TABLE investor_xyz699.holdings OWNER TO postgres;

--
-- Name: investor_details; Type: TABLE; Schema: investor_xyz699; Owner: postgres
--

CREATE TABLE investor_xyz699.investor_details (
    govt_id text NOT NULL
);


ALTER TABLE investor_xyz699.investor_details OWNER TO postgres;

--
-- Name: holdings; Type: TABLE; Schema: investor_xyz789; Owner: postgres
--

CREATE TABLE investor_xyz789.holdings (
    ticker text NOT NULL,
    quantity integer
);


ALTER TABLE investor_xyz789.holdings OWNER TO postgres;

--
-- Name: investor_details; Type: TABLE; Schema: investor_xyz789; Owner: postgres
--

CREATE TABLE investor_xyz789.investor_details (
    govt_id text NOT NULL
);


ALTER TABLE investor_xyz789.investor_details OWNER TO postgres;

--
-- Name: companies; Type: TABLE; Schema: platform; Owner: postgres
--

CREATE TABLE platform.companies (
    ticker text NOT NULL,
    face_value integer,
    company_name text
);


ALTER TABLE platform.companies OWNER TO postgres;

--
-- Name: investor; Type: TABLE; Schema: platform; Owner: postgres
--

CREATE TABLE platform.investor (
    investor_name text,
    govt_id text NOT NULL
);


ALTER TABLE platform.investor OWNER TO postgres;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.companies (
    ticker text NOT NULL,
    face_value integer,
    company_name text
);


ALTER TABLE public.companies OWNER TO postgres;

--
-- Data for Name: holdings; Type: TABLE DATA; Schema: investor_abc123; Owner: postgres
--

COPY investor_abc123.holdings (ticker, quantity) FROM stdin;
\.


--
-- Data for Name: investor_details; Type: TABLE DATA; Schema: investor_abc123; Owner: postgres
--

COPY investor_abc123.investor_details (govt_id) FROM stdin;
ABC123
\.


--
-- Data for Name: holdings; Type: TABLE DATA; Schema: investor_xyz699; Owner: postgres
--

COPY investor_xyz699.holdings (ticker, quantity) FROM stdin;
fff	10
INTC	20
AAPL	200
\.


--
-- Data for Name: investor_details; Type: TABLE DATA; Schema: investor_xyz699; Owner: postgres
--

COPY investor_xyz699.investor_details (govt_id) FROM stdin;
xyz699
\.


--
-- Data for Name: holdings; Type: TABLE DATA; Schema: investor_xyz789; Owner: postgres
--

COPY investor_xyz789.holdings (ticker, quantity) FROM stdin;
\.


--
-- Data for Name: investor_details; Type: TABLE DATA; Schema: investor_xyz789; Owner: postgres
--

COPY investor_xyz789.investor_details (govt_id) FROM stdin;
XYZ789
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: platform; Owner: postgres
--

COPY platform.companies (ticker, face_value, company_name) FROM stdin;
MSFT	200	Microsoft Inc.
AAPL	150	Apple Inc.
GOOGL	300	Google LLC
AMZN	250	Amazon Inc.
TSLA	180	Tesla Inc.
META	220	Meta Platforms
NVDA	270	NVIDIA Corp
INTC	190	Intel Corp
CSCO	210	Cisco Systems
ORCL	230	Oracle Corp
GAVAli	69	Gavali
DMART	69	Gavali
cultfit	150	Apple Inc.
\.


--
-- Data for Name: investor; Type: TABLE DATA; Schema: platform; Owner: postgres
--

COPY platform.investor (investor_name, govt_id) FROM stdin;
John Doe	ABC123
Akshay Pawar	xyz699
Jane Doe	XYZ789
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.companies (ticker, face_value, company_name) FROM stdin;
AAPL	100	Apple Inc.
MSFT	200	Microsoft Inc.
\.


--
-- Name: holdings holdings_pkey; Type: CONSTRAINT; Schema: investor_abc123; Owner: postgres
--

ALTER TABLE ONLY investor_abc123.holdings
    ADD CONSTRAINT holdings_pkey PRIMARY KEY (ticker);


--
-- Name: investor_details investor_details_pkey; Type: CONSTRAINT; Schema: investor_abc123; Owner: postgres
--

ALTER TABLE ONLY investor_abc123.investor_details
    ADD CONSTRAINT investor_details_pkey PRIMARY KEY (govt_id);


--
-- Name: holdings holdings_pkey; Type: CONSTRAINT; Schema: investor_xyz699; Owner: postgres
--

ALTER TABLE ONLY investor_xyz699.holdings
    ADD CONSTRAINT holdings_pkey PRIMARY KEY (ticker);


--
-- Name: investor_details investor_details_pkey; Type: CONSTRAINT; Schema: investor_xyz699; Owner: postgres
--

ALTER TABLE ONLY investor_xyz699.investor_details
    ADD CONSTRAINT investor_details_pkey PRIMARY KEY (govt_id);


--
-- Name: holdings holdings_pkey; Type: CONSTRAINT; Schema: investor_xyz789; Owner: postgres
--

ALTER TABLE ONLY investor_xyz789.holdings
    ADD CONSTRAINT holdings_pkey PRIMARY KEY (ticker);


--
-- Name: investor_details investor_details_pkey; Type: CONSTRAINT; Schema: investor_xyz789; Owner: postgres
--

ALTER TABLE ONLY investor_xyz789.investor_details
    ADD CONSTRAINT investor_details_pkey PRIMARY KEY (govt_id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: platform; Owner: postgres
--

ALTER TABLE ONLY platform.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (ticker);


--
-- Name: investor investor_pkey; Type: CONSTRAINT; Schema: platform; Owner: postgres
--

ALTER TABLE ONLY platform.investor
    ADD CONSTRAINT investor_pkey PRIMARY KEY (govt_id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (ticker);


--
-- PostgreSQL database dump complete
--

