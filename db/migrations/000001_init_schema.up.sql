--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

-- Started on 2024-03-07 18:16:30 UTC

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
-- TOC entry 222 (class 1259 OID 16428)
-- Name: maintenance_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.maintenance_logs (
    log_id integer NOT NULL,
    slot_id integer NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone,
    note text
);


ALTER TABLE public.maintenance_logs OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16427)
-- Name: maintenance_logs_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.maintenance_logs_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.maintenance_logs_log_id_seq OWNER TO postgres;

--
-- TOC entry 3388 (class 0 OID 0)
-- Dependencies: 221
-- Name: maintenance_logs_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.maintenance_logs_log_id_seq OWNED BY public.maintenance_logs.log_id;


--
-- TOC entry 216 (class 1259 OID 16390)
-- Name: parking_lots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.parking_lots (
    lot_id integer NOT NULL,
    name character varying(255) NOT NULL,
    capacity integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.parking_lots OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16389)
-- Name: parking_lots_lot_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.parking_lots_lot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.parking_lots_lot_id_seq OWNER TO postgres;

--
-- TOC entry 3389 (class 0 OID 0)
-- Dependencies: 215
-- Name: parking_lots_lot_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.parking_lots_lot_id_seq OWNED BY public.parking_lots.lot_id;


--
-- TOC entry 220 (class 1259 OID 16414)
-- Name: parking_sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.parking_sessions (
    session_id integer NOT NULL,
    slot_id integer NOT NULL,
    vehicle_number character varying(255) NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone,
    fee integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.parking_sessions OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16413)
-- Name: parking_sessions_session_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.parking_sessions_session_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.parking_sessions_session_id_seq OWNER TO postgres;

--
-- TOC entry 3390 (class 0 OID 0)
-- Dependencies: 219
-- Name: parking_sessions_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.parking_sessions_session_id_seq OWNED BY public.parking_sessions.session_id;


--
-- TOC entry 218 (class 1259 OID 16399)
-- Name: parking_slots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.parking_slots (
    slot_id integer NOT NULL,
    lot_id integer NOT NULL,
    slot_number integer NOT NULL,
    status character varying(50) DEFAULT 'available'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.parking_slots OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16398)
-- Name: parking_slots_slot_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.parking_slots_slot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.parking_slots_slot_id_seq OWNER TO postgres;

--
-- TOC entry 3391 (class 0 OID 0)
-- Dependencies: 217
-- Name: parking_slots_slot_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.parking_slots_slot_id_seq OWNED BY public.parking_slots.slot_id;


--
-- TOC entry 3228 (class 2604 OID 16431)
-- Name: maintenance_logs log_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.maintenance_logs ALTER COLUMN log_id SET DEFAULT nextval('public.maintenance_logs_log_id_seq'::regclass);


--
-- TOC entry 3218 (class 2604 OID 16393)
-- Name: parking_lots lot_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_lots ALTER COLUMN lot_id SET DEFAULT nextval('public.parking_lots_lot_id_seq'::regclass);


--
-- TOC entry 3225 (class 2604 OID 16417)
-- Name: parking_sessions session_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_sessions ALTER COLUMN session_id SET DEFAULT nextval('public.parking_sessions_session_id_seq'::regclass);


--
-- TOC entry 3221 (class 2604 OID 16402)
-- Name: parking_slots slot_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_slots ALTER COLUMN slot_id SET DEFAULT nextval('public.parking_slots_slot_id_seq'::regclass);


--
-- TOC entry 3236 (class 2606 OID 16435)
-- Name: maintenance_logs maintenance_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.maintenance_logs
    ADD CONSTRAINT maintenance_logs_pkey PRIMARY KEY (log_id);


--
-- TOC entry 3230 (class 2606 OID 16397)
-- Name: parking_lots parking_lots_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_lots
    ADD CONSTRAINT parking_lots_pkey PRIMARY KEY (lot_id);


--
-- TOC entry 3234 (class 2606 OID 16421)
-- Name: parking_sessions parking_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_sessions
    ADD CONSTRAINT parking_sessions_pkey PRIMARY KEY (session_id);


--
-- TOC entry 3232 (class 2606 OID 16407)
-- Name: parking_slots parking_slots_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_slots
    ADD CONSTRAINT parking_slots_pkey PRIMARY KEY (slot_id);


--
-- TOC entry 3239 (class 2606 OID 16436)
-- Name: maintenance_logs maintenance_logs_slot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.maintenance_logs
    ADD CONSTRAINT maintenance_logs_slot_id_fkey FOREIGN KEY (slot_id) REFERENCES public.parking_slots(slot_id);


--
-- TOC entry 3238 (class 2606 OID 16422)
-- Name: parking_sessions parking_sessions_slot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_sessions
    ADD CONSTRAINT parking_sessions_slot_id_fkey FOREIGN KEY (slot_id) REFERENCES public.parking_slots(slot_id);


--
-- TOC entry 3237 (class 2606 OID 16408)
-- Name: parking_slots parking_slots_lot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_slots
    ADD CONSTRAINT parking_slots_lot_id_fkey FOREIGN KEY (lot_id) REFERENCES public.parking_lots(lot_id);


-- Completed on 2024-03-07 18:16:30 UTC

--
-- PostgreSQL database dump complete
--

