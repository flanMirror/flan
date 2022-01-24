--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Name: antenna_src_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.antenna_src_enum AS ENUM (
    'home',
    'all',
    'users',
    'list',
    'group'
);


ALTER TYPE public.antenna_src_enum OWNER TO misskey;

--
-- Name: log_level_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.log_level_enum AS ENUM (
    'error',
    'warning',
    'info',
    'success',
    'debug'
);


ALTER TYPE public.log_level_enum OWNER TO misskey;

--
-- Name: muted_note_reason_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.muted_note_reason_enum AS ENUM (
    'word',
    'manual',
    'spam',
    'other'
);


ALTER TYPE public.muted_note_reason_enum OWNER TO misskey;

--
-- Name: note_visibility_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.note_visibility_enum AS ENUM (
    'public',
    'home',
    'followers',
    'specified'
);


ALTER TYPE public.note_visibility_enum OWNER TO misskey;

--
-- Name: notification_type_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.notification_type_enum AS ENUM (
    'follow',
    'mention',
    'reply',
    'renote',
    'quote',
    'reaction',
    'pollVote',
    'receiveFollowRequest',
    'followRequestAccepted',
    'groupInvited',
    'app'
);


ALTER TYPE public.notification_type_enum OWNER TO misskey;

--
-- Name: page_visibility_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.page_visibility_enum AS ENUM (
    'public',
    'followers',
    'specified'
);


ALTER TYPE public.page_visibility_enum OWNER TO misskey;

--
-- Name: poll_notevisibility_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.poll_notevisibility_enum AS ENUM (
    'public',
    'home',
    'followers',
    'specified'
);


ALTER TYPE public.poll_notevisibility_enum OWNER TO misskey;

--
-- Name: relay_status_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.relay_status_enum AS ENUM (
    'requesting',
    'accepted',
    'rejected'
);


ALTER TYPE public.relay_status_enum OWNER TO misskey;

--
-- Name: user_profile_ffvisibility_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.user_profile_ffvisibility_enum AS ENUM (
    'public',
    'followers',
    'private'
);


ALTER TYPE public.user_profile_ffvisibility_enum OWNER TO misskey;

--
-- Name: user_profile_mutingnotificationtypes_enum; Type: TYPE; Schema: public; Owner: misskey
--

CREATE TYPE public.user_profile_mutingnotificationtypes_enum AS ENUM (
    'follow',
    'mention',
    'reply',
    'renote',
    'quote',
    'reaction',
    'pollVote',
    'receiveFollowRequest',
    'followRequestAccepted',
    'groupInvited',
    'app'
);


ALTER TYPE public.user_profile_mutingnotificationtypes_enum OWNER TO misskey;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: __chart__active_users; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__active_users (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_users character varying[] NOT NULL,
    ___remote_users character varying[] NOT NULL
);


ALTER TABLE public.__chart__active_users OWNER TO misskey;

--
-- Name: __chart__active_users_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__active_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__active_users_id_seq OWNER TO misskey;

--
-- Name: __chart__active_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__active_users_id_seq OWNED BY public.__chart__active_users.id;


--
-- Name: __chart__drive; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__drive (
    id integer NOT NULL,
    date integer NOT NULL,
    "___local_totalCount" bigint NOT NULL,
    "___local_totalSize" bigint NOT NULL,
    "___local_incCount" bigint NOT NULL,
    "___local_incSize" bigint NOT NULL,
    "___local_decCount" bigint NOT NULL,
    "___local_decSize" bigint NOT NULL,
    "___remote_totalCount" bigint NOT NULL,
    "___remote_totalSize" bigint NOT NULL,
    "___remote_incCount" bigint NOT NULL,
    "___remote_incSize" bigint NOT NULL,
    "___remote_decCount" bigint NOT NULL,
    "___remote_decSize" bigint NOT NULL
);


ALTER TABLE public.__chart__drive OWNER TO misskey;

--
-- Name: __chart__drive_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__drive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__drive_id_seq OWNER TO misskey;

--
-- Name: __chart__drive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__drive_id_seq OWNED BY public.__chart__drive.id;


--
-- Name: __chart__federation; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__federation (
    id integer NOT NULL,
    date integer NOT NULL,
    ___instance_total bigint NOT NULL,
    ___instance_inc bigint NOT NULL,
    ___instance_dec bigint NOT NULL
);


ALTER TABLE public.__chart__federation OWNER TO misskey;

--
-- Name: __chart__federation_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__federation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__federation_id_seq OWNER TO misskey;

--
-- Name: __chart__federation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__federation_id_seq OWNED BY public.__chart__federation.id;


--
-- Name: __chart__hashtag; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__hashtag (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_users character varying[] NOT NULL,
    ___remote_users character varying[] NOT NULL
);


ALTER TABLE public.__chart__hashtag OWNER TO misskey;

--
-- Name: __chart__hashtag_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__hashtag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__hashtag_id_seq OWNER TO misskey;

--
-- Name: __chart__hashtag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__hashtag_id_seq OWNED BY public.__chart__hashtag.id;


--
-- Name: __chart__instance; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__instance (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___requests_failed bigint NOT NULL,
    ___requests_succeeded bigint NOT NULL,
    ___requests_received bigint NOT NULL,
    ___notes_total bigint NOT NULL,
    ___notes_inc bigint NOT NULL,
    ___notes_dec bigint NOT NULL,
    ___notes_diffs_normal bigint NOT NULL,
    ___notes_diffs_reply bigint NOT NULL,
    ___notes_diffs_renote bigint NOT NULL,
    ___users_total bigint NOT NULL,
    ___users_inc bigint NOT NULL,
    ___users_dec bigint NOT NULL,
    ___following_total bigint NOT NULL,
    ___following_inc bigint NOT NULL,
    ___following_dec bigint NOT NULL,
    ___followers_total bigint NOT NULL,
    ___followers_inc bigint NOT NULL,
    ___followers_dec bigint NOT NULL,
    "___drive_totalFiles" bigint NOT NULL,
    "___drive_totalUsage" bigint NOT NULL,
    "___drive_incFiles" bigint NOT NULL,
    "___drive_incUsage" bigint NOT NULL,
    "___drive_decFiles" bigint NOT NULL,
    "___drive_decUsage" bigint NOT NULL
);


ALTER TABLE public.__chart__instance OWNER TO misskey;

--
-- Name: __chart__instance_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__instance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__instance_id_seq OWNER TO misskey;

--
-- Name: __chart__instance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__instance_id_seq OWNED BY public.__chart__instance.id;


--
-- Name: __chart__network; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__network (
    id integer NOT NULL,
    date integer NOT NULL,
    "___incomingRequests" bigint NOT NULL,
    "___outgoingRequests" bigint NOT NULL,
    "___totalTime" bigint NOT NULL,
    "___incomingBytes" bigint NOT NULL,
    "___outgoingBytes" bigint NOT NULL
);


ALTER TABLE public.__chart__network OWNER TO misskey;

--
-- Name: __chart__network_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__network_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__network_id_seq OWNER TO misskey;

--
-- Name: __chart__network_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__network_id_seq OWNED BY public.__chart__network.id;


--
-- Name: __chart__notes; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__notes (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_total bigint NOT NULL,
    ___local_inc bigint NOT NULL,
    ___local_dec bigint NOT NULL,
    ___local_diffs_normal bigint NOT NULL,
    ___local_diffs_reply bigint NOT NULL,
    ___local_diffs_renote bigint NOT NULL,
    ___remote_total bigint NOT NULL,
    ___remote_inc bigint NOT NULL,
    ___remote_dec bigint NOT NULL,
    ___remote_diffs_normal bigint NOT NULL,
    ___remote_diffs_reply bigint NOT NULL,
    ___remote_diffs_renote bigint NOT NULL
);


ALTER TABLE public.__chart__notes OWNER TO misskey;

--
-- Name: __chart__notes_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__notes_id_seq OWNER TO misskey;

--
-- Name: __chart__notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__notes_id_seq OWNED BY public.__chart__notes.id;


--
-- Name: __chart__per_user_drive; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__per_user_drive (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    "___totalCount" bigint NOT NULL,
    "___totalSize" bigint NOT NULL,
    "___incCount" bigint NOT NULL,
    "___incSize" bigint NOT NULL,
    "___decCount" bigint NOT NULL,
    "___decSize" bigint NOT NULL
);


ALTER TABLE public.__chart__per_user_drive OWNER TO misskey;

--
-- Name: __chart__per_user_drive_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__per_user_drive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__per_user_drive_id_seq OWNER TO misskey;

--
-- Name: __chart__per_user_drive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__per_user_drive_id_seq OWNED BY public.__chart__per_user_drive.id;


--
-- Name: __chart__per_user_following; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__per_user_following (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_followings_total bigint NOT NULL,
    ___local_followings_inc bigint NOT NULL,
    ___local_followings_dec bigint NOT NULL,
    ___local_followers_total bigint NOT NULL,
    ___local_followers_inc bigint NOT NULL,
    ___local_followers_dec bigint NOT NULL,
    ___remote_followings_total bigint NOT NULL,
    ___remote_followings_inc bigint NOT NULL,
    ___remote_followings_dec bigint NOT NULL,
    ___remote_followers_total bigint NOT NULL,
    ___remote_followers_inc bigint NOT NULL,
    ___remote_followers_dec bigint NOT NULL
);


ALTER TABLE public.__chart__per_user_following OWNER TO misskey;

--
-- Name: __chart__per_user_following_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__per_user_following_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__per_user_following_id_seq OWNER TO misskey;

--
-- Name: __chart__per_user_following_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__per_user_following_id_seq OWNED BY public.__chart__per_user_following.id;


--
-- Name: __chart__per_user_notes; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__per_user_notes (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___total bigint NOT NULL,
    ___inc bigint NOT NULL,
    ___dec bigint NOT NULL,
    ___diffs_normal bigint NOT NULL,
    ___diffs_reply bigint NOT NULL,
    ___diffs_renote bigint NOT NULL
);


ALTER TABLE public.__chart__per_user_notes OWNER TO misskey;

--
-- Name: __chart__per_user_notes_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__per_user_notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__per_user_notes_id_seq OWNER TO misskey;

--
-- Name: __chart__per_user_notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__per_user_notes_id_seq OWNED BY public.__chart__per_user_notes.id;


--
-- Name: __chart__per_user_reaction; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__per_user_reaction (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_count bigint NOT NULL,
    ___remote_count bigint NOT NULL
);


ALTER TABLE public.__chart__per_user_reaction OWNER TO misskey;

--
-- Name: __chart__per_user_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__per_user_reaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__per_user_reaction_id_seq OWNER TO misskey;

--
-- Name: __chart__per_user_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__per_user_reaction_id_seq OWNED BY public.__chart__per_user_reaction.id;


--
-- Name: __chart__test; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__test (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128),
    ___foo_total bigint NOT NULL,
    ___foo_inc bigint NOT NULL,
    ___foo_dec bigint NOT NULL
);


ALTER TABLE public.__chart__test OWNER TO misskey;

--
-- Name: __chart__test_grouped; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__test_grouped (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128),
    ___foo_total bigint NOT NULL,
    ___foo_inc bigint NOT NULL,
    ___foo_dec bigint NOT NULL
);


ALTER TABLE public.__chart__test_grouped OWNER TO misskey;

--
-- Name: __chart__test_grouped_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__test_grouped_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__test_grouped_id_seq OWNER TO misskey;

--
-- Name: __chart__test_grouped_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__test_grouped_id_seq OWNED BY public.__chart__test_grouped.id;


--
-- Name: __chart__test_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__test_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__test_id_seq OWNER TO misskey;

--
-- Name: __chart__test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__test_id_seq OWNED BY public.__chart__test.id;


--
-- Name: __chart__test_unique; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__test_unique (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128),
    ___foo character varying[] DEFAULT '{}'::character varying[] NOT NULL
);


ALTER TABLE public.__chart__test_unique OWNER TO misskey;

--
-- Name: __chart__test_unique_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__test_unique_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__test_unique_id_seq OWNER TO misskey;

--
-- Name: __chart__test_unique_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__test_unique_id_seq OWNED BY public.__chart__test_unique.id;


--
-- Name: __chart__users; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart__users (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_total bigint NOT NULL,
    ___local_inc bigint NOT NULL,
    ___local_dec bigint NOT NULL,
    ___remote_total bigint NOT NULL,
    ___remote_inc bigint NOT NULL,
    ___remote_dec bigint NOT NULL
);


ALTER TABLE public.__chart__users OWNER TO misskey;

--
-- Name: __chart__users_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart__users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart__users_id_seq OWNER TO misskey;

--
-- Name: __chart__users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart__users_id_seq OWNED BY public.__chart__users.id;


--
-- Name: __chart_day__active_users; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__active_users (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_users character varying[] NOT NULL,
    ___remote_users character varying[] NOT NULL
);


ALTER TABLE public.__chart_day__active_users OWNER TO misskey;

--
-- Name: __chart_day__active_users_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__active_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__active_users_id_seq OWNER TO misskey;

--
-- Name: __chart_day__active_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__active_users_id_seq OWNED BY public.__chart_day__active_users.id;


--
-- Name: __chart_day__drive; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__drive (
    id integer NOT NULL,
    date integer NOT NULL,
    "___local_totalCount" bigint NOT NULL,
    "___local_totalSize" bigint NOT NULL,
    "___local_incCount" bigint NOT NULL,
    "___local_incSize" bigint NOT NULL,
    "___local_decCount" bigint NOT NULL,
    "___local_decSize" bigint NOT NULL,
    "___remote_totalCount" bigint NOT NULL,
    "___remote_totalSize" bigint NOT NULL,
    "___remote_incCount" bigint NOT NULL,
    "___remote_incSize" bigint NOT NULL,
    "___remote_decCount" bigint NOT NULL,
    "___remote_decSize" bigint NOT NULL
);


ALTER TABLE public.__chart_day__drive OWNER TO misskey;

--
-- Name: __chart_day__drive_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__drive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__drive_id_seq OWNER TO misskey;

--
-- Name: __chart_day__drive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__drive_id_seq OWNED BY public.__chart_day__drive.id;


--
-- Name: __chart_day__federation; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__federation (
    id integer NOT NULL,
    date integer NOT NULL,
    ___instance_total bigint NOT NULL,
    ___instance_inc bigint NOT NULL,
    ___instance_dec bigint NOT NULL
);


ALTER TABLE public.__chart_day__federation OWNER TO misskey;

--
-- Name: __chart_day__federation_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__federation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__federation_id_seq OWNER TO misskey;

--
-- Name: __chart_day__federation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__federation_id_seq OWNED BY public.__chart_day__federation.id;


--
-- Name: __chart_day__hashtag; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__hashtag (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_users character varying[] NOT NULL,
    ___remote_users character varying[] NOT NULL
);


ALTER TABLE public.__chart_day__hashtag OWNER TO misskey;

--
-- Name: __chart_day__hashtag_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__hashtag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__hashtag_id_seq OWNER TO misskey;

--
-- Name: __chart_day__hashtag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__hashtag_id_seq OWNED BY public.__chart_day__hashtag.id;


--
-- Name: __chart_day__instance; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__instance (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___requests_failed bigint NOT NULL,
    ___requests_succeeded bigint NOT NULL,
    ___requests_received bigint NOT NULL,
    ___notes_total bigint NOT NULL,
    ___notes_inc bigint NOT NULL,
    ___notes_dec bigint NOT NULL,
    ___notes_diffs_normal bigint NOT NULL,
    ___notes_diffs_reply bigint NOT NULL,
    ___notes_diffs_renote bigint NOT NULL,
    ___users_total bigint NOT NULL,
    ___users_inc bigint NOT NULL,
    ___users_dec bigint NOT NULL,
    ___following_total bigint NOT NULL,
    ___following_inc bigint NOT NULL,
    ___following_dec bigint NOT NULL,
    ___followers_total bigint NOT NULL,
    ___followers_inc bigint NOT NULL,
    ___followers_dec bigint NOT NULL,
    "___drive_totalFiles" bigint NOT NULL,
    "___drive_totalUsage" bigint NOT NULL,
    "___drive_incFiles" bigint NOT NULL,
    "___drive_incUsage" bigint NOT NULL,
    "___drive_decFiles" bigint NOT NULL,
    "___drive_decUsage" bigint NOT NULL
);


ALTER TABLE public.__chart_day__instance OWNER TO misskey;

--
-- Name: __chart_day__instance_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__instance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__instance_id_seq OWNER TO misskey;

--
-- Name: __chart_day__instance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__instance_id_seq OWNED BY public.__chart_day__instance.id;


--
-- Name: __chart_day__network; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__network (
    id integer NOT NULL,
    date integer NOT NULL,
    "___incomingRequests" bigint NOT NULL,
    "___outgoingRequests" bigint NOT NULL,
    "___totalTime" bigint NOT NULL,
    "___incomingBytes" bigint NOT NULL,
    "___outgoingBytes" bigint NOT NULL
);


ALTER TABLE public.__chart_day__network OWNER TO misskey;

--
-- Name: __chart_day__network_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__network_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__network_id_seq OWNER TO misskey;

--
-- Name: __chart_day__network_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__network_id_seq OWNED BY public.__chart_day__network.id;


--
-- Name: __chart_day__notes; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__notes (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_total bigint NOT NULL,
    ___local_inc bigint NOT NULL,
    ___local_dec bigint NOT NULL,
    ___local_diffs_normal bigint NOT NULL,
    ___local_diffs_reply bigint NOT NULL,
    ___local_diffs_renote bigint NOT NULL,
    ___remote_total bigint NOT NULL,
    ___remote_inc bigint NOT NULL,
    ___remote_dec bigint NOT NULL,
    ___remote_diffs_normal bigint NOT NULL,
    ___remote_diffs_reply bigint NOT NULL,
    ___remote_diffs_renote bigint NOT NULL
);


ALTER TABLE public.__chart_day__notes OWNER TO misskey;

--
-- Name: __chart_day__notes_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__notes_id_seq OWNER TO misskey;

--
-- Name: __chart_day__notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__notes_id_seq OWNED BY public.__chart_day__notes.id;


--
-- Name: __chart_day__per_user_drive; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__per_user_drive (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    "___totalCount" bigint NOT NULL,
    "___totalSize" bigint NOT NULL,
    "___incCount" bigint NOT NULL,
    "___incSize" bigint NOT NULL,
    "___decCount" bigint NOT NULL,
    "___decSize" bigint NOT NULL
);


ALTER TABLE public.__chart_day__per_user_drive OWNER TO misskey;

--
-- Name: __chart_day__per_user_drive_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__per_user_drive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__per_user_drive_id_seq OWNER TO misskey;

--
-- Name: __chart_day__per_user_drive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__per_user_drive_id_seq OWNED BY public.__chart_day__per_user_drive.id;


--
-- Name: __chart_day__per_user_following; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__per_user_following (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_followings_total bigint NOT NULL,
    ___local_followings_inc bigint NOT NULL,
    ___local_followings_dec bigint NOT NULL,
    ___local_followers_total bigint NOT NULL,
    ___local_followers_inc bigint NOT NULL,
    ___local_followers_dec bigint NOT NULL,
    ___remote_followings_total bigint NOT NULL,
    ___remote_followings_inc bigint NOT NULL,
    ___remote_followings_dec bigint NOT NULL,
    ___remote_followers_total bigint NOT NULL,
    ___remote_followers_inc bigint NOT NULL,
    ___remote_followers_dec bigint NOT NULL
);


ALTER TABLE public.__chart_day__per_user_following OWNER TO misskey;

--
-- Name: __chart_day__per_user_following_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__per_user_following_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__per_user_following_id_seq OWNER TO misskey;

--
-- Name: __chart_day__per_user_following_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__per_user_following_id_seq OWNED BY public.__chart_day__per_user_following.id;


--
-- Name: __chart_day__per_user_notes; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__per_user_notes (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___total bigint NOT NULL,
    ___inc bigint NOT NULL,
    ___dec bigint NOT NULL,
    ___diffs_normal bigint NOT NULL,
    ___diffs_reply bigint NOT NULL,
    ___diffs_renote bigint NOT NULL
);


ALTER TABLE public.__chart_day__per_user_notes OWNER TO misskey;

--
-- Name: __chart_day__per_user_notes_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__per_user_notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__per_user_notes_id_seq OWNER TO misskey;

--
-- Name: __chart_day__per_user_notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__per_user_notes_id_seq OWNED BY public.__chart_day__per_user_notes.id;


--
-- Name: __chart_day__per_user_reaction; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__per_user_reaction (
    id integer NOT NULL,
    date integer NOT NULL,
    "group" character varying(128) NOT NULL,
    ___local_count bigint NOT NULL,
    ___remote_count bigint NOT NULL
);


ALTER TABLE public.__chart_day__per_user_reaction OWNER TO misskey;

--
-- Name: __chart_day__per_user_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__per_user_reaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__per_user_reaction_id_seq OWNER TO misskey;

--
-- Name: __chart_day__per_user_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__per_user_reaction_id_seq OWNED BY public.__chart_day__per_user_reaction.id;


--
-- Name: __chart_day__users; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.__chart_day__users (
    id integer NOT NULL,
    date integer NOT NULL,
    ___local_total bigint NOT NULL,
    ___local_inc bigint NOT NULL,
    ___local_dec bigint NOT NULL,
    ___remote_total bigint NOT NULL,
    ___remote_inc bigint NOT NULL,
    ___remote_dec bigint NOT NULL
);


ALTER TABLE public.__chart_day__users OWNER TO misskey;

--
-- Name: __chart_day__users_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.__chart_day__users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.__chart_day__users_id_seq OWNER TO misskey;

--
-- Name: __chart_day__users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.__chart_day__users_id_seq OWNED BY public.__chart_day__users.id;


--
-- Name: abuse_user_report; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.abuse_user_report (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "targetUserId" character varying(32) NOT NULL,
    "reporterId" character varying(32) NOT NULL,
    "assigneeId" character varying(32),
    resolved boolean DEFAULT false NOT NULL,
    comment character varying(2048) DEFAULT '{}'::character varying[] NOT NULL,
    "targetUserHost" character varying(128),
    "reporterHost" character varying(128)
);


ALTER TABLE public.abuse_user_report OWNER TO misskey;

--
-- Name: COLUMN abuse_user_report."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.abuse_user_report."createdAt" IS 'The created date of the AbuseUserReport.';


--
-- Name: COLUMN abuse_user_report."targetUserHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.abuse_user_report."targetUserHost" IS '[Denormalized]';


--
-- Name: COLUMN abuse_user_report."reporterHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.abuse_user_report."reporterHost" IS '[Denormalized]';


--
-- Name: access_token; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.access_token (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    token character varying(128) NOT NULL,
    hash character varying(128) NOT NULL,
    "userId" character varying(32) NOT NULL,
    "appId" character varying(32) DEFAULT NULL::character varying,
    "lastUsedAt" timestamp with time zone,
    session character varying(128) DEFAULT NULL::character varying,
    name character varying(128) DEFAULT NULL::character varying,
    description character varying(512) DEFAULT NULL::character varying,
    "iconUrl" character varying(512) DEFAULT NULL::character varying,
    permission character varying(64)[] DEFAULT '{}'::character varying[] NOT NULL,
    fetched boolean DEFAULT false NOT NULL
);


ALTER TABLE public.access_token OWNER TO misskey;

--
-- Name: COLUMN access_token."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.access_token."createdAt" IS 'The created date of the AccessToken.';


--
-- Name: ad; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.ad (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "expiresAt" timestamp with time zone NOT NULL,
    place character varying(32) NOT NULL,
    priority character varying(32) NOT NULL,
    url character varying(1024) NOT NULL,
    "imageUrl" character varying(1024) NOT NULL,
    memo character varying(8192) NOT NULL,
    ratio integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.ad OWNER TO misskey;

--
-- Name: COLUMN ad."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.ad."createdAt" IS 'The created date of the Ad.';


--
-- Name: COLUMN ad."expiresAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.ad."expiresAt" IS 'The expired date of the Ad.';


--
-- Name: announcement; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.announcement (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    text character varying(8192) NOT NULL,
    title character varying(256) NOT NULL,
    "imageUrl" character varying(1024),
    "updatedAt" timestamp with time zone
);


ALTER TABLE public.announcement OWNER TO misskey;

--
-- Name: COLUMN announcement."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.announcement."createdAt" IS 'The created date of the Announcement.';


--
-- Name: COLUMN announcement."updatedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.announcement."updatedAt" IS 'The updated date of the Announcement.';


--
-- Name: announcement_read; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.announcement_read (
    id character varying(32) NOT NULL,
    "userId" character varying(32) NOT NULL,
    "announcementId" character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL
);


ALTER TABLE public.announcement_read OWNER TO misskey;

--
-- Name: COLUMN announcement_read."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.announcement_read."createdAt" IS 'The created date of the AnnouncementRead.';


--
-- Name: antenna; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.antenna (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    name character varying(128) NOT NULL,
    src public.antenna_src_enum NOT NULL,
    "userListId" character varying(32),
    keywords jsonb DEFAULT '[]'::jsonb NOT NULL,
    "withFile" boolean NOT NULL,
    expression character varying(2048),
    notify boolean NOT NULL,
    "caseSensitive" boolean DEFAULT false NOT NULL,
    "withReplies" boolean DEFAULT false NOT NULL,
    "userGroupJoiningId" character varying(32),
    users character varying(1024)[] DEFAULT '{}'::character varying[] NOT NULL,
    "excludeKeywords" jsonb DEFAULT '[]'::jsonb NOT NULL
);


ALTER TABLE public.antenna OWNER TO misskey;

--
-- Name: COLUMN antenna."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.antenna."createdAt" IS 'The created date of the Antenna.';


--
-- Name: COLUMN antenna."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.antenna."userId" IS 'The owner ID.';


--
-- Name: COLUMN antenna.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.antenna.name IS 'The name of the Antenna.';


--
-- Name: antenna_note; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.antenna_note (
    id character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    "antennaId" character varying(32) NOT NULL,
    read boolean DEFAULT false NOT NULL
);


ALTER TABLE public.antenna_note OWNER TO misskey;

--
-- Name: COLUMN antenna_note."noteId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.antenna_note."noteId" IS 'The note ID.';


--
-- Name: COLUMN antenna_note."antennaId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.antenna_note."antennaId" IS 'The antenna ID.';


--
-- Name: app; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.app (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32),
    secret character varying(64) NOT NULL,
    name character varying(128) NOT NULL,
    description character varying(512) NOT NULL,
    permission character varying(64)[] NOT NULL,
    "callbackUrl" character varying(512)
);


ALTER TABLE public.app OWNER TO misskey;

--
-- Name: COLUMN app."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app."createdAt" IS 'The created date of the App.';


--
-- Name: COLUMN app."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app."userId" IS 'The owner ID.';


--
-- Name: COLUMN app.secret; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app.secret IS 'The secret key of the App.';


--
-- Name: COLUMN app.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app.name IS 'The name of the App.';


--
-- Name: COLUMN app.description; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app.description IS 'The description of the App.';


--
-- Name: COLUMN app.permission; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app.permission IS 'The permission of the App.';


--
-- Name: COLUMN app."callbackUrl"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.app."callbackUrl" IS 'The callbackUrl of the App.';


--
-- Name: attestation_challenge; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.attestation_challenge (
    id character varying(32) NOT NULL,
    "userId" character varying(32) NOT NULL,
    challenge character varying(64) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "registrationChallenge" boolean DEFAULT false NOT NULL
);


ALTER TABLE public.attestation_challenge OWNER TO misskey;

--
-- Name: COLUMN attestation_challenge.challenge; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.attestation_challenge.challenge IS 'Hex-encoded sha256 hash of the challenge.';


--
-- Name: COLUMN attestation_challenge."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.attestation_challenge."createdAt" IS 'The date challenge was created for expiry purposes.';


--
-- Name: COLUMN attestation_challenge."registrationChallenge"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.attestation_challenge."registrationChallenge" IS 'Indicates that the challenge is only for registration purposes if true to prevent the challenge for being used as authentication.';


--
-- Name: auth_session; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.auth_session (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    token character varying(128) NOT NULL,
    "userId" character varying(32),
    "appId" character varying(32) NOT NULL
);


ALTER TABLE public.auth_session OWNER TO misskey;

--
-- Name: COLUMN auth_session."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.auth_session."createdAt" IS 'The created date of the AuthSession.';


--
-- Name: blocking; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.blocking (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "blockeeId" character varying(32) NOT NULL,
    "blockerId" character varying(32) NOT NULL
);


ALTER TABLE public.blocking OWNER TO misskey;

--
-- Name: COLUMN blocking."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.blocking."createdAt" IS 'The created date of the Blocking.';


--
-- Name: COLUMN blocking."blockeeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.blocking."blockeeId" IS 'The blockee user ID.';


--
-- Name: COLUMN blocking."blockerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.blocking."blockerId" IS 'The blocker user ID.';


--
-- Name: channel; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.channel (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "lastNotedAt" timestamp with time zone,
    "userId" character varying(32),
    name character varying(128) NOT NULL,
    description character varying(2048),
    "bannerId" character varying(32),
    "notesCount" integer DEFAULT 0 NOT NULL,
    "usersCount" integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.channel OWNER TO misskey;

--
-- Name: COLUMN channel."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel."createdAt" IS 'The created date of the Channel.';


--
-- Name: COLUMN channel."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel."userId" IS 'The owner ID.';


--
-- Name: COLUMN channel.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel.name IS 'The name of the Channel.';


--
-- Name: COLUMN channel.description; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel.description IS 'The description of the Channel.';


--
-- Name: COLUMN channel."bannerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel."bannerId" IS 'The ID of banner Channel.';


--
-- Name: COLUMN channel."notesCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel."notesCount" IS 'The count of notes.';


--
-- Name: COLUMN channel."usersCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel."usersCount" IS 'The count of users.';


--
-- Name: channel_following; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.channel_following (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "followeeId" character varying(32) NOT NULL,
    "followerId" character varying(32) NOT NULL
);


ALTER TABLE public.channel_following OWNER TO misskey;

--
-- Name: COLUMN channel_following."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel_following."createdAt" IS 'The created date of the ChannelFollowing.';


--
-- Name: COLUMN channel_following."followeeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel_following."followeeId" IS 'The followee channel ID.';


--
-- Name: COLUMN channel_following."followerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel_following."followerId" IS 'The follower user ID.';


--
-- Name: channel_note_pining; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.channel_note_pining (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "channelId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL
);


ALTER TABLE public.channel_note_pining OWNER TO misskey;

--
-- Name: COLUMN channel_note_pining."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.channel_note_pining."createdAt" IS 'The created date of the ChannelNotePining.';


--
-- Name: clip; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.clip (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    name character varying(128) NOT NULL,
    "isPublic" boolean DEFAULT false NOT NULL,
    description character varying(2048) DEFAULT NULL::character varying
);


ALTER TABLE public.clip OWNER TO misskey;

--
-- Name: COLUMN clip."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip."createdAt" IS 'The created date of the Clip.';


--
-- Name: COLUMN clip."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip."userId" IS 'The owner ID.';


--
-- Name: COLUMN clip.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip.name IS 'The name of the Clip.';


--
-- Name: COLUMN clip.description; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip.description IS 'The description of the Clip.';


--
-- Name: clip_note; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.clip_note (
    id character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    "clipId" character varying(32) NOT NULL
);


ALTER TABLE public.clip_note OWNER TO misskey;

--
-- Name: COLUMN clip_note."noteId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip_note."noteId" IS 'The note ID.';


--
-- Name: COLUMN clip_note."clipId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.clip_note."clipId" IS 'The clip ID.';


--
-- Name: drive_file; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.drive_file (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32),
    "userHost" character varying(128),
    md5 character varying(32) NOT NULL,
    name character varying(256) NOT NULL,
    type character varying(128) NOT NULL,
    size integer NOT NULL,
    comment character varying(512),
    properties jsonb DEFAULT '{}'::jsonb NOT NULL,
    "storedInternal" boolean NOT NULL,
    url character varying(512) NOT NULL,
    "thumbnailUrl" character varying(512),
    "webpublicUrl" character varying(512),
    "accessKey" character varying(256),
    "thumbnailAccessKey" character varying(256),
    "webpublicAccessKey" character varying(256),
    uri character varying(512),
    src character varying(512),
    "folderId" character varying(32),
    "isSensitive" boolean DEFAULT false NOT NULL,
    "isLink" boolean DEFAULT false NOT NULL,
    blurhash character varying(128)
);


ALTER TABLE public.drive_file OWNER TO misskey;

--
-- Name: COLUMN drive_file."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."createdAt" IS 'The created date of the DriveFile.';


--
-- Name: COLUMN drive_file."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."userId" IS 'The owner ID.';


--
-- Name: COLUMN drive_file."userHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."userHost" IS 'The host of owner. It will be null if the user in local.';


--
-- Name: COLUMN drive_file.md5; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.md5 IS 'The MD5 hash of the DriveFile.';


--
-- Name: COLUMN drive_file.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.name IS 'The file name of the DriveFile.';


--
-- Name: COLUMN drive_file.type; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.type IS 'The content type (MIME) of the DriveFile.';


--
-- Name: COLUMN drive_file.size; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.size IS 'The file size (bytes) of the DriveFile.';


--
-- Name: COLUMN drive_file.comment; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.comment IS 'The comment of the DriveFile.';


--
-- Name: COLUMN drive_file.properties; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.properties IS 'The any properties of the DriveFile. For example, it includes image width/height.';


--
-- Name: COLUMN drive_file.url; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.url IS 'The URL of the DriveFile.';


--
-- Name: COLUMN drive_file."thumbnailUrl"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."thumbnailUrl" IS 'The URL of the thumbnail of the DriveFile.';


--
-- Name: COLUMN drive_file."webpublicUrl"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."webpublicUrl" IS 'The URL of the webpublic of the DriveFile.';


--
-- Name: COLUMN drive_file.uri; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.uri IS 'The URI of the DriveFile. it will be null when the DriveFile is local.';


--
-- Name: COLUMN drive_file."folderId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."folderId" IS 'The parent folder ID. If null, it means the DriveFile is located in root.';


--
-- Name: COLUMN drive_file."isSensitive"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."isSensitive" IS 'Whether the DriveFile is NSFW.';


--
-- Name: COLUMN drive_file."isLink"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file."isLink" IS 'Whether the DriveFile is direct link to remote server.';


--
-- Name: COLUMN drive_file.blurhash; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_file.blurhash IS 'The BlurHash string.';


--
-- Name: drive_folder; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.drive_folder (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    name character varying(128) NOT NULL,
    "userId" character varying(32),
    "parentId" character varying(32)
);


ALTER TABLE public.drive_folder OWNER TO misskey;

--
-- Name: COLUMN drive_folder."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_folder."createdAt" IS 'The created date of the DriveFolder.';


--
-- Name: COLUMN drive_folder.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_folder.name IS 'The name of the DriveFolder.';


--
-- Name: COLUMN drive_folder."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_folder."userId" IS 'The owner ID.';


--
-- Name: COLUMN drive_folder."parentId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.drive_folder."parentId" IS 'The parent folder ID. If null, it means the DriveFolder is located in root.';


--
-- Name: emoji; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.emoji (
    id character varying(32) NOT NULL,
    "updatedAt" timestamp with time zone,
    name character varying(128) NOT NULL,
    host character varying(128),
    url character varying(512) NOT NULL,
    uri character varying(512),
    type character varying(64),
    aliases character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    category character varying(128)
);


ALTER TABLE public.emoji OWNER TO misskey;

--
-- Name: follow_request; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.follow_request (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "followeeId" character varying(32) NOT NULL,
    "followerId" character varying(32) NOT NULL,
    "requestId" character varying(128),
    "followerHost" character varying(128),
    "followerInbox" character varying(512),
    "followerSharedInbox" character varying(512),
    "followeeHost" character varying(128),
    "followeeInbox" character varying(512),
    "followeeSharedInbox" character varying(512)
);


ALTER TABLE public.follow_request OWNER TO misskey;

--
-- Name: COLUMN follow_request."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."createdAt" IS 'The created date of the FollowRequest.';


--
-- Name: COLUMN follow_request."followeeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followeeId" IS 'The followee user ID.';


--
-- Name: COLUMN follow_request."followerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followerId" IS 'The follower user ID.';


--
-- Name: COLUMN follow_request."requestId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."requestId" IS 'id of Follow Activity.';


--
-- Name: COLUMN follow_request."followerHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followerHost" IS '[Denormalized]';


--
-- Name: COLUMN follow_request."followerInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followerInbox" IS '[Denormalized]';


--
-- Name: COLUMN follow_request."followerSharedInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followerSharedInbox" IS '[Denormalized]';


--
-- Name: COLUMN follow_request."followeeHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followeeHost" IS '[Denormalized]';


--
-- Name: COLUMN follow_request."followeeInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followeeInbox" IS '[Denormalized]';


--
-- Name: COLUMN follow_request."followeeSharedInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.follow_request."followeeSharedInbox" IS '[Denormalized]';


--
-- Name: following; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.following (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "followeeId" character varying(32) NOT NULL,
    "followerId" character varying(32) NOT NULL,
    "followerHost" character varying(128),
    "followerInbox" character varying(512),
    "followerSharedInbox" character varying(512),
    "followeeHost" character varying(128),
    "followeeInbox" character varying(512),
    "followeeSharedInbox" character varying(512)
);


ALTER TABLE public.following OWNER TO misskey;

--
-- Name: COLUMN following."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."createdAt" IS 'The created date of the Following.';


--
-- Name: COLUMN following."followeeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followeeId" IS 'The followee user ID.';


--
-- Name: COLUMN following."followerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followerId" IS 'The follower user ID.';


--
-- Name: COLUMN following."followerHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followerHost" IS '[Denormalized]';


--
-- Name: COLUMN following."followerInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followerInbox" IS '[Denormalized]';


--
-- Name: COLUMN following."followerSharedInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followerSharedInbox" IS '[Denormalized]';


--
-- Name: COLUMN following."followeeHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followeeHost" IS '[Denormalized]';


--
-- Name: COLUMN following."followeeInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followeeInbox" IS '[Denormalized]';


--
-- Name: COLUMN following."followeeSharedInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.following."followeeSharedInbox" IS '[Denormalized]';


--
-- Name: gallery_like; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.gallery_like (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "postId" character varying(32) NOT NULL
);


ALTER TABLE public.gallery_like OWNER TO misskey;

--
-- Name: gallery_post; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.gallery_post (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    title character varying(256) NOT NULL,
    description character varying(2048),
    "userId" character varying(32) NOT NULL,
    "fileIds" character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    "isSensitive" boolean DEFAULT false NOT NULL,
    "likedCount" integer DEFAULT 0 NOT NULL,
    tags character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL
);


ALTER TABLE public.gallery_post OWNER TO misskey;

--
-- Name: COLUMN gallery_post."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.gallery_post."createdAt" IS 'The created date of the GalleryPost.';


--
-- Name: COLUMN gallery_post."updatedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.gallery_post."updatedAt" IS 'The updated date of the GalleryPost.';


--
-- Name: COLUMN gallery_post."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.gallery_post."userId" IS 'The ID of author.';


--
-- Name: COLUMN gallery_post."isSensitive"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.gallery_post."isSensitive" IS 'Whether the post is sensitive.';


--
-- Name: hashtag; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.hashtag (
    id character varying(32) NOT NULL,
    name character varying(128) NOT NULL,
    "mentionedUserIds" character varying(32)[] NOT NULL,
    "mentionedUsersCount" integer DEFAULT 0 NOT NULL,
    "mentionedLocalUserIds" character varying(32)[] NOT NULL,
    "mentionedLocalUsersCount" integer DEFAULT 0 NOT NULL,
    "mentionedRemoteUserIds" character varying(32)[] NOT NULL,
    "mentionedRemoteUsersCount" integer DEFAULT 0 NOT NULL,
    "attachedUserIds" character varying(32)[] NOT NULL,
    "attachedUsersCount" integer DEFAULT 0 NOT NULL,
    "attachedLocalUserIds" character varying(32)[] NOT NULL,
    "attachedLocalUsersCount" integer DEFAULT 0 NOT NULL,
    "attachedRemoteUserIds" character varying(32)[] NOT NULL,
    "attachedRemoteUsersCount" integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.hashtag OWNER TO misskey;

--
-- Name: instance; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.instance (
    id character varying(32) NOT NULL,
    "caughtAt" timestamp with time zone NOT NULL,
    host character varying(128) NOT NULL,
    "usersCount" integer DEFAULT 0 NOT NULL,
    "notesCount" integer DEFAULT 0 NOT NULL,
    "followingCount" integer DEFAULT 0 NOT NULL,
    "followersCount" integer DEFAULT 0 NOT NULL,
    "driveUsage" bigint DEFAULT 0 NOT NULL,
    "driveFiles" integer DEFAULT 0 NOT NULL,
    "latestRequestSentAt" timestamp with time zone,
    "latestStatus" integer,
    "latestRequestReceivedAt" timestamp with time zone,
    "lastCommunicatedAt" timestamp with time zone NOT NULL,
    "isNotResponding" boolean DEFAULT false NOT NULL,
    "softwareName" character varying(64) DEFAULT NULL::character varying,
    "softwareVersion" character varying(64) DEFAULT NULL::character varying,
    "openRegistrations" boolean,
    name character varying(256) DEFAULT NULL::character varying,
    description character varying(4096) DEFAULT NULL::character varying,
    "maintainerName" character varying(128) DEFAULT NULL::character varying,
    "maintainerEmail" character varying(256) DEFAULT NULL::character varying,
    "infoUpdatedAt" timestamp with time zone,
    "isSuspended" boolean DEFAULT false NOT NULL,
    "iconUrl" character varying(256) DEFAULT NULL::character varying,
    "themeColor" character varying(64) DEFAULT NULL::character varying,
    "faviconUrl" character varying(256) DEFAULT NULL::character varying
);


ALTER TABLE public.instance OWNER TO misskey;

--
-- Name: COLUMN instance."caughtAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.instance."caughtAt" IS 'The caught date of the Instance.';


--
-- Name: COLUMN instance.host; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.instance.host IS 'The host of the Instance.';


--
-- Name: COLUMN instance."usersCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.instance."usersCount" IS 'The count of the users of the Instance.';


--
-- Name: COLUMN instance."notesCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.instance."notesCount" IS 'The count of the notes of the Instance.';


--
-- Name: COLUMN instance."softwareName"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.instance."softwareName" IS 'The software of the Instance.';


--
-- Name: messaging_message; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.messaging_message (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "recipientId" character varying(32),
    text character varying(4096),
    "isRead" boolean DEFAULT false NOT NULL,
    "fileId" character varying(32),
    "groupId" character varying(32),
    reads character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    uri character varying(512)
);


ALTER TABLE public.messaging_message OWNER TO misskey;

--
-- Name: COLUMN messaging_message."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.messaging_message."createdAt" IS 'The created date of the MessagingMessage.';


--
-- Name: COLUMN messaging_message."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.messaging_message."userId" IS 'The sender user ID.';


--
-- Name: COLUMN messaging_message."recipientId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.messaging_message."recipientId" IS 'The recipient user ID.';


--
-- Name: COLUMN messaging_message."groupId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.messaging_message."groupId" IS 'The recipient group ID.';


--
-- Name: meta; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.meta (
    id character varying(32) NOT NULL,
    name character varying(128),
    description character varying(1024),
    "maintainerName" character varying(128),
    "maintainerEmail" character varying(128),
    "disableRegistration" boolean DEFAULT false NOT NULL,
    "disableLocalTimeline" boolean DEFAULT false NOT NULL,
    "disableGlobalTimeline" boolean DEFAULT false NOT NULL,
    "useStarForReactionFallback" boolean DEFAULT false NOT NULL,
    langs character varying(64)[] DEFAULT '{}'::character varying[] NOT NULL,
    "hiddenTags" character varying(256)[] DEFAULT '{}'::character varying[] NOT NULL,
    "blockedHosts" character varying(256)[] DEFAULT '{}'::character varying[] NOT NULL,
    "mascotImageUrl" character varying(512) DEFAULT '/assets/ai.png'::character varying,
    "bannerUrl" character varying(512),
    "errorImageUrl" character varying(512) DEFAULT 'https://xn--931a.moe/aiart/yubitun.png'::character varying,
    "iconUrl" character varying(512),
    "cacheRemoteFiles" boolean DEFAULT true NOT NULL,
    "enableRecaptcha" boolean DEFAULT false NOT NULL,
    "recaptchaSiteKey" character varying(64),
    "recaptchaSecretKey" character varying(64),
    "localDriveCapacityMb" integer DEFAULT 1024 NOT NULL,
    "remoteDriveCapacityMb" integer DEFAULT 32 NOT NULL,
    "maxNoteTextLength" integer DEFAULT 500 NOT NULL,
    "summalyProxy" character varying(128),
    "enableEmail" boolean DEFAULT false NOT NULL,
    email character varying(128),
    "smtpSecure" boolean DEFAULT false NOT NULL,
    "smtpHost" character varying(128),
    "smtpPort" integer,
    "smtpUser" character varying(128),
    "smtpPass" character varying(128),
    "enableServiceWorker" boolean DEFAULT false NOT NULL,
    "swPublicKey" character varying(128),
    "swPrivateKey" character varying(128),
    "enableTwitterIntegration" boolean DEFAULT false NOT NULL,
    "twitterConsumerKey" character varying(128),
    "twitterConsumerSecret" character varying(128),
    "enableGithubIntegration" boolean DEFAULT false NOT NULL,
    "githubClientId" character varying(128),
    "githubClientSecret" character varying(128),
    "enableDiscordIntegration" boolean DEFAULT false NOT NULL,
    "discordClientId" character varying(128),
    "discordClientSecret" character varying(128),
    "pinnedUsers" character varying(256)[] DEFAULT '{}'::character varying[] NOT NULL,
    "ToSUrl" character varying(512),
    "repositoryUrl" character varying(512) DEFAULT 'https://github.com/misskey-dev/misskey'::character varying NOT NULL,
    "feedbackUrl" character varying(512) DEFAULT 'https://github.com/misskey-dev/misskey/issues/new'::character varying,
    "useObjectStorage" boolean DEFAULT false NOT NULL,
    "objectStorageBucket" character varying(512),
    "objectStoragePrefix" character varying(512),
    "objectStorageBaseUrl" character varying(512),
    "objectStorageEndpoint" character varying(512),
    "objectStorageRegion" character varying(512),
    "objectStorageAccessKey" character varying(512),
    "objectStorageSecretKey" character varying(512),
    "objectStoragePort" integer,
    "objectStorageUseSSL" boolean DEFAULT true NOT NULL,
    "proxyRemoteFiles" boolean DEFAULT false NOT NULL,
    "proxyAccountId" character varying(32),
    "objectStorageUseProxy" boolean DEFAULT true NOT NULL,
    "enableHcaptcha" boolean DEFAULT false NOT NULL,
    "hcaptchaSiteKey" character varying(64),
    "hcaptchaSecretKey" character varying(64),
    "objectStorageSetPublicRead" boolean DEFAULT false NOT NULL,
    "pinnedPages" character varying(512)[] DEFAULT '{/featured,/channels,/explore,/pages,/about-misskey}'::character varying[] NOT NULL,
    "backgroundImageUrl" character varying(512),
    "logoImageUrl" character varying(512),
    "pinnedClipId" character varying(32),
    "objectStorageS3ForcePathStyle" boolean DEFAULT true NOT NULL,
    "deeplAuthKey" character varying(128),
    "deeplIsPro" boolean DEFAULT false NOT NULL,
    "emailRequiredForSignup" boolean DEFAULT false NOT NULL
);


ALTER TABLE public.meta OWNER TO misskey;

--
-- Name: COLUMN meta."localDriveCapacityMb"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.meta."localDriveCapacityMb" IS 'Drive capacity of a local user (MB)';


--
-- Name: COLUMN meta."remoteDriveCapacityMb"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.meta."remoteDriveCapacityMb" IS 'Drive capacity of a remote user (MB)';


--
-- Name: COLUMN meta."maxNoteTextLength"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.meta."maxNoteTextLength" IS 'Max allowed note text length in characters';


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.migrations (
    id integer NOT NULL,
    "timestamp" bigint NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.migrations OWNER TO misskey;

--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: misskey
--

CREATE SEQUENCE public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.migrations_id_seq OWNER TO misskey;

--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: misskey
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: moderation_log; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.moderation_log (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    type character varying(128) NOT NULL,
    info jsonb NOT NULL
);


ALTER TABLE public.moderation_log OWNER TO misskey;

--
-- Name: COLUMN moderation_log."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.moderation_log."createdAt" IS 'The created date of the ModerationLog.';


--
-- Name: muted_note; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.muted_note (
    id character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    "userId" character varying(32) NOT NULL,
    reason public.muted_note_reason_enum NOT NULL
);


ALTER TABLE public.muted_note OWNER TO misskey;

--
-- Name: COLUMN muted_note."noteId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muted_note."noteId" IS 'The note ID.';


--
-- Name: COLUMN muted_note."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muted_note."userId" IS 'The user ID.';


--
-- Name: COLUMN muted_note.reason; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muted_note.reason IS 'The reason of the MutedNote.';


--
-- Name: muting; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.muting (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "muteeId" character varying(32) NOT NULL,
    "muterId" character varying(32) NOT NULL
);


ALTER TABLE public.muting OWNER TO misskey;

--
-- Name: COLUMN muting."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muting."createdAt" IS 'The created date of the Muting.';


--
-- Name: COLUMN muting."muteeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muting."muteeId" IS 'The mutee user ID.';


--
-- Name: COLUMN muting."muterId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.muting."muterId" IS 'The muter user ID.';


--
-- Name: note; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "replyId" character varying(32),
    "renoteId" character varying(32),
    text text,
    name character varying(256),
    cw character varying(512),
    "userId" character varying(32) NOT NULL,
    "localOnly" boolean DEFAULT false NOT NULL,
    "renoteCount" smallint DEFAULT 0 NOT NULL,
    "repliesCount" smallint DEFAULT 0 NOT NULL,
    reactions jsonb DEFAULT '{}'::jsonb NOT NULL,
    visibility public.note_visibility_enum NOT NULL,
    uri character varying(512),
    score integer DEFAULT 0 NOT NULL,
    "fileIds" character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    "attachedFileTypes" character varying(256)[] DEFAULT '{}'::character varying[] NOT NULL,
    "visibleUserIds" character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    mentions character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    "mentionedRemoteUsers" text DEFAULT '[]'::text NOT NULL,
    emojis character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    tags character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    "hasPoll" boolean DEFAULT false NOT NULL,
    "userHost" character varying(128),
    "replyUserId" character varying(32),
    "replyUserHost" character varying(128),
    "renoteUserId" character varying(32),
    "renoteUserHost" character varying(128),
    url character varying(512),
    "channelId" character varying(32) DEFAULT NULL::character varying,
    "threadId" character varying(256)
);


ALTER TABLE public.note OWNER TO misskey;

--
-- Name: COLUMN note."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."createdAt" IS 'The created date of the Note.';


--
-- Name: COLUMN note."replyId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."replyId" IS 'The ID of reply target.';


--
-- Name: COLUMN note."renoteId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."renoteId" IS 'The ID of renote target.';


--
-- Name: COLUMN note."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."userId" IS 'The ID of author.';


--
-- Name: COLUMN note.uri; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note.uri IS 'The URI of a note. it will be null when the note is local.';


--
-- Name: COLUMN note."userHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."userHost" IS '[Denormalized]';


--
-- Name: COLUMN note."replyUserId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."replyUserId" IS '[Denormalized]';


--
-- Name: COLUMN note."replyUserHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."replyUserHost" IS '[Denormalized]';


--
-- Name: COLUMN note."renoteUserId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."renoteUserId" IS '[Denormalized]';


--
-- Name: COLUMN note."renoteUserHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."renoteUserHost" IS '[Denormalized]';


--
-- Name: COLUMN note.url; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note.url IS 'The human readable url of a note. it will be null when the note is local.';


--
-- Name: COLUMN note."channelId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note."channelId" IS 'The ID of source channel.';


--
-- Name: note_favorite; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note_favorite (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL
);


ALTER TABLE public.note_favorite OWNER TO misskey;

--
-- Name: COLUMN note_favorite."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_favorite."createdAt" IS 'The created date of the NoteFavorite.';


--
-- Name: note_reaction; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note_reaction (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    reaction character varying(260) NOT NULL
);


ALTER TABLE public.note_reaction OWNER TO misskey;

--
-- Name: COLUMN note_reaction."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_reaction."createdAt" IS 'The created date of the NoteReaction.';


--
-- Name: note_thread_muting; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note_thread_muting (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "threadId" character varying(256) NOT NULL
);


ALTER TABLE public.note_thread_muting OWNER TO misskey;

--
-- Name: note_unread; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note_unread (
    id character varying(32) NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    "noteUserId" character varying(32) NOT NULL,
    "isSpecified" boolean NOT NULL,
    "isMentioned" boolean NOT NULL,
    "noteChannelId" character varying(32)
);


ALTER TABLE public.note_unread OWNER TO misskey;

--
-- Name: COLUMN note_unread."noteUserId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_unread."noteUserId" IS '[Denormalized]';


--
-- Name: COLUMN note_unread."noteChannelId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_unread."noteChannelId" IS '[Denormalized]';


--
-- Name: note_watching; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.note_watching (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    "noteUserId" character varying(32) NOT NULL
);


ALTER TABLE public.note_watching OWNER TO misskey;

--
-- Name: COLUMN note_watching."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_watching."createdAt" IS 'The created date of the NoteWatching.';


--
-- Name: COLUMN note_watching."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_watching."userId" IS 'The watcher ID.';


--
-- Name: COLUMN note_watching."noteId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_watching."noteId" IS 'The target Note ID.';


--
-- Name: COLUMN note_watching."noteUserId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.note_watching."noteUserId" IS '[Denormalized]';


--
-- Name: notification; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.notification (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "notifieeId" character varying(32) NOT NULL,
    "notifierId" character varying(32),
    "isRead" boolean DEFAULT false NOT NULL,
    "noteId" character varying(32),
    reaction character varying(128),
    choice integer,
    "followRequestId" character varying(32),
    type public.notification_type_enum NOT NULL,
    "userGroupInvitationId" character varying(32),
    "customBody" character varying(2048),
    "customHeader" character varying(256),
    "customIcon" character varying(1024),
    "appAccessTokenId" character varying(32)
);


ALTER TABLE public.notification OWNER TO misskey;

--
-- Name: COLUMN notification."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.notification."createdAt" IS 'The created date of the Notification.';


--
-- Name: COLUMN notification."notifieeId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.notification."notifieeId" IS 'The ID of recipient user of the Notification.';


--
-- Name: COLUMN notification."notifierId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.notification."notifierId" IS 'The ID of sender user of the Notification.';


--
-- Name: COLUMN notification."isRead"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.notification."isRead" IS 'Whether the Notification is read.';


--
-- Name: COLUMN notification.type; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.notification.type IS 'The type of the Notification.';


--
-- Name: page; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.page (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    title character varying(256) NOT NULL,
    name character varying(256) NOT NULL,
    summary character varying(256),
    "alignCenter" boolean NOT NULL,
    font character varying(32) NOT NULL,
    "userId" character varying(32) NOT NULL,
    "eyeCatchingImageId" character varying(32),
    content jsonb DEFAULT '[]'::jsonb NOT NULL,
    variables jsonb DEFAULT '[]'::jsonb NOT NULL,
    visibility public.page_visibility_enum NOT NULL,
    "visibleUserIds" character varying(32)[] DEFAULT '{}'::character varying[] NOT NULL,
    "likedCount" integer DEFAULT 0 NOT NULL,
    "hideTitleWhenPinned" boolean DEFAULT false NOT NULL,
    script character varying(16384) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.page OWNER TO misskey;

--
-- Name: COLUMN page."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.page."createdAt" IS 'The created date of the Page.';


--
-- Name: COLUMN page."updatedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.page."updatedAt" IS 'The updated date of the Page.';


--
-- Name: COLUMN page."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.page."userId" IS 'The ID of author.';


--
-- Name: page_like; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.page_like (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "pageId" character varying(32) NOT NULL
);


ALTER TABLE public.page_like OWNER TO misskey;

--
-- Name: password_reset_request; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.password_reset_request (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    token character varying(256) NOT NULL,
    "userId" character varying(32) NOT NULL
);


ALTER TABLE public.password_reset_request OWNER TO misskey;

--
-- Name: poll; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.poll (
    "noteId" character varying(32) NOT NULL,
    "expiresAt" timestamp with time zone,
    multiple boolean NOT NULL,
    choices character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    votes integer[] NOT NULL,
    "noteVisibility" public.poll_notevisibility_enum NOT NULL,
    "userId" character varying(32) NOT NULL,
    "userHost" character varying(128)
);


ALTER TABLE public.poll OWNER TO misskey;

--
-- Name: COLUMN poll."noteVisibility"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.poll."noteVisibility" IS '[Denormalized]';


--
-- Name: COLUMN poll."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.poll."userId" IS '[Denormalized]';


--
-- Name: COLUMN poll."userHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.poll."userHost" IS '[Denormalized]';


--
-- Name: poll_vote; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.poll_vote (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL,
    choice integer NOT NULL
);


ALTER TABLE public.poll_vote OWNER TO misskey;

--
-- Name: COLUMN poll_vote."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.poll_vote."createdAt" IS 'The created date of the PollVote.';


--
-- Name: promo_note; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.promo_note (
    "noteId" character varying(32) NOT NULL,
    "expiresAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL
);


ALTER TABLE public.promo_note OWNER TO misskey;

--
-- Name: COLUMN promo_note."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.promo_note."userId" IS '[Denormalized]';


--
-- Name: promo_read; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.promo_read (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL
);


ALTER TABLE public.promo_read OWNER TO misskey;

--
-- Name: COLUMN promo_read."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.promo_read."createdAt" IS 'The created date of the PromoRead.';


--
-- Name: registration_ticket; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.registration_ticket (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    code character varying(64) NOT NULL
);


ALTER TABLE public.registration_ticket OWNER TO misskey;

--
-- Name: registry_item; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.registry_item (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    key character varying(1024) NOT NULL,
    scope character varying(1024)[] DEFAULT '{}'::character varying[] NOT NULL,
    domain character varying(512),
    value jsonb DEFAULT '{}'::jsonb
);


ALTER TABLE public.registry_item OWNER TO misskey;

--
-- Name: COLUMN registry_item."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.registry_item."createdAt" IS 'The created date of the RegistryItem.';


--
-- Name: COLUMN registry_item."updatedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.registry_item."updatedAt" IS 'The updated date of the RegistryItem.';


--
-- Name: COLUMN registry_item."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.registry_item."userId" IS 'The owner ID.';


--
-- Name: COLUMN registry_item.key; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.registry_item.key IS 'The key of the RegistryItem.';


--
-- Name: COLUMN registry_item.value; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.registry_item.value IS 'The value of the RegistryItem.';


--
-- Name: relay; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.relay (
    id character varying(32) NOT NULL,
    inbox character varying(512) NOT NULL,
    status public.relay_status_enum NOT NULL
);


ALTER TABLE public.relay OWNER TO misskey;

--
-- Name: reversi_game; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.reversi_game (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "startedAt" timestamp with time zone,
    "user1Id" character varying(32) NOT NULL,
    "user2Id" character varying(32) NOT NULL,
    "user1Accepted" boolean DEFAULT false NOT NULL,
    "user2Accepted" boolean DEFAULT false NOT NULL,
    black integer,
    "isStarted" boolean DEFAULT false NOT NULL,
    "isEnded" boolean DEFAULT false NOT NULL,
    "winnerId" character varying(32),
    surrendered character varying(32),
    logs jsonb DEFAULT '[]'::jsonb NOT NULL,
    map character varying(64)[] NOT NULL,
    bw character varying(32) NOT NULL,
    "isLlotheo" boolean DEFAULT false NOT NULL,
    "canPutEverywhere" boolean DEFAULT false NOT NULL,
    "loopedBoard" boolean DEFAULT false NOT NULL,
    form1 jsonb,
    form2 jsonb,
    crc32 character varying(32)
);


ALTER TABLE public.reversi_game OWNER TO misskey;

--
-- Name: COLUMN reversi_game."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.reversi_game."createdAt" IS 'The created date of the ReversiGame.';


--
-- Name: COLUMN reversi_game."startedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.reversi_game."startedAt" IS 'The started date of the ReversiGame.';


--
-- Name: reversi_matching; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.reversi_matching (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "parentId" character varying(32) NOT NULL,
    "childId" character varying(32) NOT NULL
);


ALTER TABLE public.reversi_matching OWNER TO misskey;

--
-- Name: COLUMN reversi_matching."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.reversi_matching."createdAt" IS 'The created date of the ReversiMatching.';


--
-- Name: signin; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.signin (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    ip character varying(128) NOT NULL,
    headers jsonb NOT NULL,
    success boolean NOT NULL
);


ALTER TABLE public.signin OWNER TO misskey;

--
-- Name: COLUMN signin."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.signin."createdAt" IS 'The created date of the Signin.';


--
-- Name: sw_subscription; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.sw_subscription (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    endpoint character varying(512) NOT NULL,
    auth character varying(256) NOT NULL,
    publickey character varying(128) NOT NULL
);


ALTER TABLE public.sw_subscription OWNER TO misskey;

--
-- Name: used_username; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.used_username (
    username character varying(128) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL
);


ALTER TABLE public.used_username OWNER TO misskey;

--
-- Name: user; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public."user" (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone,
    "lastFetchedAt" timestamp with time zone,
    username character varying(128) NOT NULL,
    "usernameLower" character varying(128) NOT NULL,
    name character varying(128),
    "followersCount" integer DEFAULT 0 NOT NULL,
    "followingCount" integer DEFAULT 0 NOT NULL,
    "notesCount" integer DEFAULT 0 NOT NULL,
    "avatarId" character varying(32),
    "bannerId" character varying(32),
    tags character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    "avatarUrl" character varying(512),
    "bannerUrl" character varying(512),
    "isSuspended" boolean DEFAULT false NOT NULL,
    "isSilenced" boolean DEFAULT false NOT NULL,
    "isLocked" boolean DEFAULT false NOT NULL,
    "isBot" boolean DEFAULT false NOT NULL,
    "isCat" boolean DEFAULT false NOT NULL,
    "isAdmin" boolean DEFAULT false NOT NULL,
    "isModerator" boolean DEFAULT false NOT NULL,
    emojis character varying(128)[] DEFAULT '{}'::character varying[] NOT NULL,
    host character varying(128),
    inbox character varying(512),
    "sharedInbox" character varying(512),
    featured character varying(512),
    uri character varying(512),
    token character(16),
    "avatarBlurhash" character varying(128),
    "bannerBlurhash" character varying(128),
    "isExplorable" boolean DEFAULT true NOT NULL,
    "followersUri" character varying(512) DEFAULT NULL::character varying,
    "lastActiveDate" timestamp with time zone,
    "hideOnlineStatus" boolean DEFAULT false NOT NULL,
    "isDeleted" boolean DEFAULT false NOT NULL
);


ALTER TABLE public."user" OWNER TO misskey;

--
-- Name: COLUMN "user"."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."createdAt" IS 'The created date of the User.';


--
-- Name: COLUMN "user"."updatedAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."updatedAt" IS 'The updated date of the User.';


--
-- Name: COLUMN "user".username; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".username IS 'The username of the User.';


--
-- Name: COLUMN "user"."usernameLower"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."usernameLower" IS 'The username (lowercased) of the User.';


--
-- Name: COLUMN "user".name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".name IS 'The name of the User.';


--
-- Name: COLUMN "user"."followersCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."followersCount" IS 'The count of followers.';


--
-- Name: COLUMN "user"."followingCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."followingCount" IS 'The count of following.';


--
-- Name: COLUMN "user"."notesCount"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."notesCount" IS 'The count of notes.';


--
-- Name: COLUMN "user"."avatarId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."avatarId" IS 'The ID of avatar DriveFile.';


--
-- Name: COLUMN "user"."bannerId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."bannerId" IS 'The ID of banner DriveFile.';


--
-- Name: COLUMN "user"."isSuspended"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isSuspended" IS 'Whether the User is suspended.';


--
-- Name: COLUMN "user"."isSilenced"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isSilenced" IS 'Whether the User is silenced.';


--
-- Name: COLUMN "user"."isLocked"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isLocked" IS 'Whether the User is locked.';


--
-- Name: COLUMN "user"."isBot"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isBot" IS 'Whether the User is a bot.';


--
-- Name: COLUMN "user"."isCat"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isCat" IS 'Whether the User is a cat.';


--
-- Name: COLUMN "user"."isAdmin"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isAdmin" IS 'Whether the User is the admin.';


--
-- Name: COLUMN "user"."isModerator"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isModerator" IS 'Whether the User is a moderator.';


--
-- Name: COLUMN "user".host; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".host IS 'The host of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user".inbox; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".inbox IS 'The inbox URL of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user"."sharedInbox"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."sharedInbox" IS 'The sharedInbox URL of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user".featured; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".featured IS 'The featured URL of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user".uri; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".uri IS 'The URI of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user".token; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user".token IS 'The native access token of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user"."isExplorable"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isExplorable" IS 'Whether the User is explorable.';


--
-- Name: COLUMN "user"."followersUri"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."followersUri" IS 'The URI of the user Follower Collection. It will be null if the origin of the user is local.';


--
-- Name: COLUMN "user"."isDeleted"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public."user"."isDeleted" IS 'Whether the User is deleted.';


--
-- Name: user_group; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_group (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    name character varying(256) NOT NULL,
    "userId" character varying(32) NOT NULL,
    "isPrivate" boolean DEFAULT false NOT NULL
);


ALTER TABLE public.user_group OWNER TO misskey;

--
-- Name: COLUMN user_group."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group."createdAt" IS 'The created date of the UserGroup.';


--
-- Name: COLUMN user_group."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group."userId" IS 'The ID of owner.';


--
-- Name: user_group_invitation; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_group_invitation (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "userGroupId" character varying(32) NOT NULL
);


ALTER TABLE public.user_group_invitation OWNER TO misskey;

--
-- Name: COLUMN user_group_invitation."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_invitation."createdAt" IS 'The created date of the UserGroupInvitation.';


--
-- Name: COLUMN user_group_invitation."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_invitation."userId" IS 'The user ID.';


--
-- Name: COLUMN user_group_invitation."userGroupId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_invitation."userGroupId" IS 'The group ID.';


--
-- Name: user_group_invite; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_group_invite (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "userGroupId" character varying(32) NOT NULL
);


ALTER TABLE public.user_group_invite OWNER TO misskey;

--
-- Name: user_group_joining; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_group_joining (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "userGroupId" character varying(32) NOT NULL
);


ALTER TABLE public.user_group_joining OWNER TO misskey;

--
-- Name: COLUMN user_group_joining."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_joining."createdAt" IS 'The created date of the UserGroupJoining.';


--
-- Name: COLUMN user_group_joining."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_joining."userId" IS 'The user ID.';


--
-- Name: COLUMN user_group_joining."userGroupId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_group_joining."userGroupId" IS 'The group ID.';


--
-- Name: user_keypair; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_keypair (
    "userId" character varying(32) NOT NULL,
    "publicKey" character varying(4096) NOT NULL,
    "privateKey" character varying(4096) NOT NULL
);


ALTER TABLE public.user_keypair OWNER TO misskey;

--
-- Name: user_list; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_list (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    name character varying(128) NOT NULL
);


ALTER TABLE public.user_list OWNER TO misskey;

--
-- Name: COLUMN user_list."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list."createdAt" IS 'The created date of the UserList.';


--
-- Name: COLUMN user_list."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list."userId" IS 'The owner ID.';


--
-- Name: COLUMN user_list.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list.name IS 'The name of the UserList.';


--
-- Name: user_list_joining; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_list_joining (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "userListId" character varying(32) NOT NULL
);


ALTER TABLE public.user_list_joining OWNER TO misskey;

--
-- Name: COLUMN user_list_joining."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list_joining."createdAt" IS 'The created date of the UserListJoining.';


--
-- Name: COLUMN user_list_joining."userId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list_joining."userId" IS 'The user ID.';


--
-- Name: COLUMN user_list_joining."userListId"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_list_joining."userListId" IS 'The list ID.';


--
-- Name: user_note_pining; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_note_pining (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "userId" character varying(32) NOT NULL,
    "noteId" character varying(32) NOT NULL
);


ALTER TABLE public.user_note_pining OWNER TO misskey;

--
-- Name: COLUMN user_note_pining."createdAt"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_note_pining."createdAt" IS 'The created date of the UserNotePinings.';


--
-- Name: user_pending; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_pending (
    id character varying(32) NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    code character varying(128) NOT NULL,
    username character varying(128) NOT NULL,
    email character varying(128) NOT NULL,
    password character varying(128) NOT NULL
);


ALTER TABLE public.user_pending OWNER TO misskey;

--
-- Name: user_profile; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_profile (
    "userId" character varying(32) NOT NULL,
    location character varying(128),
    birthday character(10),
    description character varying(2048),
    fields jsonb DEFAULT '[]'::jsonb NOT NULL,
    url character varying(512),
    email character varying(128),
    "emailVerifyCode" character varying(128),
    "emailVerified" boolean DEFAULT false NOT NULL,
    "twoFactorTempSecret" character varying(128),
    "twoFactorSecret" character varying(128),
    "twoFactorEnabled" boolean DEFAULT false NOT NULL,
    password character varying(128),
    "clientData" jsonb DEFAULT '{}'::jsonb NOT NULL,
    "autoAcceptFollowed" boolean DEFAULT false NOT NULL,
    "alwaysMarkNsfw" boolean DEFAULT false NOT NULL,
    "carefulBot" boolean DEFAULT false NOT NULL,
    "userHost" character varying(128),
    "securityKeysAvailable" boolean DEFAULT false NOT NULL,
    "usePasswordLessLogin" boolean DEFAULT false NOT NULL,
    "pinnedPageId" character varying(32),
    room jsonb DEFAULT '{}'::jsonb NOT NULL,
    integrations jsonb DEFAULT '{}'::jsonb NOT NULL,
    "injectFeaturedNote" boolean DEFAULT true NOT NULL,
    "enableWordMute" boolean DEFAULT false NOT NULL,
    "mutedWords" jsonb DEFAULT '[]'::jsonb NOT NULL,
    "mutingNotificationTypes" public.user_profile_mutingnotificationtypes_enum[] DEFAULT '{}'::public.user_profile_mutingnotificationtypes_enum[] NOT NULL,
    "noCrawle" boolean DEFAULT false NOT NULL,
    "receiveAnnouncementEmail" boolean DEFAULT true NOT NULL,
    "emailNotificationTypes" jsonb DEFAULT '["follow", "receiveFollowRequest", "groupInvited"]'::jsonb NOT NULL,
    lang character varying(32),
    "mutedInstances" jsonb DEFAULT '[]'::jsonb NOT NULL,
    "publicReactions" boolean DEFAULT false NOT NULL,
    "ffVisibility" public.user_profile_ffvisibility_enum DEFAULT 'public'::public.user_profile_ffvisibility_enum NOT NULL
);


ALTER TABLE public.user_profile OWNER TO misskey;

--
-- Name: COLUMN user_profile.location; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.location IS 'The location of the User.';


--
-- Name: COLUMN user_profile.birthday; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.birthday IS 'The birthday (YYYY-MM-DD) of the User.';


--
-- Name: COLUMN user_profile.description; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.description IS 'The description (bio) of the User.';


--
-- Name: COLUMN user_profile.url; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.url IS 'Remote URL of the user.';


--
-- Name: COLUMN user_profile.email; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.email IS 'The email address of the User.';


--
-- Name: COLUMN user_profile.password; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.password IS 'The password hash of the User. It will be null if the origin of the user is local.';


--
-- Name: COLUMN user_profile."clientData"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile."clientData" IS 'The client-specific data of the User.';


--
-- Name: COLUMN user_profile."userHost"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile."userHost" IS '[Denormalized]';


--
-- Name: COLUMN user_profile.room; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile.room IS 'The room data of the User.';


--
-- Name: COLUMN user_profile."noCrawle"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile."noCrawle" IS 'Whether reject index by crawler.';


--
-- Name: COLUMN user_profile."mutedInstances"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_profile."mutedInstances" IS 'List of instances muted by the user.';


--
-- Name: user_publickey; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_publickey (
    "userId" character varying(32) NOT NULL,
    "keyId" character varying(256) NOT NULL,
    "keyPem" character varying(4096) NOT NULL
);


ALTER TABLE public.user_publickey OWNER TO misskey;

--
-- Name: user_security_key; Type: TABLE; Schema: public; Owner: misskey
--

CREATE TABLE public.user_security_key (
    id character varying NOT NULL,
    "userId" character varying(32) NOT NULL,
    "publicKey" character varying NOT NULL,
    "lastUsed" timestamp with time zone NOT NULL,
    name character varying(30) NOT NULL
);


ALTER TABLE public.user_security_key OWNER TO misskey;

--
-- Name: COLUMN user_security_key.id; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_security_key.id IS 'Variable-length id given to navigator.credentials.get()';


--
-- Name: COLUMN user_security_key."publicKey"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_security_key."publicKey" IS 'Variable-length public key used to verify attestations (hex-encoded).';


--
-- Name: COLUMN user_security_key."lastUsed"; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_security_key."lastUsed" IS 'The date of the last time the UserSecurityKey was successfully validated.';


--
-- Name: COLUMN user_security_key.name; Type: COMMENT; Schema: public; Owner: misskey
--

COMMENT ON COLUMN public.user_security_key.name IS 'User-defined name for this key';


--
-- Name: __chart__active_users id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__active_users ALTER COLUMN id SET DEFAULT nextval('public.__chart__active_users_id_seq'::regclass);


--
-- Name: __chart__drive id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__drive ALTER COLUMN id SET DEFAULT nextval('public.__chart__drive_id_seq'::regclass);


--
-- Name: __chart__federation id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__federation ALTER COLUMN id SET DEFAULT nextval('public.__chart__federation_id_seq'::regclass);


--
-- Name: __chart__hashtag id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__hashtag ALTER COLUMN id SET DEFAULT nextval('public.__chart__hashtag_id_seq'::regclass);


--
-- Name: __chart__instance id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__instance ALTER COLUMN id SET DEFAULT nextval('public.__chart__instance_id_seq'::regclass);


--
-- Name: __chart__network id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__network ALTER COLUMN id SET DEFAULT nextval('public.__chart__network_id_seq'::regclass);


--
-- Name: __chart__notes id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__notes ALTER COLUMN id SET DEFAULT nextval('public.__chart__notes_id_seq'::regclass);


--
-- Name: __chart__per_user_drive id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_drive ALTER COLUMN id SET DEFAULT nextval('public.__chart__per_user_drive_id_seq'::regclass);


--
-- Name: __chart__per_user_following id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_following ALTER COLUMN id SET DEFAULT nextval('public.__chart__per_user_following_id_seq'::regclass);


--
-- Name: __chart__per_user_notes id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_notes ALTER COLUMN id SET DEFAULT nextval('public.__chart__per_user_notes_id_seq'::regclass);


--
-- Name: __chart__per_user_reaction id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_reaction ALTER COLUMN id SET DEFAULT nextval('public.__chart__per_user_reaction_id_seq'::regclass);


--
-- Name: __chart__test id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test ALTER COLUMN id SET DEFAULT nextval('public.__chart__test_id_seq'::regclass);


--
-- Name: __chart__test_grouped id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test_grouped ALTER COLUMN id SET DEFAULT nextval('public.__chart__test_grouped_id_seq'::regclass);


--
-- Name: __chart__test_unique id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test_unique ALTER COLUMN id SET DEFAULT nextval('public.__chart__test_unique_id_seq'::regclass);


--
-- Name: __chart__users id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__users ALTER COLUMN id SET DEFAULT nextval('public.__chart__users_id_seq'::regclass);


--
-- Name: __chart_day__active_users id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__active_users ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__active_users_id_seq'::regclass);


--
-- Name: __chart_day__drive id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__drive ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__drive_id_seq'::regclass);


--
-- Name: __chart_day__federation id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__federation ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__federation_id_seq'::regclass);


--
-- Name: __chart_day__hashtag id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__hashtag ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__hashtag_id_seq'::regclass);


--
-- Name: __chart_day__instance id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__instance ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__instance_id_seq'::regclass);


--
-- Name: __chart_day__network id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__network ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__network_id_seq'::regclass);


--
-- Name: __chart_day__notes id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__notes ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__notes_id_seq'::regclass);


--
-- Name: __chart_day__per_user_drive id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_drive ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__per_user_drive_id_seq'::regclass);


--
-- Name: __chart_day__per_user_following id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_following ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__per_user_following_id_seq'::regclass);


--
-- Name: __chart_day__per_user_notes id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_notes ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__per_user_notes_id_seq'::regclass);


--
-- Name: __chart_day__per_user_reaction id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_reaction ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__per_user_reaction_id_seq'::regclass);


--
-- Name: __chart_day__users id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__users ALTER COLUMN id SET DEFAULT nextval('public.__chart_day__users_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Data for Name: __chart__active_users; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__active_users (id, date, ___local_users, ___remote_users) FROM stdin;
\.


--
-- Data for Name: __chart__drive; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__drive (id, date, "___local_totalCount", "___local_totalSize", "___local_incCount", "___local_incSize", "___local_decCount", "___local_decSize", "___remote_totalCount", "___remote_totalSize", "___remote_incCount", "___remote_incSize", "___remote_decCount", "___remote_decSize") FROM stdin;
\.


--
-- Data for Name: __chart__federation; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__federation (id, date, ___instance_total, ___instance_inc, ___instance_dec) FROM stdin;
\.


--
-- Data for Name: __chart__hashtag; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__hashtag (id, date, "group", ___local_users, ___remote_users) FROM stdin;
\.


--
-- Data for Name: __chart__instance; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__instance (id, date, "group", ___requests_failed, ___requests_succeeded, ___requests_received, ___notes_total, ___notes_inc, ___notes_dec, ___notes_diffs_normal, ___notes_diffs_reply, ___notes_diffs_renote, ___users_total, ___users_inc, ___users_dec, ___following_total, ___following_inc, ___following_dec, ___followers_total, ___followers_inc, ___followers_dec, "___drive_totalFiles", "___drive_totalUsage", "___drive_incFiles", "___drive_incUsage", "___drive_decFiles", "___drive_decUsage") FROM stdin;
\.


--
-- Data for Name: __chart__network; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__network (id, date, "___incomingRequests", "___outgoingRequests", "___totalTime", "___incomingBytes", "___outgoingBytes") FROM stdin;
\.


--
-- Data for Name: __chart__notes; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__notes (id, date, ___local_total, ___local_inc, ___local_dec, ___local_diffs_normal, ___local_diffs_reply, ___local_diffs_renote, ___remote_total, ___remote_inc, ___remote_dec, ___remote_diffs_normal, ___remote_diffs_reply, ___remote_diffs_renote) FROM stdin;
\.


--
-- Data for Name: __chart__per_user_drive; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__per_user_drive (id, date, "group", "___totalCount", "___totalSize", "___incCount", "___incSize", "___decCount", "___decSize") FROM stdin;
\.


--
-- Data for Name: __chart__per_user_following; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__per_user_following (id, date, "group", ___local_followings_total, ___local_followings_inc, ___local_followings_dec, ___local_followers_total, ___local_followers_inc, ___local_followers_dec, ___remote_followings_total, ___remote_followings_inc, ___remote_followings_dec, ___remote_followers_total, ___remote_followers_inc, ___remote_followers_dec) FROM stdin;
\.


--
-- Data for Name: __chart__per_user_notes; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__per_user_notes (id, date, "group", ___total, ___inc, ___dec, ___diffs_normal, ___diffs_reply, ___diffs_renote) FROM stdin;
\.


--
-- Data for Name: __chart__per_user_reaction; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__per_user_reaction (id, date, "group", ___local_count, ___remote_count) FROM stdin;
\.


--
-- Data for Name: __chart__test; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__test (id, date, "group", ___foo_total, ___foo_inc, ___foo_dec) FROM stdin;
\.


--
-- Data for Name: __chart__test_grouped; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__test_grouped (id, date, "group", ___foo_total, ___foo_inc, ___foo_dec) FROM stdin;
\.


--
-- Data for Name: __chart__test_unique; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__test_unique (id, date, "group", ___foo) FROM stdin;
\.


--
-- Data for Name: __chart__users; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart__users (id, date, ___local_total, ___local_inc, ___local_dec, ___remote_total, ___remote_inc, ___remote_dec) FROM stdin;
\.


--
-- Data for Name: __chart_day__active_users; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__active_users (id, date, ___local_users, ___remote_users) FROM stdin;
\.


--
-- Data for Name: __chart_day__drive; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__drive (id, date, "___local_totalCount", "___local_totalSize", "___local_incCount", "___local_incSize", "___local_decCount", "___local_decSize", "___remote_totalCount", "___remote_totalSize", "___remote_incCount", "___remote_incSize", "___remote_decCount", "___remote_decSize") FROM stdin;
\.


--
-- Data for Name: __chart_day__federation; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__federation (id, date, ___instance_total, ___instance_inc, ___instance_dec) FROM stdin;
\.


--
-- Data for Name: __chart_day__hashtag; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__hashtag (id, date, "group", ___local_users, ___remote_users) FROM stdin;
\.


--
-- Data for Name: __chart_day__instance; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__instance (id, date, "group", ___requests_failed, ___requests_succeeded, ___requests_received, ___notes_total, ___notes_inc, ___notes_dec, ___notes_diffs_normal, ___notes_diffs_reply, ___notes_diffs_renote, ___users_total, ___users_inc, ___users_dec, ___following_total, ___following_inc, ___following_dec, ___followers_total, ___followers_inc, ___followers_dec, "___drive_totalFiles", "___drive_totalUsage", "___drive_incFiles", "___drive_incUsage", "___drive_decFiles", "___drive_decUsage") FROM stdin;
\.


--
-- Data for Name: __chart_day__network; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__network (id, date, "___incomingRequests", "___outgoingRequests", "___totalTime", "___incomingBytes", "___outgoingBytes") FROM stdin;
\.


--
-- Data for Name: __chart_day__notes; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__notes (id, date, ___local_total, ___local_inc, ___local_dec, ___local_diffs_normal, ___local_diffs_reply, ___local_diffs_renote, ___remote_total, ___remote_inc, ___remote_dec, ___remote_diffs_normal, ___remote_diffs_reply, ___remote_diffs_renote) FROM stdin;
\.


--
-- Data for Name: __chart_day__per_user_drive; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__per_user_drive (id, date, "group", "___totalCount", "___totalSize", "___incCount", "___incSize", "___decCount", "___decSize") FROM stdin;
\.


--
-- Data for Name: __chart_day__per_user_following; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__per_user_following (id, date, "group", ___local_followings_total, ___local_followings_inc, ___local_followings_dec, ___local_followers_total, ___local_followers_inc, ___local_followers_dec, ___remote_followings_total, ___remote_followings_inc, ___remote_followings_dec, ___remote_followers_total, ___remote_followers_inc, ___remote_followers_dec) FROM stdin;
\.


--
-- Data for Name: __chart_day__per_user_notes; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__per_user_notes (id, date, "group", ___total, ___inc, ___dec, ___diffs_normal, ___diffs_reply, ___diffs_renote) FROM stdin;
\.


--
-- Data for Name: __chart_day__per_user_reaction; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__per_user_reaction (id, date, "group", ___local_count, ___remote_count) FROM stdin;
\.


--
-- Data for Name: __chart_day__users; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.__chart_day__users (id, date, ___local_total, ___local_inc, ___local_dec, ___remote_total, ___remote_inc, ___remote_dec) FROM stdin;
\.


--
-- Data for Name: abuse_user_report; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.abuse_user_report (id, "createdAt", "targetUserId", "reporterId", "assigneeId", resolved, comment, "targetUserHost", "reporterHost") FROM stdin;
\.


--
-- Data for Name: access_token; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.access_token (id, "createdAt", token, hash, "userId", "appId", "lastUsedAt", session, name, description, "iconUrl", permission, fetched) FROM stdin;
\.


--
-- Data for Name: ad; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.ad (id, "createdAt", "expiresAt", place, priority, url, "imageUrl", memo, ratio) FROM stdin;
\.


--
-- Data for Name: announcement; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.announcement (id, "createdAt", text, title, "imageUrl", "updatedAt") FROM stdin;
\.


--
-- Data for Name: announcement_read; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.announcement_read (id, "userId", "announcementId", "createdAt") FROM stdin;
\.


--
-- Data for Name: antenna; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.antenna (id, "createdAt", "userId", name, src, "userListId", keywords, "withFile", expression, notify, "caseSensitive", "withReplies", "userGroupJoiningId", users, "excludeKeywords") FROM stdin;
\.


--
-- Data for Name: antenna_note; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.antenna_note (id, "noteId", "antennaId", read) FROM stdin;
\.


--
-- Data for Name: app; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.app (id, "createdAt", "userId", secret, name, description, permission, "callbackUrl") FROM stdin;
\.


--
-- Data for Name: attestation_challenge; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.attestation_challenge (id, "userId", challenge, "createdAt", "registrationChallenge") FROM stdin;
\.


--
-- Data for Name: auth_session; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.auth_session (id, "createdAt", token, "userId", "appId") FROM stdin;
\.


--
-- Data for Name: blocking; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.blocking (id, "createdAt", "blockeeId", "blockerId") FROM stdin;
\.


--
-- Data for Name: channel; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.channel (id, "createdAt", "lastNotedAt", "userId", name, description, "bannerId", "notesCount", "usersCount") FROM stdin;
\.


--
-- Data for Name: channel_following; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.channel_following (id, "createdAt", "followeeId", "followerId") FROM stdin;
\.


--
-- Data for Name: channel_note_pining; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.channel_note_pining (id, "createdAt", "channelId", "noteId") FROM stdin;
\.


--
-- Data for Name: clip; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.clip (id, "createdAt", "userId", name, "isPublic", description) FROM stdin;
\.


--
-- Data for Name: clip_note; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.clip_note (id, "noteId", "clipId") FROM stdin;
\.


--
-- Data for Name: drive_file; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.drive_file (id, "createdAt", "userId", "userHost", md5, name, type, size, comment, properties, "storedInternal", url, "thumbnailUrl", "webpublicUrl", "accessKey", "thumbnailAccessKey", "webpublicAccessKey", uri, src, "folderId", "isSensitive", "isLink", blurhash) FROM stdin;
\.


--
-- Data for Name: drive_folder; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.drive_folder (id, "createdAt", name, "userId", "parentId") FROM stdin;
\.


--
-- Data for Name: emoji; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.emoji (id, "updatedAt", name, host, url, uri, type, aliases, category) FROM stdin;
\.


--
-- Data for Name: follow_request; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.follow_request (id, "createdAt", "followeeId", "followerId", "requestId", "followerHost", "followerInbox", "followerSharedInbox", "followeeHost", "followeeInbox", "followeeSharedInbox") FROM stdin;
\.


--
-- Data for Name: following; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.following (id, "createdAt", "followeeId", "followerId", "followerHost", "followerInbox", "followerSharedInbox", "followeeHost", "followeeInbox", "followeeSharedInbox") FROM stdin;
\.


--
-- Data for Name: gallery_like; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.gallery_like (id, "createdAt", "userId", "postId") FROM stdin;
\.


--
-- Data for Name: gallery_post; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.gallery_post (id, "createdAt", "updatedAt", title, description, "userId", "fileIds", "isSensitive", "likedCount", tags) FROM stdin;
\.


--
-- Data for Name: hashtag; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.hashtag (id, name, "mentionedUserIds", "mentionedUsersCount", "mentionedLocalUserIds", "mentionedLocalUsersCount", "mentionedRemoteUserIds", "mentionedRemoteUsersCount", "attachedUserIds", "attachedUsersCount", "attachedLocalUserIds", "attachedLocalUsersCount", "attachedRemoteUserIds", "attachedRemoteUsersCount") FROM stdin;
\.


--
-- Data for Name: instance; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.instance (id, "caughtAt", host, "usersCount", "notesCount", "followingCount", "followersCount", "driveUsage", "driveFiles", "latestRequestSentAt", "latestStatus", "latestRequestReceivedAt", "lastCommunicatedAt", "isNotResponding", "softwareName", "softwareVersion", "openRegistrations", name, description, "maintainerName", "maintainerEmail", "infoUpdatedAt", "isSuspended", "iconUrl", "themeColor", "faviconUrl") FROM stdin;
\.


--
-- Data for Name: messaging_message; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.messaging_message (id, "createdAt", "userId", "recipientId", text, "isRead", "fileId", "groupId", reads, uri) FROM stdin;
\.


--
-- Data for Name: meta; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.meta (id, name, description, "maintainerName", "maintainerEmail", "disableRegistration", "disableLocalTimeline", "disableGlobalTimeline", "useStarForReactionFallback", langs, "hiddenTags", "blockedHosts", "mascotImageUrl", "bannerUrl", "errorImageUrl", "iconUrl", "cacheRemoteFiles", "enableRecaptcha", "recaptchaSiteKey", "recaptchaSecretKey", "localDriveCapacityMb", "remoteDriveCapacityMb", "maxNoteTextLength", "summalyProxy", "enableEmail", email, "smtpSecure", "smtpHost", "smtpPort", "smtpUser", "smtpPass", "enableServiceWorker", "swPublicKey", "swPrivateKey", "enableTwitterIntegration", "twitterConsumerKey", "twitterConsumerSecret", "enableGithubIntegration", "githubClientId", "githubClientSecret", "enableDiscordIntegration", "discordClientId", "discordClientSecret", "pinnedUsers", "ToSUrl", "repositoryUrl", "feedbackUrl", "useObjectStorage", "objectStorageBucket", "objectStoragePrefix", "objectStorageBaseUrl", "objectStorageEndpoint", "objectStorageRegion", "objectStorageAccessKey", "objectStorageSecretKey", "objectStoragePort", "objectStorageUseSSL", "proxyRemoteFiles", "proxyAccountId", "objectStorageUseProxy", "enableHcaptcha", "hcaptchaSiteKey", "hcaptchaSecretKey", "objectStorageSetPublicRead", "pinnedPages", "backgroundImageUrl", "logoImageUrl", "pinnedClipId", "objectStorageS3ForcePathStyle", "deeplAuthKey", "deeplIsPro", "emailRequiredForSignup") FROM stdin;
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.migrations (id, "timestamp", name) FROM stdin;
1	1000000000000	Init1000000000000
2	1556348509290	Pages1556348509290
3	1556746559567	UserProfile1556746559567
4	1557476068003	PinnedUsers1557476068003
5	1557761316509	AddSomeUrls1557761316509
6	1557932705754	ObjectStorageSetting1557932705754
7	1558072954435	PageLike1558072954435
8	1558103093633	UserGroup1558103093633
9	1558257926829	UserGroupInvite1558257926829
10	1558266512381	UserListJoining1558266512381
11	1561706992953	webauthn1561706992953
12	1561873850023	ChartIndexes1561873850023
13	1562422242907	PasswordLessLogin1562422242907
14	1562444565093	PinnedPage1562444565093
15	1562448332510	PageTitleHideOption1562448332510
16	1562869971568	ModerationLog1562869971568
17	1563757595828	UsedUsername1563757595828
18	1565634203341	room1565634203341
19	1571220798684	CustomEmojiCategory1571220798684
20	1572760203493	nodeinfo1572760203493
21	1576269851876	TalkFederationId1576269851876
22	1576869585998	ProxyRemoteFiles1576869585998
23	1579267006611	v121579267006611
24	1579270193251	v1221579270193251
25	1579282808087	v1231579282808087
26	1579544426412	v1241579544426412
27	1579977526288	v1251579977526288
28	1579993013959	v1261579993013959
29	1580069531114	v1271580069531114
30	1580148575182	v1281580148575182
31	1580154400017	v1291580154400017
32	1580276619901	v12101580276619901
33	1580331224276	v12111580331224276
34	1580508795118	v12121580508795118
35	1580543501339	v12131580543501339
36	1580864313253	v12141580864313253
37	1581526429287	userGroupInvitation1581526429287
38	1581695816408	userGroupAntenna1581695816408
39	1581708415836	driveUserFolderIdIndex1581708415836
40	1581979837262	promo1581979837262
41	1582019042083	featuredInjecttion1582019042083
42	1582210532752	antennaExclude1582210532752
43	1582875306439	noteReactionLength1582875306439
44	1585361548360	miauth1585361548360
45	1585385921215	customNotification1585385921215
46	1585772678853	apUrl1585772678853
47	1586624197029	AddObjectStorageUseProxy1586624197029
48	1586641139527	remoteReaction1586641139527
49	1586708940386	pageAiScript1586708940386
50	1588044505511	hCaptcha1588044505511
51	1589023282116	pubRelay1589023282116
52	1595075960584	blurhash1595075960584
53	1595077605646	blurhashForAvatarBanner1595077605646
54	1595676934834	instanceIconUrl1595676934834
55	1595771249699	wordMute1595771249699
56	1595782306083	wordMute21595782306083
57	1596548170836	channel1596548170836
58	1596786425167	channel21596786425167
59	1597230137744	objectStorageSetPublicRead1597230137744
60	1597236229720	IncludingNotificationTypes1597236229720
61	1597385880794	addSensitiveIndex1597385880794
62	1597459042300	channelUnread1597459042300
63	1597893996136	ChannelNoteIdDescIndex1597893996136
64	1600353287890	mutingNotificationTypes1600353287890
65	1603094348345	refineAbuseUserReport1603094348345
66	1603095701770	refineAbuseUserReport21603095701770
67	1603776877564	instanceThemeColor1603776877564
68	1603781553011	instanceFavicon1603781553011
69	1604821689616	deleteAutoWatch1604821689616
70	1605408848373	clipDescription1605408848373
71	1605408971051	comments1605408971051
72	1605585339718	instancePinnedPages1605585339718
73	1605965516823	instanceImages1605965516823
74	1606191203881	noCrawle1606191203881
75	1607151207216	instancePinnedClip1607151207216
76	1607353487793	isExplorable1607353487793
77	1610277136869	registry1610277136869
78	1610277585759	registry21610277585759
79	1610283021566	registry31610283021566
80	1611354329133	followersUri1611354329133
81	1611397665007	gallery1611397665007
82	1611547387175	objectStorageS3ForcePathStyle1611547387175
83	1612619156584	announcementEmail1612619156584
84	1613155914446	emailNotificationTypes1613155914446
85	1613181457597	userLang1613181457597
86	1613503367223	useBigintForDriveUsage1613503367223
87	1615965918224	chartV21615965918224
88	1615966519402	chartV221615966519402
89	1618637372000	userLastActiveDate1618637372000
90	1618639857000	userHideOnlineStatus1618639857000
91	1619942102890	passwordReset1619942102890
92	1620019354680	ad1620019354680
93	1620364649428	ad21620364649428
94	1621479946000	addNoteIndexes1621479946000
95	1622679304522	userProfileDescriptionLength1622679304522
96	1622681548499	logMessageLength1622681548499
97	1629004542760	chartReindex1629004542760
98	1629024377804	deeplIntegration1629024377804
99	1629288472000	fixChannelUserId1629288472000
100	1629512953000	isUserDeleted1629512953000
101	1629778475000	deeplIntegration21629778475000
102	1629968054000	userInstanceBlocks1629968054000
103	1633068642000	emailRequiredForSignup1633068642000
104	1633071909016	userPending1633071909016
105	1634486652000	userPublicReactions1634486652000
106	1634902659689	deleteLog1634902659689
107	1635500777168	noteThreadMute1635500777168
108	1636197624383	ffVisibility1636197624383
109	1636697408073	removeViaMobile1636697408073
110	1639325650583	chartV31639325650583
\.


--
-- Data for Name: moderation_log; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.moderation_log (id, "createdAt", "userId", type, info) FROM stdin;
\.


--
-- Data for Name: muted_note; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.muted_note (id, "noteId", "userId", reason) FROM stdin;
\.


--
-- Data for Name: muting; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.muting (id, "createdAt", "muteeId", "muterId") FROM stdin;
\.


--
-- Data for Name: note; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note (id, "createdAt", "replyId", "renoteId", text, name, cw, "userId", "localOnly", "renoteCount", "repliesCount", reactions, visibility, uri, score, "fileIds", "attachedFileTypes", "visibleUserIds", mentions, "mentionedRemoteUsers", emojis, tags, "hasPoll", "userHost", "replyUserId", "replyUserHost", "renoteUserId", "renoteUserHost", url, "channelId", "threadId") FROM stdin;
\.


--
-- Data for Name: note_favorite; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note_favorite (id, "createdAt", "userId", "noteId") FROM stdin;
\.


--
-- Data for Name: note_reaction; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note_reaction (id, "createdAt", "userId", "noteId", reaction) FROM stdin;
\.


--
-- Data for Name: note_thread_muting; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note_thread_muting (id, "createdAt", "userId", "threadId") FROM stdin;
\.


--
-- Data for Name: note_unread; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note_unread (id, "userId", "noteId", "noteUserId", "isSpecified", "isMentioned", "noteChannelId") FROM stdin;
\.


--
-- Data for Name: note_watching; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.note_watching (id, "createdAt", "userId", "noteId", "noteUserId") FROM stdin;
\.


--
-- Data for Name: notification; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.notification (id, "createdAt", "notifieeId", "notifierId", "isRead", "noteId", reaction, choice, "followRequestId", type, "userGroupInvitationId", "customBody", "customHeader", "customIcon", "appAccessTokenId") FROM stdin;
\.


--
-- Data for Name: page; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.page (id, "createdAt", "updatedAt", title, name, summary, "alignCenter", font, "userId", "eyeCatchingImageId", content, variables, visibility, "visibleUserIds", "likedCount", "hideTitleWhenPinned", script) FROM stdin;
\.


--
-- Data for Name: page_like; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.page_like (id, "createdAt", "userId", "pageId") FROM stdin;
\.


--
-- Data for Name: password_reset_request; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.password_reset_request (id, "createdAt", token, "userId") FROM stdin;
\.


--
-- Data for Name: poll; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.poll ("noteId", "expiresAt", multiple, choices, votes, "noteVisibility", "userId", "userHost") FROM stdin;
\.


--
-- Data for Name: poll_vote; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.poll_vote (id, "createdAt", "userId", "noteId", choice) FROM stdin;
\.


--
-- Data for Name: promo_note; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.promo_note ("noteId", "expiresAt", "userId") FROM stdin;
\.


--
-- Data for Name: promo_read; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.promo_read (id, "createdAt", "userId", "noteId") FROM stdin;
\.


--
-- Data for Name: registration_ticket; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.registration_ticket (id, "createdAt", code) FROM stdin;
\.


--
-- Data for Name: registry_item; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.registry_item (id, "createdAt", "updatedAt", "userId", key, scope, domain, value) FROM stdin;
\.


--
-- Data for Name: relay; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.relay (id, inbox, status) FROM stdin;
\.


--
-- Data for Name: reversi_game; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.reversi_game (id, "createdAt", "startedAt", "user1Id", "user2Id", "user1Accepted", "user2Accepted", black, "isStarted", "isEnded", "winnerId", surrendered, logs, map, bw, "isLlotheo", "canPutEverywhere", "loopedBoard", form1, form2, crc32) FROM stdin;
\.


--
-- Data for Name: reversi_matching; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.reversi_matching (id, "createdAt", "parentId", "childId") FROM stdin;
\.


--
-- Data for Name: signin; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.signin (id, "createdAt", "userId", ip, headers, success) FROM stdin;
\.


--
-- Data for Name: sw_subscription; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.sw_subscription (id, "createdAt", "userId", endpoint, auth, publickey) FROM stdin;
\.


--
-- Data for Name: used_username; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.used_username (username, "createdAt") FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public."user" (id, "createdAt", "updatedAt", "lastFetchedAt", username, "usernameLower", name, "followersCount", "followingCount", "notesCount", "avatarId", "bannerId", tags, "avatarUrl", "bannerUrl", "isSuspended", "isSilenced", "isLocked", "isBot", "isCat", "isAdmin", "isModerator", emojis, host, inbox, "sharedInbox", featured, uri, token, "avatarBlurhash", "bannerBlurhash", "isExplorable", "followersUri", "lastActiveDate", "hideOnlineStatus", "isDeleted") FROM stdin;
\.


--
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_group (id, "createdAt", name, "userId", "isPrivate") FROM stdin;
\.


--
-- Data for Name: user_group_invitation; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_group_invitation (id, "createdAt", "userId", "userGroupId") FROM stdin;
\.


--
-- Data for Name: user_group_invite; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_group_invite (id, "createdAt", "userId", "userGroupId") FROM stdin;
\.


--
-- Data for Name: user_group_joining; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_group_joining (id, "createdAt", "userId", "userGroupId") FROM stdin;
\.


--
-- Data for Name: user_keypair; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_keypair ("userId", "publicKey", "privateKey") FROM stdin;
\.


--
-- Data for Name: user_list; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_list (id, "createdAt", "userId", name) FROM stdin;
\.


--
-- Data for Name: user_list_joining; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_list_joining (id, "createdAt", "userId", "userListId") FROM stdin;
\.


--
-- Data for Name: user_note_pining; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_note_pining (id, "createdAt", "userId", "noteId") FROM stdin;
\.


--
-- Data for Name: user_pending; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_pending (id, "createdAt", code, username, email, password) FROM stdin;
\.


--
-- Data for Name: user_profile; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_profile ("userId", location, birthday, description, fields, url, email, "emailVerifyCode", "emailVerified", "twoFactorTempSecret", "twoFactorSecret", "twoFactorEnabled", password, "clientData", "autoAcceptFollowed", "alwaysMarkNsfw", "carefulBot", "userHost", "securityKeysAvailable", "usePasswordLessLogin", "pinnedPageId", room, integrations, "injectFeaturedNote", "enableWordMute", "mutedWords", "mutingNotificationTypes", "noCrawle", "receiveAnnouncementEmail", "emailNotificationTypes", lang, "mutedInstances", "publicReactions", "ffVisibility") FROM stdin;
\.


--
-- Data for Name: user_publickey; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_publickey ("userId", "keyId", "keyPem") FROM stdin;
\.


--
-- Data for Name: user_security_key; Type: TABLE DATA; Schema: public; Owner: misskey
--

COPY public.user_security_key (id, "userId", "publicKey", "lastUsed", name) FROM stdin;
\.


--
-- Name: __chart__active_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__active_users_id_seq', 1, false);


--
-- Name: __chart__drive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__drive_id_seq', 1, false);


--
-- Name: __chart__federation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__federation_id_seq', 1, false);


--
-- Name: __chart__hashtag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__hashtag_id_seq', 1, false);


--
-- Name: __chart__instance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__instance_id_seq', 1, false);


--
-- Name: __chart__network_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__network_id_seq', 1, false);


--
-- Name: __chart__notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__notes_id_seq', 1, false);


--
-- Name: __chart__per_user_drive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__per_user_drive_id_seq', 1, false);


--
-- Name: __chart__per_user_following_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__per_user_following_id_seq', 1, false);


--
-- Name: __chart__per_user_notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__per_user_notes_id_seq', 1, false);


--
-- Name: __chart__per_user_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__per_user_reaction_id_seq', 1, false);


--
-- Name: __chart__test_grouped_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__test_grouped_id_seq', 1, false);


--
-- Name: __chart__test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__test_id_seq', 1, false);


--
-- Name: __chart__test_unique_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__test_unique_id_seq', 1, false);


--
-- Name: __chart__users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart__users_id_seq', 1, false);


--
-- Name: __chart_day__active_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__active_users_id_seq', 1, false);


--
-- Name: __chart_day__drive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__drive_id_seq', 1, false);


--
-- Name: __chart_day__federation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__federation_id_seq', 1, false);


--
-- Name: __chart_day__hashtag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__hashtag_id_seq', 1, false);


--
-- Name: __chart_day__instance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__instance_id_seq', 1, false);


--
-- Name: __chart_day__network_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__network_id_seq', 1, false);


--
-- Name: __chart_day__notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__notes_id_seq', 1, false);


--
-- Name: __chart_day__per_user_drive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__per_user_drive_id_seq', 1, false);


--
-- Name: __chart_day__per_user_following_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__per_user_following_id_seq', 1, false);


--
-- Name: __chart_day__per_user_notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__per_user_notes_id_seq', 1, false);


--
-- Name: __chart_day__per_user_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__per_user_reaction_id_seq', 1, false);


--
-- Name: __chart_day__users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.__chart_day__users_id_seq', 1, false);


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: misskey
--

SELECT pg_catalog.setval('public.migrations_id_seq', 110, true);


--
-- Name: ad PK_0193d5ef09746e88e9ea92c634d; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.ad
    ADD CONSTRAINT "PK_0193d5ef09746e88e9ea92c634d" PRIMARY KEY (id);


--
-- Name: __chart__notes PK_0aec823fa85c7f901bdb3863b14; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__notes
    ADD CONSTRAINT "PK_0aec823fa85c7f901bdb3863b14" PRIMARY KEY (id);


--
-- Name: user_publickey PK_10c146e4b39b443ede016f6736d; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_publickey
    ADD CONSTRAINT "PK_10c146e4b39b443ede016f6736d" PRIMARY KEY ("userId");


--
-- Name: user_list_joining PK_11abb3768da1c5f8de101c9df45; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_list_joining
    ADD CONSTRAINT "PK_11abb3768da1c5f8de101c9df45" PRIMARY KEY (id);


--
-- Name: __chart__instance PK_1267c67c7c2d47b4903975f2c00; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__instance
    ADD CONSTRAINT "PK_1267c67c7c2d47b4903975f2c00" PRIMARY KEY (id);


--
-- Name: __chart_day__hashtag PK_13d5a3b089344e5557f8e0980b4; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__hashtag
    ADD CONSTRAINT "PK_13d5a3b089344e5557f8e0980b4" PRIMARY KEY (id);


--
-- Name: user_group_joining PK_15f2425885253c5507e1599cfe7; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_joining
    ADD CONSTRAINT "PK_15f2425885253c5507e1599cfe7" PRIMARY KEY (id);


--
-- Name: user_group_invitation PK_160c63ec02bf23f6a5c5e8140d6; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invitation
    ADD CONSTRAINT "PK_160c63ec02bf23f6a5c5e8140d6" PRIMARY KEY (id);


--
-- Name: note_unread PK_1904eda61a784f57e6e51fa9c1f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_unread
    ADD CONSTRAINT "PK_1904eda61a784f57e6e51fa9c1f" PRIMARY KEY (id);


--
-- Name: auth_session PK_19354ed146424a728c1112a8cbf; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.auth_session
    ADD CONSTRAINT "PK_19354ed146424a728c1112a8cbf" PRIMARY KEY (id);


--
-- Name: __chart_day__per_user_drive PK_1ae135254c137011645da7f4045; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_drive
    ADD CONSTRAINT "PK_1ae135254c137011645da7f4045" PRIMARY KEY (id);


--
-- Name: __chart_day__notes PK_1fa4139e1f338272b758d05e090; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__notes
    ADD CONSTRAINT "PK_1fa4139e1f338272b758d05e090" PRIMARY KEY (id);


--
-- Name: muting PK_2e92d06c8b5c602eeb27ca9ba48; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muting
    ADD CONSTRAINT "PK_2e92d06c8b5c602eeb27ca9ba48" PRIMARY KEY (id);


--
-- Name: __chart__active_users PK_317237a9f733b970604a11e314f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__active_users
    ADD CONSTRAINT "PK_317237a9f733b970604a11e314f" PRIMARY KEY (id);


--
-- Name: __chart__per_user_notes PK_334acf6e915af2f29edc11b8e50; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_notes
    ADD CONSTRAINT "PK_334acf6e915af2f29edc11b8e50" PRIMARY KEY (id);


--
-- Name: user_group_invite PK_3893884af0d3a5f4d01e7921a97; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invite
    ADD CONSTRAINT "PK_3893884af0d3a5f4d01e7921a97" PRIMARY KEY (id);


--
-- Name: user_group PK_3c29fba6fe013ec8724378ce7c9; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group
    ADD CONSTRAINT "PK_3c29fba6fe013ec8724378ce7c9" PRIMARY KEY (id);


--
-- Name: user_security_key PK_3e508571121ab39c5f85d10c166; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_security_key
    ADD CONSTRAINT "PK_3e508571121ab39c5f85d10c166" PRIMARY KEY (id);


--
-- Name: __chart__test_unique PK_409bac9c97cc612d8500012319d; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test_unique
    ADD CONSTRAINT "PK_409bac9c97cc612d8500012319d" PRIMARY KEY (id);


--
-- Name: drive_file PK_43ddaaaf18c9e68029b7cbb032e; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_file
    ADD CONSTRAINT "PK_43ddaaaf18c9e68029b7cbb032e" PRIMARY KEY (id);


--
-- Name: channel_note_pining PK_44f7474496bcf2e4b741681146d; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_note_pining
    ADD CONSTRAINT "PK_44f7474496bcf2e4b741681146d" PRIMARY KEY (id);


--
-- Name: __chart_day__instance PK_479a8ff9d959274981087043023; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__instance
    ADD CONSTRAINT "PK_479a8ff9d959274981087043023" PRIMARY KEY (id);


--
-- Name: note_watching PK_49286fdb23725945a74aa27d757; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_watching
    ADD CONSTRAINT "PK_49286fdb23725945a74aa27d757" PRIMARY KEY (id);


--
-- Name: announcement_read PK_4b90ad1f42681d97b2683890c5e; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.announcement_read
    ADD CONSTRAINT "PK_4b90ad1f42681d97b2683890c5e" PRIMARY KEY (id);


--
-- Name: __chart__users PK_4dfcf2c78d03524b9eb2c99d328; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__users
    ADD CONSTRAINT "PK_4dfcf2c78d03524b9eb2c99d328" PRIMARY KEY (id);


--
-- Name: user_profile PK_51cb79b5555effaf7d69ba1cff9; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_profile
    ADD CONSTRAINT "PK_51cb79b5555effaf7d69ba1cff9" PRIMARY KEY ("userId");


--
-- Name: follow_request PK_53a9aa3725f7a3deb150b39dbfc; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.follow_request
    ADD CONSTRAINT "PK_53a9aa3725f7a3deb150b39dbfc" PRIMARY KEY (id);


--
-- Name: __chart_day__per_user_notes PK_58bab6b6d3ad9310cbc7460fd28; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_notes
    ADD CONSTRAINT "PK_58bab6b6d3ad9310cbc7460fd28" PRIMARY KEY (id);


--
-- Name: channel PK_590f33ee6ee7d76437acf362e39; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT "PK_590f33ee6ee7d76437acf362e39" PRIMARY KEY (id);


--
-- Name: promo_read PK_61917c1541002422b703318b7c9; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.promo_read
    ADD CONSTRAINT "PK_61917c1541002422b703318b7c9" PRIMARY KEY (id);


--
-- Name: registry_item PK_64b3f7e6008b4d89b826cd3af95; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.registry_item
    ADD CONSTRAINT "PK_64b3f7e6008b4d89b826cd3af95" PRIMARY KEY (id);


--
-- Name: __chart_day__per_user_following PK_68ce6b67da57166da66fc8fb27e; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_following
    ADD CONSTRAINT "PK_68ce6b67da57166da66fc8fb27e" PRIMARY KEY (id);


--
-- Name: notification PK_705b6c7cdf9b2c2ff7ac7872cb7; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "PK_705b6c7cdf9b2c2ff7ac7872cb7" PRIMARY KEY (id);


--
-- Name: page PK_742f4117e065c5b6ad21b37ba1f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page
    ADD CONSTRAINT "PK_742f4117e065c5b6ad21b37ba1f" PRIMARY KEY (id);


--
-- Name: note_reaction PK_767ec729b108799b587a3fcc9cf; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_reaction
    ADD CONSTRAINT "PK_767ec729b108799b587a3fcc9cf" PRIMARY KEY (id);


--
-- Name: reversi_game PK_76b30eeba71b1193ad7c5311c3f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_game
    ADD CONSTRAINT "PK_76b30eeba71b1193ad7c5311c3f" PRIMARY KEY (id);


--
-- Name: relay PK_78ebc9cfddf4292633b7ba57aee; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.relay
    ADD CONSTRAINT "PK_78ebc9cfddf4292633b7ba57aee" PRIMARY KEY (id);


--
-- Name: used_username PK_78fd79d2d24c6ac2f4cc9a31a5d; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.used_username
    ADD CONSTRAINT "PK_78fd79d2d24c6ac2f4cc9a31a5d" PRIMARY KEY (username);


--
-- Name: drive_folder PK_7a0c089191f5ebdc214e0af808a; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_folder
    ADD CONSTRAINT "PK_7a0c089191f5ebdc214e0af808a" PRIMARY KEY (id);


--
-- Name: __chart_day__federation PK_7ca721c769f31698e0e1331e8e6; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__federation
    ADD CONSTRAINT "PK_7ca721c769f31698e0e1331e8e6" PRIMARY KEY (id);


--
-- Name: page_like PK_813f034843af992d3ae0f43c64c; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page_like
    ADD CONSTRAINT "PK_813f034843af992d3ae0f43c64c" PRIMARY KEY (id);


--
-- Name: gallery_like PK_853ab02be39b8de45cd720cc15f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.gallery_like
    ADD CONSTRAINT "PK_853ab02be39b8de45cd720cc15f" PRIMARY KEY (id);


--
-- Name: __chart__per_user_following PK_85bb1b540363a29c2fec83bd907; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_following
    ADD CONSTRAINT "PK_85bb1b540363a29c2fec83bd907" PRIMARY KEY (id);


--
-- Name: abuse_user_report PK_87873f5f5cc5c321a1306b2d18c; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.abuse_user_report
    ADD CONSTRAINT "PK_87873f5f5cc5c321a1306b2d18c" PRIMARY KEY (id);


--
-- Name: user_list PK_87bab75775fd9b1ff822b656402; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_list
    ADD CONSTRAINT "PK_87bab75775fd9b1ff822b656402" PRIMARY KEY (id);


--
-- Name: reversi_matching PK_880bd0afbab232f21c8b9d146cf; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_matching
    ADD CONSTRAINT "PK_880bd0afbab232f21c8b9d146cf" PRIMARY KEY (id);


--
-- Name: muted_note PK_897e2eff1c0b9b64e55ca1418a4; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muted_note
    ADD CONSTRAINT "PK_897e2eff1c0b9b64e55ca1418a4" PRIMARY KEY (id);


--
-- Name: __chart_day__per_user_reaction PK_8af24e2d51ff781a354fe595eda; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_reaction
    ADD CONSTRAINT "PK_8af24e2d51ff781a354fe595eda" PRIMARY KEY (id);


--
-- Name: channel_following PK_8b104be7f7415113f2a02cd5bdd; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_following
    ADD CONSTRAINT "PK_8b104be7f7415113f2a02cd5bdd" PRIMARY KEY (id);


--
-- Name: migrations PK_8c82d7f526340ab734260ea46be; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT "PK_8c82d7f526340ab734260ea46be" PRIMARY KEY (id);


--
-- Name: gallery_post PK_8e90d7b6015f2c4518881b14753; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.gallery_post
    ADD CONSTRAINT "PK_8e90d7b6015f2c4518881b14753" PRIMARY KEY (id);


--
-- Name: app PK_9478629fc093d229df09e560aea; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.app
    ADD CONSTRAINT "PK_9478629fc093d229df09e560aea" PRIMARY KEY (id);


--
-- Name: note PK_96d0c172a4fba276b1bbed43058; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT "PK_96d0c172a4fba276b1bbed43058" PRIMARY KEY (id);


--
-- Name: __chart__per_user_reaction PK_984f54dae441e65b633e8d27a7f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_reaction
    ADD CONSTRAINT "PK_984f54dae441e65b633e8d27a7f" PRIMARY KEY (id);


--
-- Name: signin PK_9e96ddc025712616fc492b3b588; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.signin
    ADD CONSTRAINT "PK_9e96ddc025712616fc492b3b588" PRIMARY KEY (id);


--
-- Name: user_note_pining PK_a6a2dad4ae000abce2ea9d9b103; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_note_pining
    ADD CONSTRAINT "PK_a6a2dad4ae000abce2ea9d9b103" PRIMARY KEY (id);


--
-- Name: note_favorite PK_af0da35a60b9fa4463a62082b36; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_favorite
    ADD CONSTRAINT "PK_af0da35a60b9fa4463a62082b36" PRIMARY KEY (id);


--
-- Name: __chart_day__active_users PK_b1790489b14f005ae8f404f5795; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__active_users
    ADD CONSTRAINT "PK_b1790489b14f005ae8f404f5795" PRIMARY KEY (id);


--
-- Name: __chart__federation PK_b39dcd31a0fe1a7757e348e85fd; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__federation
    ADD CONSTRAINT "PK_b39dcd31a0fe1a7757e348e85fd" PRIMARY KEY (id);


--
-- Name: __chart__test PK_b4bc31dffbd1b785276a3ecfc1e; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test
    ADD CONSTRAINT "PK_b4bc31dffbd1b785276a3ecfc1e" PRIMARY KEY (id);


--
-- Name: __chart__network PK_bc4290c2e27fad14ef0c1ca93f3; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__network
    ADD CONSTRAINT "PK_bc4290c2e27fad14ef0c1ca93f3" PRIMARY KEY (id);


--
-- Name: antenna PK_c170b99775e1dccca947c9f2d5f; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna
    ADD CONSTRAINT "PK_c170b99775e1dccca947c9f2d5f" PRIMARY KEY (id);


--
-- Name: __chart__hashtag PK_c32f1ea2b44a5d2f7881e37f8f9; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__hashtag
    ADD CONSTRAINT "PK_c32f1ea2b44a5d2f7881e37f8f9" PRIMARY KEY (id);


--
-- Name: meta PK_c4c17a6c2bd7651338b60fc590b; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.meta
    ADD CONSTRAINT "PK_c4c17a6c2bd7651338b60fc590b" PRIMARY KEY (id);


--
-- Name: following PK_c76c6e044bdf76ecf8bfb82a645; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.following
    ADD CONSTRAINT "PK_c76c6e044bdf76ecf8bfb82a645" PRIMARY KEY (id);


--
-- Name: __chart_day__network PK_cac499d6f471042dfed1e7e0132; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__network
    ADD CONSTRAINT "PK_cac499d6f471042dfed1e7e0132" PRIMARY KEY (id);


--
-- Name: user PK_cace4a159ff9f2512dd42373760; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY (id);


--
-- Name: hashtag PK_cb36eb8af8412bfa978f1165d78; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.hashtag
    ADD CONSTRAINT "PK_cb36eb8af8412bfa978f1165d78" PRIMARY KEY (id);


--
-- Name: moderation_log PK_d0adca6ecfd068db83e4526cc26; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.moderation_log
    ADD CONSTRAINT "PK_d0adca6ecfd068db83e4526cc26" PRIMARY KEY (id);


--
-- Name: attestation_challenge PK_d0ba6786e093f1bcb497572a6b5; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.attestation_challenge
    ADD CONSTRAINT "PK_d0ba6786e093f1bcb497572a6b5" PRIMARY KEY (id, "userId");


--
-- Name: __chart__per_user_drive PK_d0ef23d24d666e1a44a0cd3d208; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_drive
    ADD CONSTRAINT "PK_d0ef23d24d666e1a44a0cd3d208" PRIMARY KEY (id);


--
-- Name: user_pending PK_d4c84e013c98ec02d19b8fbbafa; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_pending
    ADD CONSTRAINT "PK_d4c84e013c98ec02d19b8fbbafa" PRIMARY KEY (id);


--
-- Name: __chart_day__users PK_d7f7185abb9851f70c4726c54bd; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__users
    ADD CONSTRAINT "PK_d7f7185abb9851f70c4726c54bd" PRIMARY KEY (id);


--
-- Name: poll PK_da851e06d0dfe2ef397d8b1bf1b; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.poll
    ADD CONSTRAINT "PK_da851e06d0dfe2ef397d8b1bf1b" PRIMARY KEY ("noteId");


--
-- Name: messaging_message PK_db398fd79dc95d0eb8c30456eaa; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.messaging_message
    ADD CONSTRAINT "PK_db398fd79dc95d0eb8c30456eaa" PRIMARY KEY (id);


--
-- Name: emoji PK_df74ce05e24999ee01ea0bc50a3; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.emoji
    ADD CONSTRAINT "PK_df74ce05e24999ee01ea0bc50a3" PRIMARY KEY (id);


--
-- Name: announcement PK_e0ef0550174fd1099a308fd18a0; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.announcement
    ADD CONSTRAINT "PK_e0ef0550174fd1099a308fd18a0" PRIMARY KEY (id);


--
-- Name: promo_note PK_e263909ca4fe5d57f8d4230dd5c; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.promo_note
    ADD CONSTRAINT "PK_e263909ca4fe5d57f8d4230dd5c" PRIMARY KEY ("noteId");


--
-- Name: blocking PK_e5d9a541cc1965ee7e048ea09dd; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.blocking
    ADD CONSTRAINT "PK_e5d9a541cc1965ee7e048ea09dd" PRIMARY KEY (id);


--
-- Name: __chart_day__drive PK_e7ec0de057c77c40fc8d8b62151; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__drive
    ADD CONSTRAINT "PK_e7ec0de057c77c40fc8d8b62151" PRIMARY KEY (id);


--
-- Name: sw_subscription PK_e8f763631530051b95eb6279b91; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.sw_subscription
    ADD CONSTRAINT "PK_e8f763631530051b95eb6279b91" PRIMARY KEY (id);


--
-- Name: clip_note PK_e94cda2f40a99b57e032a1a738b; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.clip_note
    ADD CONSTRAINT "PK_e94cda2f40a99b57e032a1a738b" PRIMARY KEY (id);


--
-- Name: instance PK_eaf60e4a0c399c9935413e06474; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.instance
    ADD CONSTRAINT "PK_eaf60e4a0c399c9935413e06474" PRIMARY KEY (id);


--
-- Name: note_thread_muting PK_ec5936d94d1a0369646d12a3a47; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_thread_muting
    ADD CONSTRAINT "PK_ec5936d94d1a0369646d12a3a47" PRIMARY KEY (id);


--
-- Name: clip PK_f0685dac8d4dd056d7255670b75; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.clip
    ADD CONSTRAINT "PK_f0685dac8d4dd056d7255670b75" PRIMARY KEY (id);


--
-- Name: registration_ticket PK_f11696b6fafcf3662d4292734f8; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.registration_ticket
    ADD CONSTRAINT "PK_f11696b6fafcf3662d4292734f8" PRIMARY KEY (id);


--
-- Name: access_token PK_f20f028607b2603deabd8182d12; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.access_token
    ADD CONSTRAINT "PK_f20f028607b2603deabd8182d12" PRIMARY KEY (id);


--
-- Name: user_keypair PK_f4853eb41ab722fe05f81cedeb6; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_keypair
    ADD CONSTRAINT "PK_f4853eb41ab722fe05f81cedeb6" PRIMARY KEY ("userId");


--
-- Name: __chart__test_grouped PK_f4a2b175d308695af30d4293272; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__test_grouped
    ADD CONSTRAINT "PK_f4a2b175d308695af30d4293272" PRIMARY KEY (id);


--
-- Name: __chart__drive PK_f96bc548a765cd4b3b354221ce7; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__drive
    ADD CONSTRAINT "PK_f96bc548a765cd4b3b354221ce7" PRIMARY KEY (id);


--
-- Name: antenna_note PK_fb28d94d0989a3872df19fd6ef8; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna_note
    ADD CONSTRAINT "PK_fb28d94d0989a3872df19fd6ef8" PRIMARY KEY (id);


--
-- Name: password_reset_request PK_fcf4b02eae1403a2edaf87fd074; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.password_reset_request
    ADD CONSTRAINT "PK_fcf4b02eae1403a2edaf87fd074" PRIMARY KEY (id);


--
-- Name: poll_vote PK_fd002d371201c472490ba89c6a0; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.poll_vote
    ADD CONSTRAINT "PK_fd002d371201c472490ba89c6a0" PRIMARY KEY (id);


--
-- Name: user REL_58f5c71eaab331645112cf8cfa; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "REL_58f5c71eaab331645112cf8cfa" UNIQUE ("avatarId");


--
-- Name: user REL_afc64b53f8db3707ceb34eb28e; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "REL_afc64b53f8db3707ceb34eb28e" UNIQUE ("bannerId");


--
-- Name: __chart__active_users UQ_0ad37b7ef50f4ddc84363d7ccca; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__active_users
    ADD CONSTRAINT "UQ_0ad37b7ef50f4ddc84363d7ccca" UNIQUE (date);


--
-- Name: __chart_day__drive UQ_0b60ebb3aa0065f10b0616c1171; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__drive
    ADD CONSTRAINT "UQ_0b60ebb3aa0065f10b0616c1171" UNIQUE (date);


--
-- Name: __chart__drive UQ_13565815f618a1ff53886c5b28a; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__drive
    ADD CONSTRAINT "UQ_13565815f618a1ff53886c5b28a" UNIQUE (date);


--
-- Name: __chart_day__notes UQ_1a527b423ad0858a1af5a056d43; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__notes
    ADD CONSTRAINT "UQ_1a527b423ad0858a1af5a056d43" UNIQUE (date);


--
-- Name: __chart__per_user_reaction UQ_229a41ad465f9205f1f57032910; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_reaction
    ADD CONSTRAINT "UQ_229a41ad465f9205f1f57032910" UNIQUE (date, "group");


--
-- Name: __chart__hashtag UQ_25a97c02003338124b2b75fdbc8; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__hashtag
    ADD CONSTRAINT "UQ_25a97c02003338124b2b75fdbc8" UNIQUE (date, "group");


--
-- Name: __chart__per_user_drive UQ_30bf67687f483ace115c5ca6429; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_drive
    ADD CONSTRAINT "UQ_30bf67687f483ace115c5ca6429" UNIQUE (date, "group");


--
-- Name: __chart__federation UQ_36cb699c49580d4e6c2e6159f97; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__federation
    ADD CONSTRAINT "UQ_36cb699c49580d4e6c2e6159f97" UNIQUE (date);


--
-- Name: __chart__instance UQ_39ee857ab2f23493037c6b66311; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__instance
    ADD CONSTRAINT "UQ_39ee857ab2f23493037c6b66311" UNIQUE (date, "group");


--
-- Name: __chart__notes UQ_42eb716a37d381cdf566192b2be; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__notes
    ADD CONSTRAINT "UQ_42eb716a37d381cdf566192b2be" UNIQUE (date);


--
-- Name: __chart__per_user_notes UQ_5048e9daccbbbc6d567bb142d34; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_notes
    ADD CONSTRAINT "UQ_5048e9daccbbbc6d567bb142d34" UNIQUE (date, "group");


--
-- Name: __chart_day__federation UQ_617a8fe225a6e701d89e02d2c74; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__federation
    ADD CONSTRAINT "UQ_617a8fe225a6e701d89e02d2c74" UNIQUE (date);


--
-- Name: __chart_day__per_user_drive UQ_62aa5047b5aec92524f24c701d7; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_drive
    ADD CONSTRAINT "UQ_62aa5047b5aec92524f24c701d7" UNIQUE (date, "group");


--
-- Name: user_profile UQ_6dc44f1ceb65b1e72bacef2ca27; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_profile
    ADD CONSTRAINT "UQ_6dc44f1ceb65b1e72bacef2ca27" UNIQUE ("pinnedPageId");


--
-- Name: __chart__users UQ_845254b3eaf708ae8a6cac30265; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__users
    ADD CONSTRAINT "UQ_845254b3eaf708ae8a6cac30265" UNIQUE (date);


--
-- Name: __chart_day__network UQ_8bfa548c2b31f9e07db113773ee; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__network
    ADD CONSTRAINT "UQ_8bfa548c2b31f9e07db113773ee" UNIQUE (date);


--
-- Name: __chart_day__hashtag UQ_8f589cf056ff51f09d6096f6450; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__hashtag
    ADD CONSTRAINT "UQ_8f589cf056ff51f09d6096f6450" UNIQUE (date, "group");


--
-- Name: __chart__network UQ_a1efd3e0048a5f2793a47360dc6; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__network
    ADD CONSTRAINT "UQ_a1efd3e0048a5f2793a47360dc6" UNIQUE (date);


--
-- Name: user UQ_a854e557b1b14814750c7c7b0c9; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "UQ_a854e557b1b14814750c7c7b0c9" UNIQUE (token);


--
-- Name: __chart__per_user_following UQ_b77d4dd9562c3a899d9a286fcd7; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart__per_user_following
    ADD CONSTRAINT "UQ_b77d4dd9562c3a899d9a286fcd7" UNIQUE (date, "group");


--
-- Name: __chart_day__per_user_notes UQ_c5545d4b31cdc684034e33b81c3; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_notes
    ADD CONSTRAINT "UQ_c5545d4b31cdc684034e33b81c3" UNIQUE (date, "group");


--
-- Name: __chart_day__users UQ_cad6e07c20037f31cdba8a350c3; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__users
    ADD CONSTRAINT "UQ_cad6e07c20037f31cdba8a350c3" UNIQUE (date);


--
-- Name: __chart_day__per_user_reaction UQ_d54b653660d808b118e36c184c0; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_reaction
    ADD CONSTRAINT "UQ_d54b653660d808b118e36c184c0" UNIQUE (date, "group");


--
-- Name: __chart_day__active_users UQ_d5954f3df5e5e3bdfc3c03f3906; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__active_users
    ADD CONSTRAINT "UQ_d5954f3df5e5e3bdfc3c03f3906" UNIQUE (date);


--
-- Name: __chart_day__per_user_following UQ_e4849a3231f38281280ea4c0eee; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__per_user_following
    ADD CONSTRAINT "UQ_e4849a3231f38281280ea4c0eee" UNIQUE (date, "group");


--
-- Name: __chart_day__instance UQ_fea7c0278325a1a2492f2d6acbf; Type: CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.__chart_day__instance
    ADD CONSTRAINT "UQ_fea7c0278325a1a2492f2d6acbf" UNIQUE (date, "group");


--
-- Name: IDX_00ceffb0cdc238b3233294f08f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_00ceffb0cdc238b3233294f08f" ON public.drive_folder USING btree ("parentId");


--
-- Name: IDX_01f4581f114e0ebd2bbb876f0b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_01f4581f114e0ebd2bbb876f0b" ON public.note_reaction USING btree ("createdAt");


--
-- Name: IDX_02878d441ceae15ce060b73daf; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_02878d441ceae15ce060b73daf" ON public.drive_folder USING btree ("createdAt");


--
-- Name: IDX_03e7028ab8388a3f5e3ce2a861; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_03e7028ab8388a3f5e3ce2a861" ON public.note_watching USING btree ("noteId");


--
-- Name: IDX_048a757923ed8b157e9895da53; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_048a757923ed8b157e9895da53" ON public.app USING btree ("createdAt");


--
-- Name: IDX_04cc96756f89d0b7f9473e8cdf; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_04cc96756f89d0b7f9473e8cdf" ON public.abuse_user_report USING btree ("reporterId");


--
-- Name: IDX_05cca34b985d1b8edc1d1e28df; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_05cca34b985d1b8edc1d1e28df" ON public.gallery_post USING btree (tags);


--
-- Name: IDX_0610ebcfcfb4a18441a9bcdab2; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0610ebcfcfb4a18441a9bcdab2" ON public.poll USING btree ("userId");


--
-- Name: IDX_0627125f1a8a42c9a1929edb55; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0627125f1a8a42c9a1929edb55" ON public.blocking USING btree ("blockerId");


--
-- Name: IDX_080ab397c379af09b9d2169e5b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_080ab397c379af09b9d2169e5b" ON public.notification USING btree ("isRead");


--
-- Name: IDX_094b86cd36bb805d1aa1e8cc9a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_094b86cd36bb805d1aa1e8cc9a" ON public.channel USING btree ("usersCount");


--
-- Name: IDX_0a72bdfcdb97c0eca11fe7ecad; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0a72bdfcdb97c0eca11fe7ecad" ON public.registry_item USING btree (domain);


--
-- Name: IDX_0ad37b7ef50f4ddc84363d7ccc; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0ad37b7ef50f4ddc84363d7ccc" ON public.__chart__active_users USING btree (date);


--
-- Name: IDX_0b03cbcd7e6a7ce068efa8ecc2; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0b03cbcd7e6a7ce068efa8ecc2" ON public.hashtag USING btree ("attachedRemoteUsersCount");


--
-- Name: IDX_0b575fa9a4cfe638a925949285; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0b575fa9a4cfe638a925949285" ON public.password_reset_request USING btree (token);


--
-- Name: IDX_0b60ebb3aa0065f10b0616c117; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0b60ebb3aa0065f10b0616c117" ON public.__chart_day__drive USING btree (date);


--
-- Name: IDX_0c44bf4f680964145f2a68a341; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0c44bf4f680964145f2a68a341" ON public.hashtag USING btree ("attachedLocalUsersCount");


--
-- Name: IDX_0d7718e562dcedd0aa5cf2c9f7; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0d7718e562dcedd0aa5cf2c9f7" ON public.user_security_key USING btree ("publicKey");


--
-- Name: IDX_0d775946662d2575dfd2068a5f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0d775946662d2575dfd2068a5f" ON public.antenna_note USING btree ("antennaId");


--
-- Name: IDX_0d9a1738f2cf7f3b1c3334dfab; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0d9a1738f2cf7f3b1c3334dfab" ON public.relay USING btree (inbox);


--
-- Name: IDX_0e206cec573f1edff4a3062923; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0e206cec573f1edff4a3062923" ON public.hashtag USING btree ("mentionedLocalUsersCount");


--
-- Name: IDX_0e43068c3f92cab197c3d3cd86; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0e43068c3f92cab197c3d3cd86" ON public.channel_following USING btree ("followeeId");


--
-- Name: IDX_0e61efab7f88dbb79c9166dbb4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0e61efab7f88dbb79c9166dbb4" ON public.page_like USING btree ("userId");


--
-- Name: IDX_0f4fb9ad355f3effff221ef245; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0f4fb9ad355f3effff221ef245" ON public.note_favorite USING btree ("userId", "noteId");


--
-- Name: IDX_0f58c11241e649d2a638a8de94; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0f58c11241e649d2a638a8de94" ON public.channel USING btree ("notesCount");


--
-- Name: IDX_0fb627e1c2f753262a74f0562d; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_0fb627e1c2f753262a74f0562d" ON public.poll_vote USING btree ("createdAt");


--
-- Name: IDX_0ff69e8dfa9fe31bb4a4660f59; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_0ff69e8dfa9fe31bb4a4660f59" ON public.registration_ticket USING btree (code);


--
-- Name: IDX_1039988afa3bf991185b277fe0; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_1039988afa3bf991185b277fe0" ON public.user_group_invite USING btree ("userId");


--
-- Name: IDX_1129c2ef687fc272df040bafaa; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_1129c2ef687fc272df040bafaa" ON public.ad USING btree ("createdAt");


--
-- Name: IDX_118ec703e596086fc4515acb39; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_118ec703e596086fc4515acb39" ON public.announcement USING btree ("createdAt");


--
-- Name: IDX_11e71f2511589dcc8a4d3214f9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_11e71f2511589dcc8a4d3214f9" ON public.channel_following USING btree ("createdAt");


--
-- Name: IDX_12c01c0d1a79f77d9f6c15fadd; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_12c01c0d1a79f77d9f6c15fadd" ON public.follow_request USING btree ("followeeId");


--
-- Name: IDX_13565815f618a1ff53886c5b28; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_13565815f618a1ff53886c5b28" ON public.__chart__drive USING btree (date);


--
-- Name: IDX_13761f64257f40c5636d0ff95e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_13761f64257f40c5636d0ff95e" ON public.note_reaction USING btree ("userId");


--
-- Name: IDX_153536c67d05e9adb24e99fc2b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_153536c67d05e9adb24e99fc2b" ON public.note USING btree (uri);


--
-- Name: IDX_16effb2e888f6763673b579f80; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_16effb2e888f6763673b579f80" ON public.__chart__test_unique USING btree (date) WHERE ("group" IS NULL);


--
-- Name: IDX_171e64971c780ebd23fae140bb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_171e64971c780ebd23fae140bb" ON public.user_publickey USING btree ("keyId");


--
-- Name: IDX_17cb3553c700a4985dff5a30ff; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_17cb3553c700a4985dff5a30ff" ON public.note USING btree ("replyId");


--
-- Name: IDX_1a165c68a49d08f11caffbd206; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_1a165c68a49d08f11caffbd206" ON public.gallery_post USING btree ("likedCount");


--
-- Name: IDX_1a527b423ad0858a1af5a056d4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_1a527b423ad0858a1af5a056d4" ON public.__chart_day__notes USING btree (date);


--
-- Name: IDX_1eb9d9824a630321a29fd3b290; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_1eb9d9824a630321a29fd3b290" ON public.muting USING btree ("muterId", "muteeId");


--
-- Name: IDX_20e30aa35180e317e133d75316; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_20e30aa35180e317e133d75316" ON public.user_group USING btree ("createdAt");


--
-- Name: IDX_2133ef8317e4bdb839c0dcbf13; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_2133ef8317e4bdb839c0dcbf13" ON public.page USING btree ("userId", name);


--
-- Name: IDX_229a41ad465f9205f1f5703291; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_229a41ad465f9205f1f5703291" ON public.__chart__per_user_reaction USING btree (date, "group");


--
-- Name: IDX_22baca135bb8a3ea1a83d13df3; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_22baca135bb8a3ea1a83d13df3" ON public.registry_item USING btree (scope);


--
-- Name: IDX_24e0042143a18157b234df186c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_24e0042143a18157b234df186c" ON public.following USING btree ("followeeId");


--
-- Name: IDX_25a97c02003338124b2b75fdbc; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_25a97c02003338124b2b75fdbc" ON public.__chart__hashtag USING btree (date, "group");


--
-- Name: IDX_25b1dd384bec391b07b74b861c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_25b1dd384bec391b07b74b861c" ON public.note_unread USING btree ("isMentioned");


--
-- Name: IDX_25dfc71b0369b003a4cd434d0b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_25dfc71b0369b003a4cd434d0b" ON public.note USING btree ("attachedFileTypes");


--
-- Name: IDX_2710a55f826ee236ea1a62698f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2710a55f826ee236ea1a62698f" ON public.hashtag USING btree ("mentionedUsersCount");


--
-- Name: IDX_2882b8a1a07c7d281a98b6db16; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_2882b8a1a07c7d281a98b6db16" ON public.promo_read USING btree ("userId", "noteId");


--
-- Name: IDX_29c11c7deb06615076f8c95b80; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_29c11c7deb06615076f8c95b80" ON public.note_thread_muting USING btree ("userId");


--
-- Name: IDX_29e8c1d579af54d4232939f994; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_29e8c1d579af54d4232939f994" ON public.note_unread USING btree ("noteUserId");


--
-- Name: IDX_29ef80c6f13bcea998447fce43; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_29ef80c6f13bcea998447fce43" ON public.channel USING btree ("lastNotedAt");


--
-- Name: IDX_2b15aaf4a0dc5be3499af7ab6a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2b15aaf4a0dc5be3499af7ab6a" ON public.abuse_user_report USING btree (resolved);


--
-- Name: IDX_2b5ec6c574d6802c94c80313fb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2b5ec6c574d6802c94c80313fb" ON public.clip USING btree ("userId");


--
-- Name: IDX_2c308dbdc50d94dc625670055f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2c308dbdc50d94dc625670055f" ON public.signin USING btree ("userId");


--
-- Name: IDX_2c4be03b446884f9e9c502135b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2c4be03b446884f9e9c502135b" ON public.messaging_message USING btree ("groupId");


--
-- Name: IDX_2cd3b2a6b4cf0b910b260afe08; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2cd3b2a6b4cf0b910b260afe08" ON public.instance USING btree ("caughtAt");


--
-- Name: IDX_2cd4a2743a99671308f5417759; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2cd4a2743a99671308f5417759" ON public.blocking USING btree ("blockeeId");


--
-- Name: IDX_2da24ce20ad209f1d9dc032457; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_2da24ce20ad209f1d9dc032457" ON public.ad USING btree ("expiresAt");


--
-- Name: IDX_2e230dd45a10e671d781d99f3e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_2e230dd45a10e671d781d99f3e" ON public.channel_following USING btree ("followerId", "followeeId");


--
-- Name: IDX_307be5f1d1252e0388662acb96; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_307be5f1d1252e0388662acb96" ON public.following USING btree ("followerId", "followeeId");


--
-- Name: IDX_30bf67687f483ace115c5ca642; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_30bf67687f483ace115c5ca642" ON public.__chart__per_user_drive USING btree (date, "group");


--
-- Name: IDX_318cdf42a9cfc11f479bd802bb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_318cdf42a9cfc11f479bd802bb" ON public.note_watching USING btree ("createdAt");


--
-- Name: IDX_3252a5df8d5bbd16b281f7799e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3252a5df8d5bbd16b281f7799e" ON public."user" USING btree (host);


--
-- Name: IDX_335a0bf3f904406f9ef3dd51c2; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_335a0bf3f904406f9ef3dd51c2" ON public.antenna_note USING btree ("noteId", "antennaId");


--
-- Name: IDX_33f33cc8ef29d805a97ff4628b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_33f33cc8ef29d805a97ff4628b" ON public.notification USING btree (type);


--
-- Name: IDX_34500da2e38ac393f7bb6b299c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_34500da2e38ac393f7bb6b299c" ON public.instance USING btree ("isSuspended");


--
-- Name: IDX_347fec870eafea7b26c8a73bac; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_347fec870eafea7b26c8a73bac" ON public.hashtag USING btree (name);


--
-- Name: IDX_36cb699c49580d4e6c2e6159f9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_36cb699c49580d4e6c2e6159f9" ON public.__chart__federation USING btree (date);


--
-- Name: IDX_37bb9a1b4585f8a3beb24c62d6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_37bb9a1b4585f8a3beb24c62d6" ON public.drive_file USING btree (md5);


--
-- Name: IDX_39ee857ab2f23493037c6b6631; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_39ee857ab2f23493037c6b6631" ON public.__chart__instance USING btree (date, "group");


--
-- Name: IDX_3b25402709dd9882048c2bbade; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3b25402709dd9882048c2bbade" ON public.reversi_matching USING btree ("parentId");


--
-- Name: IDX_3b4e96eec8d36a8bbb9d02aa71; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3b4e96eec8d36a8bbb9d02aa71" ON public.notification USING btree ("notifierId");


--
-- Name: IDX_3befe6f999c86aff06eb0257b4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3befe6f999c86aff06eb0257b4" ON public.user_profile USING btree ("enableWordMute");


--
-- Name: IDX_3c601b70a1066d2c8b517094cb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3c601b70a1066d2c8b517094cb" ON public.notification USING btree ("notifieeId");


--
-- Name: IDX_3ca50563facd913c425e7a89ee; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3ca50563facd913c425e7a89ee" ON public.gallery_post USING btree ("fileIds");


--
-- Name: IDX_3d6b372788ab01be58853003c9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3d6b372788ab01be58853003c9" ON public.user_group USING btree ("userId");


--
-- Name: IDX_3f5b0899ef90527a3462d7c2cb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_3f5b0899ef90527a3462d7c2cb" ON public.app USING btree ("userId");


--
-- Name: IDX_410cd649884b501c02d6e72738; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_410cd649884b501c02d6e72738" ON public.user_note_pining USING btree ("userId", "noteId");


--
-- Name: IDX_42eb716a37d381cdf566192b2b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_42eb716a37d381cdf566192b2b" ON public.__chart__notes USING btree (date);


--
-- Name: IDX_44499765eec6b5489d72c4253b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_44499765eec6b5489d72c4253b" ON public.note_watching USING btree ("noteUserId");


--
-- Name: IDX_45145e4953780f3cd5656f0ea6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_45145e4953780f3cd5656f0ea6" ON public.note_reaction USING btree ("noteId");


--
-- Name: IDX_47efb914aed1f72dd39a306c7b; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_47efb914aed1f72dd39a306c7b" ON public.attestation_challenge USING btree (challenge);


--
-- Name: IDX_47f4b1892f5d6ba8efb3057d81; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_47f4b1892f5d6ba8efb3057d81" ON public.note_favorite USING btree ("userId");


--
-- Name: IDX_4bb7fd4a34492ae0e6cc8d30ac; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_4bb7fd4a34492ae0e6cc8d30ac" ON public.password_reset_request USING btree ("userId");


--
-- Name: IDX_4c02d38a976c3ae132228c6fce; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_4c02d38a976c3ae132228c6fce" ON public.hashtag USING btree ("mentionedRemoteUsersCount");


--
-- Name: IDX_4ce6fb9c70529b4c8ac46c9bfa; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_4ce6fb9c70529b4c8ac46c9bfa" ON public.page_like USING btree ("userId", "pageId");


--
-- Name: IDX_4e5c4c99175638ec0761714ab0; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_4e5c4c99175638ec0761714ab0" ON public.user_pending USING btree (code);


--
-- Name: IDX_4ebbf7f93cdc10e8d1ef2fc6cd; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_4ebbf7f93cdc10e8d1ef2fc6cd" ON public.abuse_user_report USING btree ("targetUserHost");


--
-- Name: IDX_4f4d35e1256c84ae3d1f0eab10; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_4f4d35e1256c84ae3d1f0eab10" ON public.emoji USING btree (name, host);


--
-- Name: IDX_5048e9daccbbbc6d567bb142d3; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_5048e9daccbbbc6d567bb142d3" ON public.__chart__per_user_notes USING btree (date, "group");


--
-- Name: IDX_50bd7164c5b78f1f4a42c4d21f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_50bd7164c5b78f1f4a42c4d21f" ON public.poll_vote USING btree ("userId", "noteId", choice);


--
-- Name: IDX_51c063b6a133a9cb87145450f5; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_51c063b6a133a9cb87145450f5" ON public.note USING btree ("fileIds");


--
-- Name: IDX_52ccc804d7c69037d558bac4c9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_52ccc804d7c69037d558bac4c9" ON public.note USING btree ("renoteId");


--
-- Name: IDX_5377c307783fce2b6d352e1203; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_5377c307783fce2b6d352e1203" ON public.messaging_message USING btree ("userId");


--
-- Name: IDX_54ebcb6d27222913b908d56fd8; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_54ebcb6d27222913b908d56fd8" ON public.note USING btree (mentions);


--
-- Name: IDX_55720b33a61a7c806a8215b825; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_55720b33a61a7c806a8215b825" ON public.drive_file USING btree ("userId", "folderId", id);


--
-- Name: IDX_56b0166d34ddae49d8ef7610bb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_56b0166d34ddae49d8ef7610bb" ON public.note_unread USING btree ("userId");


--
-- Name: IDX_582f8fab771a9040a12961f3e7; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_582f8fab771a9040a12961f3e7" ON public.following USING btree ("createdAt");


--
-- Name: IDX_5900e907bb46516ddf2871327c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_5900e907bb46516ddf2871327c" ON public.emoji USING btree (host);


--
-- Name: IDX_5b87d9d19127bd5d92026017a7; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_5b87d9d19127bd5d92026017a7" ON public.note USING btree ("userId");


--
-- Name: IDX_5cc8c468090e129857e9fecce5; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_5cc8c468090e129857e9fecce5" ON public.user_group_invitation USING btree ("userGroupId");


--
-- Name: IDX_5deb01ae162d1d70b80d064c27; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_5deb01ae162d1d70b80d064c27" ON public."user" USING btree ("usernameLower", host);


--
-- Name: IDX_603a7b1e7aa0533c6c88e9bfaf; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_603a7b1e7aa0533c6c88e9bfaf" ON public.announcement_read USING btree ("announcementId");


--
-- Name: IDX_605472305f26818cc93d1baaa7; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_605472305f26818cc93d1baaa7" ON public.user_list_joining USING btree ("userListId");


--
-- Name: IDX_617a8fe225a6e701d89e02d2c7; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_617a8fe225a6e701d89e02d2c7" ON public.__chart_day__federation USING btree (date);


--
-- Name: IDX_62aa5047b5aec92524f24c701d; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_62aa5047b5aec92524f24c701d" ON public.__chart_day__per_user_drive USING btree (date, "group");


--
-- Name: IDX_62cb09e1129f6ec024ef66e183; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_62cb09e1129f6ec024ef66e183" ON public.auth_session USING btree (token);


--
-- Name: IDX_636e977ff90b23676fb5624b25; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_636e977ff90b23676fb5624b25" ON public.muted_note USING btree (reason);


--
-- Name: IDX_6446c571a0e8d0f05f01c78909; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_6446c571a0e8d0f05f01c78909" ON public.antenna USING btree ("userId");


--
-- Name: IDX_64c327441248bae40f7d92f34f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_64c327441248bae40f7d92f34f" ON public.access_token USING btree (hash);


--
-- Name: IDX_6516c5a6f3c015b4eed39978be; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_6516c5a6f3c015b4eed39978be" ON public.following USING btree ("followerId");


--
-- Name: IDX_66d2bd2ee31d14bcc23069a89f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_66d2bd2ee31d14bcc23069a89f" ON public.poll_vote USING btree ("userId");


--
-- Name: IDX_67dc758bc0566985d1b3d39986; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_67dc758bc0566985d1b3d39986" ON public.user_group_joining USING btree ("userGroupId");


--
-- Name: IDX_6a57f051d82c6d4036c141e107; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_6a57f051d82c6d4036c141e107" ON public.note_unread USING btree ("noteChannelId");


--
-- Name: IDX_6d8084ec9496e7334a4602707e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_6d8084ec9496e7334a4602707e" ON public.channel_following USING btree ("followerId");


--
-- Name: IDX_6fc0ec357d55a18646262fdfff; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_6fc0ec357d55a18646262fdfff" ON public.clip_note USING btree ("noteId", "clipId");


--
-- Name: IDX_70ab9786313d78e4201d81cdb8; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_70ab9786313d78e4201d81cdb8" ON public.muted_note USING btree ("noteId");


--
-- Name: IDX_70ba8f6af34bc924fc9e12adb8; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_70ba8f6af34bc924fc9e12adb8" ON public.access_token USING btree (token);


--
-- Name: IDX_7125a826ab192eb27e11d358a5; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_7125a826ab192eb27e11d358a5" ON public.note USING btree ("userHost");


--
-- Name: IDX_71cb7b435b7c0d4843317e7e16; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_71cb7b435b7c0d4843317e7e16" ON public.channel USING btree ("createdAt");


--
-- Name: IDX_78787741f9010886796f2320a4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_78787741f9010886796f2320a4" ON public.user_group_invite USING btree ("userId", "userGroupId");


--
-- Name: IDX_796a8c03959361f97dc2be1d5c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_796a8c03959361f97dc2be1d5c" ON public.note USING btree ("visibleUserIds");


--
-- Name: IDX_7fa20a12319c7f6dc3aed98c0a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_7fa20a12319c7f6dc3aed98c0a" ON public.poll USING btree ("userHost");


--
-- Name: IDX_80ca6e6ef65fb9ef34ea8c90f4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_80ca6e6ef65fb9ef34ea8c90f4" ON public."user" USING btree ("updatedAt");


--
-- Name: IDX_8125f950afd3093acb10d2db8a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_8125f950afd3093acb10d2db8a" ON public.channel_note_pining USING btree ("channelId");


--
-- Name: IDX_823bae55bd81b3be6e05cff438; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_823bae55bd81b3be6e05cff438" ON public.channel USING btree ("userId");


--
-- Name: IDX_8288151386172b8109f7239ab2; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_8288151386172b8109f7239ab2" ON public.announcement_read USING btree ("userId");


--
-- Name: IDX_83f0862e9bae44af52ced7099e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_83f0862e9bae44af52ced7099e" ON public.promo_note USING btree ("userId");


--
-- Name: IDX_845254b3eaf708ae8a6cac3026; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_845254b3eaf708ae8a6cac3026" ON public.__chart__users USING btree (date);


--
-- Name: IDX_860fa6f6c7df5bb887249fba22; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_860fa6f6c7df5bb887249fba22" ON public.drive_file USING btree ("userId");


--
-- Name: IDX_88937d94d7443d9a99a76fa5c0; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_88937d94d7443d9a99a76fa5c0" ON public.note USING btree (tags);


--
-- Name: IDX_89a29c9237b8c3b6b3cbb4cb30; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_89a29c9237b8c3b6b3cbb4cb30" ON public.note_unread USING btree ("isSpecified");


--
-- Name: IDX_8bfa548c2b31f9e07db113773e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_8bfa548c2b31f9e07db113773e" ON public.__chart_day__network USING btree (date);


--
-- Name: IDX_8d5afc98982185799b160e10eb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_8d5afc98982185799b160e10eb" ON public.instance USING btree (host);


--
-- Name: IDX_8f1a239bd077c8864a20c62c2c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_8f1a239bd077c8864a20c62c2c" ON public.gallery_post USING btree ("createdAt");


--
-- Name: IDX_8f589cf056ff51f09d6096f645; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_8f589cf056ff51f09d6096f645" ON public.__chart_day__hashtag USING btree (date, "group");


--
-- Name: IDX_8fd5215095473061855ceb948c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_8fd5215095473061855ceb948c" ON public.gallery_like USING btree ("userId");


--
-- Name: IDX_90148bbc2bf0854428786bfc15; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_90148bbc2bf0854428786bfc15" ON public.page USING btree ("visibleUserIds");


--
-- Name: IDX_90f7da835e4c10aca6853621e1; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_90f7da835e4c10aca6853621e1" ON public.user_list_joining USING btree ("userId", "userListId");


--
-- Name: IDX_924fa71815cfa3941d003702a0; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_924fa71815cfa3941d003702a0" ON public.announcement_read USING btree ("userId", "announcementId");


--
-- Name: IDX_92779627994ac79277f070c91e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_92779627994ac79277f070c91e" ON public.drive_file USING btree ("userHost");


--
-- Name: IDX_93060675b4a79a577f31d260c6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_93060675b4a79a577f31d260c6" ON public.muting USING btree ("muterId");


--
-- Name: IDX_9657d55550c3d37bfafaf7d4b0; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_9657d55550c3d37bfafaf7d4b0" ON public.promo_read USING btree ("userId");


--
-- Name: IDX_97754ca6f2baff9b4abb7f853d; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_97754ca6f2baff9b4abb7f853d" ON public.sw_subscription USING btree ("userId");


--
-- Name: IDX_985b836dddd8615e432d7043dd; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_985b836dddd8615e432d7043dd" ON public.gallery_post USING btree ("userId");


--
-- Name: IDX_98a1bc5cb30dfd159de056549f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_98a1bc5cb30dfd159de056549f" ON public.blocking USING btree ("blockerId", "blockeeId");


--
-- Name: IDX_9937ea48d7ae97ffb4f3f063a4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_9937ea48d7ae97ffb4f3f063a4" ON public.antenna_note USING btree (read);


--
-- Name: IDX_9949557d0e1b2c19e5344c171e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_9949557d0e1b2c19e5344c171e" ON public.access_token USING btree ("userId");


--
-- Name: IDX_NOTE_MENTIONS; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_NOTE_MENTIONS" ON public.note USING gin (mentions);


--
-- Name: IDX_NOTE_TAGS; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_NOTE_TAGS" ON public.note USING gin (tags);


--
-- Name: IDX_NOTE_VISIBLE_USER_IDS; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_NOTE_VISIBLE_USER_IDS" ON public.note USING gin ("visibleUserIds");


--
-- Name: IDX_a012eaf5c87c65da1deb5fdbfa; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a012eaf5c87c65da1deb5fdbfa" ON public.clip_note USING btree ("noteId");


--
-- Name: IDX_a08ad074601d204e0f69da9a95; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a08ad074601d204e0f69da9a95" ON public.moderation_log USING btree ("userId");


--
-- Name: IDX_a0cd75442dd10d0643a17c4a49; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a0cd75442dd10d0643a17c4a49" ON public.__chart__test_unique USING btree (date, "group");


--
-- Name: IDX_a1efd3e0048a5f2793a47360dc; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a1efd3e0048a5f2793a47360dc" ON public.__chart__network USING btree (date);


--
-- Name: IDX_a27b942a0d6dcff90e3ee9b5e8; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a27b942a0d6dcff90e3ee9b5e8" ON public."user" USING btree ("usernameLower");


--
-- Name: IDX_a319e5dbf47e8a17497623beae; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a319e5dbf47e8a17497623beae" ON public.__chart__test USING btree (date, "group");


--
-- Name: IDX_a40b8df8c989d7db937ea27cf6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a40b8df8c989d7db937ea27cf6" ON public.drive_file USING btree (type);


--
-- Name: IDX_a42c93c69989ce1d09959df4cf; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a42c93c69989ce1d09959df4cf" ON public.note_watching USING btree ("userId", "noteId");


--
-- Name: IDX_a7eba67f8b3fa27271e85d2e26; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a7eba67f8b3fa27271e85d2e26" ON public.drive_file USING btree ("isSensitive");


--
-- Name: IDX_a7fd92dd6dc519e6fb435dd108; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_a7fd92dd6dc519e6fb435dd108" ON public.follow_request USING btree ("followerId");


--
-- Name: IDX_a854e557b1b14814750c7c7b0c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a854e557b1b14814750c7c7b0c" ON public."user" USING btree (token);


--
-- Name: IDX_a8c6bfd637d3f1d67a27c48e27; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_a8c6bfd637d3f1d67a27c48e27" ON public.muted_note USING btree ("noteId", "userId");


--
-- Name: IDX_ad0c221b25672daf2df320a817; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_ad0c221b25672daf2df320a817" ON public.note_reaction USING btree ("userId", "noteId");


--
-- Name: IDX_ae1d917992dd0c9d9bbdad06c4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_ae1d917992dd0c9d9bbdad06c4" ON public.page USING btree ("userId");


--
-- Name: IDX_ae7aab18a2641d3e5f25e0c4ea; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_ae7aab18a2641d3e5f25e0c4ea" ON public.note_thread_muting USING btree ("userId", "threadId");


--
-- Name: IDX_aecfbd5ef60374918e63ee95fa; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_aecfbd5ef60374918e63ee95fa" ON public.poll_vote USING btree ("noteId");


--
-- Name: IDX_af639b066dfbca78b01a920f8a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_af639b066dfbca78b01a920f8a" ON public.page USING btree ("updatedAt");


--
-- Name: IDX_b0134ec406e8d09a540f818288; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b0134ec406e8d09a540f818288" ON public.note_watching USING btree ("userId");


--
-- Name: IDX_b11a5e627c41d4dc3170f1d370; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b11a5e627c41d4dc3170f1d370" ON public.notification USING btree ("createdAt");


--
-- Name: IDX_b14489029e4b3aaf4bba5fb524; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_b14489029e4b3aaf4bba5fb524" ON public.__chart__test_grouped USING btree (date, "group");


--
-- Name: IDX_b37dafc86e9af007e3295c2781; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b37dafc86e9af007e3295c2781" ON public.emoji USING btree (name);


--
-- Name: IDX_b46ec40746efceac604142be1c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b46ec40746efceac604142be1c" ON public.reversi_game USING btree ("createdAt");


--
-- Name: IDX_b604d92d6c7aec38627f6eaf16; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b604d92d6c7aec38627f6eaf16" ON public.reversi_matching USING btree ("createdAt");


--
-- Name: IDX_b77d4dd9562c3a899d9a286fcd; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_b77d4dd9562c3a899d9a286fcd" ON public.__chart__per_user_following USING btree (date, "group");


--
-- Name: IDX_b7fcefbdd1c18dce86687531f9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b7fcefbdd1c18dce86687531f9" ON public.user_list USING btree ("userId");


--
-- Name: IDX_b82c19c08afb292de4600d99e4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b82c19c08afb292de4600d99e4" ON public.page USING btree (name);


--
-- Name: IDX_b9a354f7941c1e779f3b33aea6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_b9a354f7941c1e779f3b33aea6" ON public.blocking USING btree ("createdAt");


--
-- Name: IDX_bb90d1956dafc4068c28aa7560; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_bb90d1956dafc4068c28aa7560" ON public.drive_file USING btree ("folderId");


--
-- Name: IDX_bd0397be22147e17210940e125; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_bd0397be22147e17210940e125" ON public.antenna_note USING btree ("noteId");


--
-- Name: IDX_be623adaa4c566baf5d29ce0c8; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_be623adaa4c566baf5d29ce0c8" ON public."user" USING btree (uri);


--
-- Name: IDX_bf3a053c07d9fb5d87317c56ee; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_bf3a053c07d9fb5d87317c56ee" ON public.access_token USING btree (session);


--
-- Name: IDX_bfbc6305547539369fe73eb144; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_bfbc6305547539369fe73eb144" ON public.user_group_invitation USING btree ("userId");


--
-- Name: IDX_bfbc6f79ba4007b4ce5097f08d; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_bfbc6f79ba4007b4ce5097f08d" ON public.user_note_pining USING btree ("userId");


--
-- Name: IDX_c426394644267453e76f036926; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_c426394644267453e76f036926" ON public.note_thread_muting USING btree ("threadId");


--
-- Name: IDX_c5545d4b31cdc684034e33b81c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_c5545d4b31cdc684034e33b81c" ON public.__chart_day__per_user_notes USING btree (date, "group");


--
-- Name: IDX_c55b2b7c284d9fef98026fc88e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_c55b2b7c284d9fef98026fc88e" ON public.drive_file USING btree ("webpublicAccessKey");


--
-- Name: IDX_c8dfad3b72196dd1d6b5db168a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_c8dfad3b72196dd1d6b5db168a" ON public.drive_file USING btree ("createdAt");


--
-- Name: IDX_cac14a4e3944454a5ce7daa514; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_cac14a4e3944454a5ce7daa514" ON public.messaging_message USING btree ("recipientId");


--
-- Name: IDX_cad6e07c20037f31cdba8a350c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_cad6e07c20037f31cdba8a350c" ON public.__chart_day__users USING btree (date);


--
-- Name: IDX_d4ebdef929896d6dc4a3c5bb48; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_d4ebdef929896d6dc4a3c5bb48" ON public.note USING btree ("threadId");


--
-- Name: IDX_d54a512b822fac7ed52800f6b4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d54a512b822fac7ed52800f6b4" ON public.follow_request USING btree ("followerId", "followeeId");


--
-- Name: IDX_d54b653660d808b118e36c184c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d54b653660d808b118e36c184c" ON public.__chart_day__per_user_reaction USING btree (date, "group");


--
-- Name: IDX_d57f9030cd3af7f63ffb1c267c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_d57f9030cd3af7f63ffb1c267c" ON public.hashtag USING btree ("attachedUsersCount");


--
-- Name: IDX_d5954f3df5e5e3bdfc3c03f390; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d5954f3df5e5e3bdfc3c03f390" ON public.__chart_day__active_users USING btree (date);


--
-- Name: IDX_d5a1b83c7cab66f167e6888188; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_d5a1b83c7cab66f167e6888188" ON public."user" USING btree ("isExplorable");


--
-- Name: IDX_d844bfc6f3f523a05189076efa; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_d844bfc6f3f523a05189076efa" ON public.user_list_joining USING btree ("userId");


--
-- Name: IDX_d85a184c2540d2deba33daf642; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d85a184c2540d2deba33daf642" ON public.drive_file USING btree ("accessKey");


--
-- Name: IDX_d8e07aa18c2d64e86201601aec; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_d8e07aa18c2d64e86201601aec" ON public.muted_note USING btree ("userId");


--
-- Name: IDX_d908433a4953cc13216cd9c274; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d908433a4953cc13216cd9c274" ON public.note_unread USING btree ("userId", "noteId");


--
-- Name: IDX_d9ecaed8c6dc43f3592c229282; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_d9ecaed8c6dc43f3592c229282" ON public.user_group_joining USING btree ("userId", "userGroupId");


--
-- Name: IDX_da522b4008a9f5d7743b87ad55; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_da522b4008a9f5d7743b87ad55" ON public.__chart__test_grouped USING btree (date) WHERE ("group" IS NULL);


--
-- Name: IDX_dab383a36f3c9db4a0c9b02cf3; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_dab383a36f3c9db4a0c9b02cf3" ON public.__chart__test USING btree (date) WHERE ("group" IS NULL);


--
-- Name: IDX_db2098070b2b5a523c58181f74; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_db2098070b2b5a523c58181f74" ON public.abuse_user_report USING btree ("createdAt");


--
-- Name: IDX_dce530b98e454793dac5ec2f5a; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_dce530b98e454793dac5ec2f5a" ON public.user_profile USING btree ("userHost");


--
-- Name: IDX_df1b5f4099e99fb0bc5eae53b6; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_df1b5f4099e99fb0bc5eae53b6" ON public.gallery_like USING btree ("userId", "postId");


--
-- Name: IDX_e10924607d058004304611a436; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e10924607d058004304611a436" ON public.user_group_invite USING btree ("userGroupId");


--
-- Name: IDX_e11e649824a45d8ed01d597fd9; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e11e649824a45d8ed01d597fd9" ON public."user" USING btree ("createdAt");


--
-- Name: IDX_e21cd3646e52ef9c94aaf17c2e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e21cd3646e52ef9c94aaf17c2e" ON public.messaging_message USING btree ("createdAt");


--
-- Name: IDX_e22bf6bda77b6adc1fd9e75c8c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e22bf6bda77b6adc1fd9e75c8c" ON public.notification USING btree ("appAccessTokenId");


--
-- Name: IDX_e247b23a3c9b45f89ec1299d06; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e247b23a3c9b45f89ec1299d06" ON public.reversi_matching USING btree ("childId");


--
-- Name: IDX_e4849a3231f38281280ea4c0ee; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_e4849a3231f38281280ea4c0ee" ON public.__chart_day__per_user_following USING btree (date, "group");


--
-- Name: IDX_e5848eac4940934e23dbc17581; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e5848eac4940934e23dbc17581" ON public.drive_file USING btree (uri);


--
-- Name: IDX_e637cba4dc4410218c4251260e; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e637cba4dc4410218c4251260e" ON public.note_unread USING btree ("noteId");


--
-- Name: IDX_e74022ce9a074b3866f70e0d27; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_e74022ce9a074b3866f70e0d27" ON public.drive_file USING btree ("thumbnailAccessKey");


--
-- Name: IDX_e7c0567f5261063592f022e9b5; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_e7c0567f5261063592f022e9b5" ON public.note USING btree ("createdAt");


--
-- Name: IDX_e9793f65f504e5a31fbaedbf2f; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_e9793f65f504e5a31fbaedbf2f" ON public.user_group_invitation USING btree ("userId", "userGroupId");


--
-- Name: IDX_ebe99317bbbe9968a0c6f579ad; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_ebe99317bbbe9968a0c6f579ad" ON public.clip_note USING btree ("clipId");


--
-- Name: IDX_ec96b4fed9dae517e0dbbe0675; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_ec96b4fed9dae517e0dbbe0675" ON public.muting USING btree ("muteeId");


--
-- Name: IDX_f1a461a618fa1755692d0e0d59; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f1a461a618fa1755692d0e0d59" ON public.attestation_challenge USING btree ("userId");


--
-- Name: IDX_f2d744d9a14d0dfb8b96cb7fc5; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f2d744d9a14d0dfb8b96cb7fc5" ON public.gallery_post USING btree ("isSensitive");


--
-- Name: IDX_f36fed37d6d4cdcc68c803cd9c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_f36fed37d6d4cdcc68c803cd9c" ON public.channel_note_pining USING btree ("channelId", "noteId");


--
-- Name: IDX_f3a1b4bd0c7cabba958a0c0b23; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f3a1b4bd0c7cabba958a0c0b23" ON public.user_group_joining USING btree ("userId");


--
-- Name: IDX_f49922d511d666848f250663c4; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f49922d511d666848f250663c4" ON public.app USING btree (secret);


--
-- Name: IDX_f4fc06e49c0171c85f1c48060d; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f4fc06e49c0171c85f1c48060d" ON public.drive_folder USING btree ("userId");


--
-- Name: IDX_f631d37835adb04792e361807c; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f631d37835adb04792e361807c" ON public.gallery_post USING btree ("updatedAt");


--
-- Name: IDX_f86d57fbca33c7a4e6897490cc; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f86d57fbca33c7a4e6897490cc" ON public.muting USING btree ("createdAt");


--
-- Name: IDX_f8d8b93740ad12c4ce8213a199; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_f8d8b93740ad12c4ce8213a199" ON public.abuse_user_report USING btree ("reporterHost");


--
-- Name: IDX_fa99d777623947a5b05f394cae; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_fa99d777623947a5b05f394cae" ON public."user" USING btree (tags);


--
-- Name: IDX_fb9d21ba0abb83223263df6bcb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_fb9d21ba0abb83223263df6bcb" ON public.registry_item USING btree ("userId");


--
-- Name: IDX_fbb4297c927a9b85e9cefa2eb1; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_fbb4297c927a9b85e9cefa2eb1" ON public.page USING btree ("createdAt");


--
-- Name: IDX_fea7c0278325a1a2492f2d6acb; Type: INDEX; Schema: public; Owner: misskey
--

CREATE UNIQUE INDEX "IDX_fea7c0278325a1a2492f2d6acb" ON public.__chart_day__instance USING btree (date, "group");


--
-- Name: IDX_ff9ca3b5f3ee3d0681367a9b44; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_ff9ca3b5f3ee3d0681367a9b44" ON public.user_security_key USING btree ("userId");


--
-- Name: IDX_note_on_channelId_and_id_desc; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_note_on_channelId_and_id_desc" ON public.note USING btree ("channelId", id DESC);


--
-- Name: IDX_seoignmeoprigmkpodgrjmkpormg; Type: INDEX; Schema: public; Owner: misskey
--

CREATE INDEX "IDX_seoignmeoprigmkpodgrjmkpormg" ON public."user" USING btree ("lastActiveDate");


--
-- Name: drive_folder FK_00ceffb0cdc238b3233294f08f2; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_folder
    ADD CONSTRAINT "FK_00ceffb0cdc238b3233294f08f2" FOREIGN KEY ("parentId") REFERENCES public.drive_folder(id) ON DELETE SET NULL;


--
-- Name: note_watching FK_03e7028ab8388a3f5e3ce2a8619; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_watching
    ADD CONSTRAINT "FK_03e7028ab8388a3f5e3ce2a8619" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: abuse_user_report FK_04cc96756f89d0b7f9473e8cdf3; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.abuse_user_report
    ADD CONSTRAINT "FK_04cc96756f89d0b7f9473e8cdf3" FOREIGN KEY ("reporterId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: blocking FK_0627125f1a8a42c9a1929edb552; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.blocking
    ADD CONSTRAINT "FK_0627125f1a8a42c9a1929edb552" FOREIGN KEY ("blockerId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: abuse_user_report FK_08b883dd5fdd6f9c4c1572b36de; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.abuse_user_report
    ADD CONSTRAINT "FK_08b883dd5fdd6f9c4c1572b36de" FOREIGN KEY ("assigneeId") REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: antenna_note FK_0d775946662d2575dfd2068a5f5; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna_note
    ADD CONSTRAINT "FK_0d775946662d2575dfd2068a5f5" FOREIGN KEY ("antennaId") REFERENCES public.antenna(id) ON DELETE CASCADE;


--
-- Name: note_favorite FK_0e00498f180193423c992bc4370; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_favorite
    ADD CONSTRAINT "FK_0e00498f180193423c992bc4370" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: channel_following FK_0e43068c3f92cab197c3d3cd86e; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_following
    ADD CONSTRAINT "FK_0e43068c3f92cab197c3d3cd86e" FOREIGN KEY ("followeeId") REFERENCES public.channel(id) ON DELETE CASCADE;


--
-- Name: page_like FK_0e61efab7f88dbb79c9166dbb48; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page_like
    ADD CONSTRAINT "FK_0e61efab7f88dbb79c9166dbb48" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_group_invite FK_1039988afa3bf991185b277fe03; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invite
    ADD CONSTRAINT "FK_1039988afa3bf991185b277fe03" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: channel_note_pining FK_10b19ef67d297ea9de325cd4502; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_note_pining
    ADD CONSTRAINT "FK_10b19ef67d297ea9de325cd4502" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: user_publickey FK_10c146e4b39b443ede016f6736d; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_publickey
    ADD CONSTRAINT "FK_10c146e4b39b443ede016f6736d" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: follow_request FK_12c01c0d1a79f77d9f6c15fadd2; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.follow_request
    ADD CONSTRAINT "FK_12c01c0d1a79f77d9f6c15fadd2" FOREIGN KEY ("followeeId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note_reaction FK_13761f64257f40c5636d0ff95ee; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_reaction
    ADD CONSTRAINT "FK_13761f64257f40c5636d0ff95ee" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note FK_17cb3553c700a4985dff5a30ff5; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT "FK_17cb3553c700a4985dff5a30ff5" FOREIGN KEY ("replyId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: following FK_24e0042143a18157b234df186c3; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.following
    ADD CONSTRAINT "FK_24e0042143a18157b234df186c3" FOREIGN KEY ("followeeId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note_thread_muting FK_29c11c7deb06615076f8c95b80a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_thread_muting
    ADD CONSTRAINT "FK_29c11c7deb06615076f8c95b80a" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: clip FK_2b5ec6c574d6802c94c80313fb2; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.clip
    ADD CONSTRAINT "FK_2b5ec6c574d6802c94c80313fb2" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: signin FK_2c308dbdc50d94dc625670055f7; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.signin
    ADD CONSTRAINT "FK_2c308dbdc50d94dc625670055f7" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: messaging_message FK_2c4be03b446884f9e9c502135be; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.messaging_message
    ADD CONSTRAINT "FK_2c4be03b446884f9e9c502135be" FOREIGN KEY ("groupId") REFERENCES public.user_group(id) ON DELETE CASCADE;


--
-- Name: blocking FK_2cd4a2743a99671308f5417759e; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.blocking
    ADD CONSTRAINT "FK_2cd4a2743a99671308f5417759e" FOREIGN KEY ("blockeeId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: page FK_3126dd7c502c9e4d7597ef7ef10; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page
    ADD CONSTRAINT "FK_3126dd7c502c9e4d7597ef7ef10" FOREIGN KEY ("eyeCatchingImageId") REFERENCES public.drive_file(id) ON DELETE CASCADE;


--
-- Name: reversi_matching FK_3b25402709dd9882048c2bbade0; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_matching
    ADD CONSTRAINT "FK_3b25402709dd9882048c2bbade0" FOREIGN KEY ("parentId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: notification FK_3b4e96eec8d36a8bbb9d02aa710; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_3b4e96eec8d36a8bbb9d02aa710" FOREIGN KEY ("notifierId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: notification FK_3c601b70a1066d2c8b517094cb9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_3c601b70a1066d2c8b517094cb9" FOREIGN KEY ("notifieeId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_group FK_3d6b372788ab01be58853003c93; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group
    ADD CONSTRAINT "FK_3d6b372788ab01be58853003c93" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: app FK_3f5b0899ef90527a3462d7c2cb3; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.app
    ADD CONSTRAINT "FK_3f5b0899ef90527a3462d7c2cb3" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: note_reaction FK_45145e4953780f3cd5656f0ea6a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_reaction
    ADD CONSTRAINT "FK_45145e4953780f3cd5656f0ea6a" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: note_favorite FK_47f4b1892f5d6ba8efb3057d81a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_favorite
    ADD CONSTRAINT "FK_47f4b1892f5d6ba8efb3057d81a" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: password_reset_request FK_4bb7fd4a34492ae0e6cc8d30ac8; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.password_reset_request
    ADD CONSTRAINT "FK_4bb7fd4a34492ae0e6cc8d30ac8" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_profile FK_51cb79b5555effaf7d69ba1cff9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_profile
    ADD CONSTRAINT "FK_51cb79b5555effaf7d69ba1cff9" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note FK_52ccc804d7c69037d558bac4c96; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT "FK_52ccc804d7c69037d558bac4c96" FOREIGN KEY ("renoteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: messaging_message FK_535def119223ac05ad3fa9ef64b; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.messaging_message
    ADD CONSTRAINT "FK_535def119223ac05ad3fa9ef64b" FOREIGN KEY ("fileId") REFERENCES public.drive_file(id) ON DELETE CASCADE;


--
-- Name: messaging_message FK_5377c307783fce2b6d352e1203b; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.messaging_message
    ADD CONSTRAINT "FK_5377c307783fce2b6d352e1203b" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note_unread FK_56b0166d34ddae49d8ef7610bb9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_unread
    ADD CONSTRAINT "FK_56b0166d34ddae49d8ef7610bb9" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user FK_58f5c71eaab331645112cf8cfa5; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "FK_58f5c71eaab331645112cf8cfa5" FOREIGN KEY ("avatarId") REFERENCES public.drive_file(id) ON DELETE SET NULL;


--
-- Name: note FK_5b87d9d19127bd5d92026017a7b; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT "FK_5b87d9d19127bd5d92026017a7b" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_group_invitation FK_5cc8c468090e129857e9fecce5a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invitation
    ADD CONSTRAINT "FK_5cc8c468090e129857e9fecce5a" FOREIGN KEY ("userGroupId") REFERENCES public.user_group(id) ON DELETE CASCADE;


--
-- Name: announcement_read FK_603a7b1e7aa0533c6c88e9bfafe; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.announcement_read
    ADD CONSTRAINT "FK_603a7b1e7aa0533c6c88e9bfafe" FOREIGN KEY ("announcementId") REFERENCES public.announcement(id) ON DELETE CASCADE;


--
-- Name: user_list_joining FK_605472305f26818cc93d1baaa74; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_list_joining
    ADD CONSTRAINT "FK_605472305f26818cc93d1baaa74" FOREIGN KEY ("userListId") REFERENCES public.user_list(id) ON DELETE CASCADE;


--
-- Name: antenna FK_6446c571a0e8d0f05f01c789096; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna
    ADD CONSTRAINT "FK_6446c571a0e8d0f05f01c789096" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: following FK_6516c5a6f3c015b4eed39978be5; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.following
    ADD CONSTRAINT "FK_6516c5a6f3c015b4eed39978be5" FOREIGN KEY ("followerId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: reversi_game FK_6649a4e8c5d5cf32fb03b5da9f6; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_game
    ADD CONSTRAINT "FK_6649a4e8c5d5cf32fb03b5da9f6" FOREIGN KEY ("user2Id") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: poll_vote FK_66d2bd2ee31d14bcc23069a89f8; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.poll_vote
    ADD CONSTRAINT "FK_66d2bd2ee31d14bcc23069a89f8" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_group_joining FK_67dc758bc0566985d1b3d399865; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_joining
    ADD CONSTRAINT "FK_67dc758bc0566985d1b3d399865" FOREIGN KEY ("userGroupId") REFERENCES public.user_group(id) ON DELETE CASCADE;


--
-- Name: user_note_pining FK_68881008f7c3588ad7ecae471cf; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_note_pining
    ADD CONSTRAINT "FK_68881008f7c3588ad7ecae471cf" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: channel_following FK_6d8084ec9496e7334a4602707e1; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_following
    ADD CONSTRAINT "FK_6d8084ec9496e7334a4602707e1" FOREIGN KEY ("followerId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_profile FK_6dc44f1ceb65b1e72bacef2ca27; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_profile
    ADD CONSTRAINT "FK_6dc44f1ceb65b1e72bacef2ca27" FOREIGN KEY ("pinnedPageId") REFERENCES public.page(id) ON DELETE SET NULL;


--
-- Name: antenna FK_709d7d32053d0dd7620f678eeb9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna
    ADD CONSTRAINT "FK_709d7d32053d0dd7620f678eeb9" FOREIGN KEY ("userListId") REFERENCES public.user_list(id) ON DELETE CASCADE;


--
-- Name: muted_note FK_70ab9786313d78e4201d81cdb89; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muted_note
    ADD CONSTRAINT "FK_70ab9786313d78e4201d81cdb89" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: notification FK_769cb6b73a1efe22ddf733ac453; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_769cb6b73a1efe22ddf733ac453" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: channel_note_pining FK_8125f950afd3093acb10d2db8a8; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel_note_pining
    ADD CONSTRAINT "FK_8125f950afd3093acb10d2db8a8" FOREIGN KEY ("channelId") REFERENCES public.channel(id) ON DELETE CASCADE;


--
-- Name: channel FK_823bae55bd81b3be6e05cff4383; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT "FK_823bae55bd81b3be6e05cff4383" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: announcement_read FK_8288151386172b8109f7239ab28; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.announcement_read
    ADD CONSTRAINT "FK_8288151386172b8109f7239ab28" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: drive_file FK_860fa6f6c7df5bb887249fba22e; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_file
    ADD CONSTRAINT "FK_860fa6f6c7df5bb887249fba22e" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: gallery_like FK_8fd5215095473061855ceb948cf; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.gallery_like
    ADD CONSTRAINT "FK_8fd5215095473061855ceb948cf" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: notification FK_8fe87814e978053a53b1beb7e98; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_8fe87814e978053a53b1beb7e98" FOREIGN KEY ("userGroupInvitationId") REFERENCES public.user_group_invitation(id) ON DELETE CASCADE;


--
-- Name: muting FK_93060675b4a79a577f31d260c67; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muting
    ADD CONSTRAINT "FK_93060675b4a79a577f31d260c67" FOREIGN KEY ("muterId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: promo_read FK_9657d55550c3d37bfafaf7d4b05; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.promo_read
    ADD CONSTRAINT "FK_9657d55550c3d37bfafaf7d4b05" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: sw_subscription FK_97754ca6f2baff9b4abb7f853dd; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.sw_subscription
    ADD CONSTRAINT "FK_97754ca6f2baff9b4abb7f853dd" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: gallery_post FK_985b836dddd8615e432d7043ddb; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.gallery_post
    ADD CONSTRAINT "FK_985b836dddd8615e432d7043ddb" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: access_token FK_9949557d0e1b2c19e5344c171e9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.access_token
    ADD CONSTRAINT "FK_9949557d0e1b2c19e5344c171e9" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: channel FK_999da2bcc7efadbfe0e92d3bc19; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT "FK_999da2bcc7efadbfe0e92d3bc19" FOREIGN KEY ("bannerId") REFERENCES public.drive_file(id) ON DELETE SET NULL;


--
-- Name: clip_note FK_a012eaf5c87c65da1deb5fdbfa3; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.clip_note
    ADD CONSTRAINT "FK_a012eaf5c87c65da1deb5fdbfa3" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: moderation_log FK_a08ad074601d204e0f69da9a954; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.moderation_log
    ADD CONSTRAINT "FK_a08ad074601d204e0f69da9a954" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: access_token FK_a3ff16c90cc87a82a0b5959e560; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.access_token
    ADD CONSTRAINT "FK_a3ff16c90cc87a82a0b5959e560" FOREIGN KEY ("appId") REFERENCES public.app(id) ON DELETE CASCADE;


--
-- Name: promo_read FK_a46a1a603ecee695d7db26da5f4; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.promo_read
    ADD CONSTRAINT "FK_a46a1a603ecee695d7db26da5f4" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: follow_request FK_a7fd92dd6dc519e6fb435dd108f; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.follow_request
    ADD CONSTRAINT "FK_a7fd92dd6dc519e6fb435dd108f" FOREIGN KEY ("followerId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: meta FK_ab1bc0c1e209daa77b8e8d212ad; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.meta
    ADD CONSTRAINT "FK_ab1bc0c1e209daa77b8e8d212ad" FOREIGN KEY ("proxyAccountId") REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: page FK_ae1d917992dd0c9d9bbdad06c4a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page
    ADD CONSTRAINT "FK_ae1d917992dd0c9d9bbdad06c4a" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: poll_vote FK_aecfbd5ef60374918e63ee95fa7; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.poll_vote
    ADD CONSTRAINT "FK_aecfbd5ef60374918e63ee95fa7" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: user FK_afc64b53f8db3707ceb34eb28e2; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "FK_afc64b53f8db3707ceb34eb28e2" FOREIGN KEY ("bannerId") REFERENCES public.drive_file(id) ON DELETE SET NULL;


--
-- Name: note_watching FK_b0134ec406e8d09a540f8182888; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_watching
    ADD CONSTRAINT "FK_b0134ec406e8d09a540f8182888" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: gallery_like FK_b1cb568bfe569e47b7051699fc8; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.gallery_like
    ADD CONSTRAINT "FK_b1cb568bfe569e47b7051699fc8" FOREIGN KEY ("postId") REFERENCES public.gallery_post(id) ON DELETE CASCADE;


--
-- Name: user_list FK_b7fcefbdd1c18dce86687531f99; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_list
    ADD CONSTRAINT "FK_b7fcefbdd1c18dce86687531f99" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: drive_file FK_bb90d1956dafc4068c28aa7560a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_file
    ADD CONSTRAINT "FK_bb90d1956dafc4068c28aa7560a" FOREIGN KEY ("folderId") REFERENCES public.drive_folder(id) ON DELETE SET NULL;


--
-- Name: antenna_note FK_bd0397be22147e17210940e125b; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna_note
    ADD CONSTRAINT "FK_bd0397be22147e17210940e125b" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: notification FK_bd7fab507621e635b32cd31892c; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_bd7fab507621e635b32cd31892c" FOREIGN KEY ("followRequestId") REFERENCES public.follow_request(id) ON DELETE CASCADE;


--
-- Name: user_group_invitation FK_bfbc6305547539369fe73eb144a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invitation
    ADD CONSTRAINT "FK_bfbc6305547539369fe73eb144a" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_note_pining FK_bfbc6f79ba4007b4ce5097f08d6; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_note_pining
    ADD CONSTRAINT "FK_bfbc6f79ba4007b4ce5097f08d6" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: auth_session FK_c072b729d71697f959bde66ade0; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.auth_session
    ADD CONSTRAINT "FK_c072b729d71697f959bde66ade0" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: messaging_message FK_cac14a4e3944454a5ce7daa5142; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.messaging_message
    ADD CONSTRAINT "FK_cac14a4e3944454a5ce7daa5142" FOREIGN KEY ("recipientId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: antenna FK_ccbf5a8c0be4511133dcc50ddeb; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.antenna
    ADD CONSTRAINT "FK_ccbf5a8c0be4511133dcc50ddeb" FOREIGN KEY ("userGroupJoiningId") REFERENCES public.user_group_joining(id) ON DELETE CASCADE;


--
-- Name: page_like FK_cf8782626dced3176038176a847; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.page_like
    ADD CONSTRAINT "FK_cf8782626dced3176038176a847" FOREIGN KEY ("pageId") REFERENCES public.page(id) ON DELETE CASCADE;


--
-- Name: user_list_joining FK_d844bfc6f3f523a05189076efaa; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_list_joining
    ADD CONSTRAINT "FK_d844bfc6f3f523a05189076efaa" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: muted_note FK_d8e07aa18c2d64e86201601aec1; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muted_note
    ADD CONSTRAINT "FK_d8e07aa18c2d64e86201601aec1" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: poll FK_da851e06d0dfe2ef397d8b1bf1b; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.poll
    ADD CONSTRAINT "FK_da851e06d0dfe2ef397d8b1bf1b" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: auth_session FK_dbe037d4bddd17b03a1dc778dee; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.auth_session
    ADD CONSTRAINT "FK_dbe037d4bddd17b03a1dc778dee" FOREIGN KEY ("appId") REFERENCES public.app(id) ON DELETE CASCADE;


--
-- Name: user_group_invite FK_e10924607d058004304611a436a; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_invite
    ADD CONSTRAINT "FK_e10924607d058004304611a436a" FOREIGN KEY ("userGroupId") REFERENCES public.user_group(id) ON DELETE CASCADE;


--
-- Name: notification FK_e22bf6bda77b6adc1fd9e75c8c9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT "FK_e22bf6bda77b6adc1fd9e75c8c9" FOREIGN KEY ("appAccessTokenId") REFERENCES public.access_token(id) ON DELETE CASCADE;


--
-- Name: reversi_matching FK_e247b23a3c9b45f89ec1299d066; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_matching
    ADD CONSTRAINT "FK_e247b23a3c9b45f89ec1299d066" FOREIGN KEY ("childId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: promo_note FK_e263909ca4fe5d57f8d4230dd5c; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.promo_note
    ADD CONSTRAINT "FK_e263909ca4fe5d57f8d4230dd5c" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: note_unread FK_e637cba4dc4410218c4251260e4; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note_unread
    ADD CONSTRAINT "FK_e637cba4dc4410218c4251260e4" FOREIGN KEY ("noteId") REFERENCES public.note(id) ON DELETE CASCADE;


--
-- Name: clip_note FK_ebe99317bbbe9968a0c6f579adf; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.clip_note
    ADD CONSTRAINT "FK_ebe99317bbbe9968a0c6f579adf" FOREIGN KEY ("clipId") REFERENCES public.clip(id) ON DELETE CASCADE;


--
-- Name: muting FK_ec96b4fed9dae517e0dbbe0675c; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.muting
    ADD CONSTRAINT "FK_ec96b4fed9dae517e0dbbe0675c" FOREIGN KEY ("muteeId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: attestation_challenge FK_f1a461a618fa1755692d0e0d592; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.attestation_challenge
    ADD CONSTRAINT "FK_f1a461a618fa1755692d0e0d592" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: note FK_f22169eb10657bded6d875ac8f9; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT "FK_f22169eb10657bded6d875ac8f9" FOREIGN KEY ("channelId") REFERENCES public.channel(id) ON DELETE CASCADE;


--
-- Name: user_group_joining FK_f3a1b4bd0c7cabba958a0c0b231; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_group_joining
    ADD CONSTRAINT "FK_f3a1b4bd0c7cabba958a0c0b231" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_keypair FK_f4853eb41ab722fe05f81cedeb6; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_keypair
    ADD CONSTRAINT "FK_f4853eb41ab722fe05f81cedeb6" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: drive_folder FK_f4fc06e49c0171c85f1c48060d2; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.drive_folder
    ADD CONSTRAINT "FK_f4fc06e49c0171c85f1c48060d2" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: reversi_game FK_f7467510c60a45ce5aca6292743; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.reversi_game
    ADD CONSTRAINT "FK_f7467510c60a45ce5aca6292743" FOREIGN KEY ("user1Id") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: registry_item FK_fb9d21ba0abb83223263df6bcb3; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.registry_item
    ADD CONSTRAINT "FK_fb9d21ba0abb83223263df6bcb3" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: user_security_key FK_ff9ca3b5f3ee3d0681367a9b447; Type: FK CONSTRAINT; Schema: public; Owner: misskey
--

ALTER TABLE ONLY public.user_security_key
    ADD CONSTRAINT "FK_ff9ca3b5f3ee3d0681367a9b447" FOREIGN KEY ("userId") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

