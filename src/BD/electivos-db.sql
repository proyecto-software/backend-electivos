PGDMP          $                y           administracion_electivos    13.3    13.3 b    8           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            9           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            :           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ;           1262    19029    administracion_electivos    DATABASE     t   CREATE DATABASE administracion_electivos WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Spain.1252';
 (   DROP DATABASE administracion_electivos;
                postgres    false            ?            1255    19030    formulario_admin()    FUNCTION     ?  CREATE FUNCTION public.formulario_admin() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
	indicador integer;
	sem_inc integer;
begin
	--cálculo indicador:
	select semestre_incompleto into sem_inc from alumno a where a.id = new.id_alumno;
	indicador = 9 - sem_inc;
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo,cantidad_electivos,estado) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_1)),new.cantidad_electivos,'false');
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo,cantidad_electivos,estado) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_2)),new.cantidad_electivos,'false');
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo,cantidad_electivos,estado) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_3)),new.cantidad_electivos,'false');
	return new;
end;
$$;
 )   DROP FUNCTION public.formulario_admin();
       public          postgres    false            ?            1255    19031    informe_curricular()    FUNCTION     ?  CREATE FUNCTION public.informe_curricular() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
	
begin
	insert into informe_curricular(rut,nrc,nombre_ramo,nota,oportunidad,semestre) values((select rut from alumno a where a.id = new.id_alumno),(select nrc from ramo r where r.id = new.id_ramo),(select nombre from ramo r where r.id = new.id_ramo),new.nota,new.oportunidad,(select semestre from ramo r where r.id = new.id_ramo));
	return new;
end;
$$;
 +   DROP FUNCTION public.informe_curricular();
       public          postgres    false            ?            1259    19032    administrador    TABLE     z   CREATE TABLE public.administrador (
    id integer NOT NULL,
    correo text,
    rut text NOT NULL,
    password text
);
 !   DROP TABLE public.administrador;
       public         heap    postgres    false            ?            1259    19038    Administrador_id_seq    SEQUENCE     ?   CREATE SEQUENCE public."Administrador_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."Administrador_id_seq";
       public          postgres    false    200            <           0    0    Administrador_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public."Administrador_id_seq" OWNED BY public.administrador.id;
          public          postgres    false    201            ?            1259    19040    alumno    TABLE     ?   CREATE TABLE public.alumno (
    id integer NOT NULL,
    rut text NOT NULL,
    nombre text,
    correo text,
    id_carrera integer NOT NULL,
    semestre_incompleto integer,
    cantidad_ramos integer
);
    DROP TABLE public.alumno;
       public         heap    postgres    false            ?            1259    19046    Alumno_id_seq    SEQUENCE     ?   CREATE SEQUENCE public."Alumno_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public."Alumno_id_seq";
       public          postgres    false    202            =           0    0    Alumno_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public."Alumno_id_seq" OWNED BY public.alumno.id;
          public          postgres    false    203            ?            1259    19048    carrera    TABLE     J   CREATE TABLE public.carrera (
    id integer NOT NULL,
    nombre text
);
    DROP TABLE public.carrera;
       public         heap    postgres    false            ?            1259    19054    Carrera_id_seq    SEQUENCE     ?   CREATE SEQUENCE public."Carrera_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public."Carrera_id_seq";
       public          postgres    false    204            >           0    0    Carrera_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public."Carrera_id_seq" OWNED BY public.carrera.id;
          public          postgres    false    205            ?            1259    19056    profesor    TABLE     \   CREATE TABLE public.profesor (
    id integer NOT NULL,
    nombre text,
    correo text
);
    DROP TABLE public.profesor;
       public         heap    postgres    false            ?            1259    19062    Profesor_id_seq    SEQUENCE     ?   CREATE SEQUENCE public."Profesor_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Profesor_id_seq";
       public          postgres    false    206            ?           0    0    Profesor_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."Profesor_id_seq" OWNED BY public.profesor.id;
          public          postgres    false    207            ?            1259    19064    electivo    TABLE     t   CREATE TABLE public.electivo (
    id integer NOT NULL,
    nombre text,
    "NRC" text,
    id_profesor integer
);
    DROP TABLE public.electivo;
       public         heap    postgres    false            ?            1259    19070    electivo_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.electivo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.electivo_id_seq;
       public          postgres    false    208            @           0    0    electivo_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.electivo_id_seq OWNED BY public.electivo.id;
          public          postgres    false    209            ?            1259    19072    informe_curricular    TABLE     ?   CREATE TABLE public.informe_curricular (
    id integer NOT NULL,
    rut text,
    nrc text,
    nombre_ramo text,
    nota double precision,
    oportunidad integer,
    semestre integer
);
 &   DROP TABLE public.informe_curricular;
       public         heap    postgres    false            ?            1259    19078    informe_curricular_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.informe_curricular_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 0   DROP SEQUENCE public.informe_curricular_id_seq;
       public          postgres    false    210            A           0    0    informe_curricular_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE public.informe_curricular_id_seq OWNED BY public.informe_curricular.id;
          public          postgres    false    211            ?            1259    19080    postulacion    TABLE     ?   CREATE TABLE public.postulacion (
    id integer NOT NULL,
    cantidad integer,
    id_electivo integer,
    aprobado boolean
);
    DROP TABLE public.postulacion;
       public         heap    postgres    false            ?            1259    19083    ramo    TABLE     ?   CREATE TABLE public.ramo (
    id integer NOT NULL,
    nrc text,
    nombre text,
    creditos integer,
    semestre integer
);
    DROP TABLE public.ramo;
       public         heap    postgres    false            ?            1259    19089    ramo_alumno    TABLE     ?   CREATE TABLE public.ramo_alumno (
    id_ramo integer NOT NULL,
    id_alumno integer NOT NULL,
    nota double precision,
    oportunidad integer
);
    DROP TABLE public.ramo_alumno;
       public         heap    postgres    false            ?            1259    19092    ramo_alumno_id_ramo_seq    SEQUENCE     ?   CREATE SEQUENCE public.ramo_alumno_id_ramo_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.ramo_alumno_id_ramo_seq;
       public          postgres    false    214            B           0    0    ramo_alumno_id_ramo_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.ramo_alumno_id_ramo_seq OWNED BY public.ramo_alumno.id_ramo;
          public          postgres    false    215            ?            1259    19094    ramo_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.ramo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.ramo_id_seq;
       public          postgres    false    213            C           0    0    ramo_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.ramo_id_seq OWNED BY public.ramo.id;
          public          postgres    false    216            ?            1259    19096    registro_electivos    TABLE     ?   CREATE TABLE public.registro_electivos (
    id integer NOT NULL,
    nombre text,
    cant_alumnos integer,
    "año" integer,
    semestre integer
);
 &   DROP TABLE public.registro_electivos;
       public         heap    postgres    false            ?            1259    19102    registro_electivos_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.registro_electivos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 0   DROP SEQUENCE public.registro_electivos_id_seq;
       public          postgres    false    217            D           0    0    registro_electivos_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE public.registro_electivos_id_seq OWNED BY public.registro_electivos.id;
          public          postgres    false    218            ?            1259    19104    registro_postulacion    TABLE     ?   CREATE TABLE public.registro_postulacion (
    id integer NOT NULL,
    rut text,
    nombre text,
    carrera text,
    indicador text,
    electivo text,
    cantidad_electivos integer,
    estado boolean
);
 (   DROP TABLE public.registro_postulacion;
       public         heap    postgres    false            ?            1259    19110    registro_postulacion_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.registro_postulacion_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.registro_postulacion_id_seq;
       public          postgres    false    219            E           0    0    registro_postulacion_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.registro_postulacion_id_seq OWNED BY public.registro_postulacion.id;
          public          postgres    false    220            ?            1259    19112 	   solicitud    TABLE     ?   CREATE TABLE public.solicitud (
    id integer NOT NULL,
    id_alumno integer NOT NULL,
    id_postulacion_1 integer NOT NULL,
    id_postulacion_2 integer,
    id_postulacion_3 integer,
    cantidad_electivos integer
);
    DROP TABLE public.solicitud;
       public         heap    postgres    false            ?            1259    19115    solicitud_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.solicitud_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.solicitud_id_seq;
       public          postgres    false    221            F           0    0    solicitud_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.solicitud_id_seq OWNED BY public.solicitud.id;
          public          postgres    false    222            m           2604    19117    administrador id    DEFAULT     v   ALTER TABLE ONLY public.administrador ALTER COLUMN id SET DEFAULT nextval('public."Administrador_id_seq"'::regclass);
 ?   ALTER TABLE public.administrador ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    201    200            n           2604    19118 	   alumno id    DEFAULT     h   ALTER TABLE ONLY public.alumno ALTER COLUMN id SET DEFAULT nextval('public."Alumno_id_seq"'::regclass);
 8   ALTER TABLE public.alumno ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202            o           2604    19119 
   carrera id    DEFAULT     j   ALTER TABLE ONLY public.carrera ALTER COLUMN id SET DEFAULT nextval('public."Carrera_id_seq"'::regclass);
 9   ALTER TABLE public.carrera ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    205    204            q           2604    19120    electivo id    DEFAULT     j   ALTER TABLE ONLY public.electivo ALTER COLUMN id SET DEFAULT nextval('public.electivo_id_seq'::regclass);
 :   ALTER TABLE public.electivo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    209    208            r           2604    19121    informe_curricular id    DEFAULT     ~   ALTER TABLE ONLY public.informe_curricular ALTER COLUMN id SET DEFAULT nextval('public.informe_curricular_id_seq'::regclass);
 D   ALTER TABLE public.informe_curricular ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    211    210            p           2604    19122    profesor id    DEFAULT     l   ALTER TABLE ONLY public.profesor ALTER COLUMN id SET DEFAULT nextval('public."Profesor_id_seq"'::regclass);
 :   ALTER TABLE public.profesor ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    206            s           2604    19123    ramo id    DEFAULT     b   ALTER TABLE ONLY public.ramo ALTER COLUMN id SET DEFAULT nextval('public.ramo_id_seq'::regclass);
 6   ALTER TABLE public.ramo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    213            t           2604    19124    ramo_alumno id_ramo    DEFAULT     z   ALTER TABLE ONLY public.ramo_alumno ALTER COLUMN id_ramo SET DEFAULT nextval('public.ramo_alumno_id_ramo_seq'::regclass);
 B   ALTER TABLE public.ramo_alumno ALTER COLUMN id_ramo DROP DEFAULT;
       public          postgres    false    215    214            u           2604    19125    registro_electivos id    DEFAULT     ~   ALTER TABLE ONLY public.registro_electivos ALTER COLUMN id SET DEFAULT nextval('public.registro_electivos_id_seq'::regclass);
 D   ALTER TABLE public.registro_electivos ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217            v           2604    19126    registro_postulacion id    DEFAULT     ?   ALTER TABLE ONLY public.registro_postulacion ALTER COLUMN id SET DEFAULT nextval('public.registro_postulacion_id_seq'::regclass);
 F   ALTER TABLE public.registro_postulacion ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219            w           2604    19127    solicitud id    DEFAULT     l   ALTER TABLE ONLY public.solicitud ALTER COLUMN id SET DEFAULT nextval('public.solicitud_id_seq'::regclass);
 ;   ALTER TABLE public.solicitud ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    222    221                      0    19032    administrador 
   TABLE DATA           B   COPY public.administrador (id, correo, rut, password) FROM stdin;
    public          postgres    false    200   bx       !          0    19040    alumno 
   TABLE DATA           j   COPY public.alumno (id, rut, nombre, correo, id_carrera, semestre_incompleto, cantidad_ramos) FROM stdin;
    public          postgres    false    202   x       #          0    19048    carrera 
   TABLE DATA           -   COPY public.carrera (id, nombre) FROM stdin;
    public          postgres    false    204   ?y       '          0    19064    electivo 
   TABLE DATA           B   COPY public.electivo (id, nombre, "NRC", id_profesor) FROM stdin;
    public          postgres    false    208   my       )          0    19072    informe_curricular 
   TABLE DATA           d   COPY public.informe_curricular (id, rut, nrc, nombre_ramo, nota, oportunidad, semestre) FROM stdin;
    public          postgres    false    210   ?y       +          0    19080    postulacion 
   TABLE DATA           J   COPY public.postulacion (id, cantidad, id_electivo, aprobado) FROM stdin;
    public          postgres    false    212   z       %          0    19056    profesor 
   TABLE DATA           6   COPY public.profesor (id, nombre, correo) FROM stdin;
    public          postgres    false    206   )z       ,          0    19083    ramo 
   TABLE DATA           C   COPY public.ramo (id, nrc, nombre, creditos, semestre) FROM stdin;
    public          postgres    false    213   }z       -          0    19089    ramo_alumno 
   TABLE DATA           L   COPY public.ramo_alumno (id_ramo, id_alumno, nota, oportunidad) FROM stdin;
    public          postgres    false    214   ?z       0          0    19096    registro_electivos 
   TABLE DATA           X   COPY public.registro_electivos (id, nombre, cant_alumnos, "año", semestre) FROM stdin;
    public          postgres    false    217    {       2          0    19104    registro_postulacion 
   TABLE DATA           y   COPY public.registro_postulacion (id, rut, nombre, carrera, indicador, electivo, cantidad_electivos, estado) FROM stdin;
    public          postgres    false    219   E{       4          0    19112 	   solicitud 
   TABLE DATA           |   COPY public.solicitud (id, id_alumno, id_postulacion_1, id_postulacion_2, id_postulacion_3, cantidad_electivos) FROM stdin;
    public          postgres    false    221   b{       G           0    0    Administrador_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."Administrador_id_seq"', 1, false);
          public          postgres    false    201            H           0    0    Alumno_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public."Alumno_id_seq"', 1, true);
          public          postgres    false    203            I           0    0    Carrera_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Carrera_id_seq"', 1, false);
          public          postgres    false    205            J           0    0    Profesor_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."Profesor_id_seq"', 1, false);
          public          postgres    false    207            K           0    0    electivo_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.electivo_id_seq', 1, false);
          public          postgres    false    209            L           0    0    informe_curricular_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.informe_curricular_id_seq', 2, true);
          public          postgres    false    211            M           0    0    ramo_alumno_id_ramo_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.ramo_alumno_id_ramo_seq', 1, false);
          public          postgres    false    215            N           0    0    ramo_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.ramo_id_seq', 2, true);
          public          postgres    false    216            O           0    0    registro_electivos_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.registro_electivos_id_seq', 4, true);
          public          postgres    false    218            P           0    0    registro_postulacion_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.registro_postulacion_id_seq', 48, true);
          public          postgres    false    220            Q           0    0    solicitud_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.solicitud_id_seq', 1, false);
          public          postgres    false    222            y           2606    19129     administrador Administrador_pkey 
   CONSTRAINT     e   ALTER TABLE ONLY public.administrador
    ADD CONSTRAINT "Administrador_pkey" PRIMARY KEY (id, rut);
 L   ALTER TABLE ONLY public.administrador DROP CONSTRAINT "Administrador_pkey";
       public            postgres    false    200    200            {           2606    19131    alumno Alumno_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_pkey" PRIMARY KEY (id, rut);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_pkey";
       public            postgres    false    202    202                       2606    19133    carrera Carrera_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.carrera
    ADD CONSTRAINT "Carrera_pkey" PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.carrera DROP CONSTRAINT "Carrera_pkey";
       public            postgres    false    204            ?           2606    19135    profesor Profesor_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.profesor
    ADD CONSTRAINT "Profesor_pkey" PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.profesor DROP CONSTRAINT "Profesor_pkey";
       public            postgres    false    206            }           2606    19137    alumno alumno_id_key 
   CONSTRAINT     M   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT alumno_id_key UNIQUE (id);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT alumno_id_key;
       public            postgres    false    202            ?           2606    19139    electivo electivo_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_pkey;
       public            postgres    false    208            ?           2606    19141 *   informe_curricular informe_curricular_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.informe_curricular
    ADD CONSTRAINT informe_curricular_pkey PRIMARY KEY (id);
 T   ALTER TABLE ONLY public.informe_curricular DROP CONSTRAINT informe_curricular_pkey;
       public            postgres    false    210            ?           2606    19143    postulacion postulacion_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_pkey;
       public            postgres    false    212            ?           2606    19145    ramo_alumno ramo_alumno_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_pkey PRIMARY KEY (id_ramo, id_alumno);
 F   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_pkey;
       public            postgres    false    214    214            ?           2606    19147    ramo ramo_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.ramo
    ADD CONSTRAINT ramo_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.ramo DROP CONSTRAINT ramo_pkey;
       public            postgres    false    213            ?           2606    19149 *   registro_electivos registro_electivos_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.registro_electivos
    ADD CONSTRAINT registro_electivos_pkey PRIMARY KEY (id);
 T   ALTER TABLE ONLY public.registro_electivos DROP CONSTRAINT registro_electivos_pkey;
       public            postgres    false    217            ?           2606    19151 .   registro_postulacion registro_postulacion_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.registro_postulacion
    ADD CONSTRAINT registro_postulacion_pkey PRIMARY KEY (id);
 X   ALTER TABLE ONLY public.registro_postulacion DROP CONSTRAINT registro_postulacion_pkey;
       public            postgres    false    219            ?           2606    19153    solicitud solicitud_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_pkey;
       public            postgres    false    221            ?           2620    19154 "   solicitud trigger_formulario_admin    TRIGGER     ?   CREATE TRIGGER trigger_formulario_admin BEFORE INSERT ON public.solicitud FOR EACH ROW EXECUTE FUNCTION public.formulario_admin();
 ;   DROP TRIGGER trigger_formulario_admin ON public.solicitud;
       public          postgres    false    221    223            ?           2620    19155 &   ramo_alumno trigger_informe_curricular    TRIGGER     ?   CREATE TRIGGER trigger_informe_curricular BEFORE INSERT ON public.ramo_alumno FOR EACH ROW EXECUTE FUNCTION public.informe_curricular();
 ?   DROP TRIGGER trigger_informe_curricular ON public.ramo_alumno;
       public          postgres    false    224    214            ?           2606    19156    alumno Alumno_id_carrera_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_id_carrera_fkey" FOREIGN KEY (id_carrera) REFERENCES public.carrera(id);
 I   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_id_carrera_fkey";
       public          postgres    false    202    2943    204            ?           2606    19161 "   electivo electivo_id_profesor_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_id_profesor_fkey FOREIGN KEY (id_profesor) REFERENCES public.profesor(id);
 L   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_id_profesor_fkey;
       public          postgres    false    208    2945    206            ?           2606    19166 (   postulacion postulacion_id_electivo_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_id_electivo_fkey FOREIGN KEY (id_electivo) REFERENCES public.electivo(id);
 R   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_id_electivo_fkey;
       public          postgres    false    212    208    2947            ?           2606    19171 &   ramo_alumno ramo_alumno_id_alumno_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 P   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_alumno_fkey;
       public          postgres    false    202    2941    214            ?           2606    19176 $   ramo_alumno ramo_alumno_id_ramo_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_ramo_fkey FOREIGN KEY (id_ramo) REFERENCES public.ramo(id);
 N   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_ramo_fkey;
       public          postgres    false    214    213    2953            ?           2606    19181 "   solicitud solicitud_id_alumno_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 L   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_alumno_fkey;
       public          postgres    false    221    2941    202            ?           2606    19186 )   solicitud solicitud_id_postulacion_1_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_1_fkey FOREIGN KEY (id_postulacion_1) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_1_fkey;
       public          postgres    false    2951    221    212            ?           2606    19191 )   solicitud solicitud_id_postulacion_2_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_2_fkey FOREIGN KEY (id_postulacion_2) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_2_fkey;
       public          postgres    false    221    212    2951            ?           2606    19196 )   solicitud solicitud_id_postulacion_3_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_3_fkey FOREIGN KEY (id_postulacion_3) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_3_fkey;
       public          postgres    false    221    212    2951                  x?????? ? ?      !   ?   x?e?A
?0EדS?M?Hw
? Rn݄4ʔ1??;y/fQw?_|x??w?+?~lI
)??p?bi{?'?G??ڍcK??o??c??㽈??
&?f?	?j????b校$??c?(?s,CfJY?|	?>?_?????'???P?Iv}=?q$?;??t???^?G?`.h?yɾK?      #      x?3??tv??2????2r<?b???? 9?K      '   (   x?3?t5?44b.#NW# ???9]??LcN#?=... l?c      )   W   x?3?42?343?330?5?L442???KO??L-:?6Q?59???????DN=SNCN3.#-&?N?ũ
)@?X?_?i
Vi????? ??y      +      x?????? ? ?      %   D   x?3??M,J?/Vp?H?,??MQ??yz?9\F?ΉE9@I?ļ?ļ|?
 ?I?$2?JP?b???? c??      ,   G   x?3?L442???KO??L-:?6Q?59???????DNSN3.#?"N???T? J,?/?șr??qqq u?      -      x?3?4?4?3?4?2?L??=... ,9K      0   5   x?3?t5?44?420?\F??F??f0?1X??5?4?2?t5?44?qc???? >8	p      2      x?????? ? ?      4      x?????? ? ?     