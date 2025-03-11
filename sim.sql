-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 05-09-2020 a las 20:11:44
-- Versión del servidor: 10.4.11-MariaDB
-- Versión de PHP: 7.4.1

set SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
set AUTOCOMMIT = 0;
start transaction;
set time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `sim`
--
create database if not exists sim default character set utf8mb4 collate utf8mb4_general_ci;
use sim;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `log`
--

create table sim.log (
  _t int(11) not NULL,
  _i int(11) not NULL,

  _V0 int(11) not NULL,
  _W0 int(11) not NULL,
  _R0 int(11) not NULL,
  _V1 int(11) not NULL,
  _W1 int(11) not NULL,
  _R1 int(11) not NULL,
  _V2 int(11) not NULL,
  _W2 int(11) not NULL,
  _R2 int(11) not NULL,
  
  _x int(11) not NULL,
  _y int(11) not NULL,
  _z int(11) not NULL,
  _xSup int(11) not NULL,
  _ySup int(11) not NULL,
  _zSup int(11) not NULL,

  _lastV0 int(11) not NULL,
  _lastW0 int(11) not NULL,
  _lastC0 int(11) not NULL,
  _lastV1 int(11) not NULL,
  _lastW1 int(11) not NULL,
  _lastC1 int(11) not NULL,
  _lastV2 int(11) not NULL,
  _lastW2 int(11) not NULL,
  _lastC2 int(11) not NULL,

  _trid int(11) not NULL,
  _obid int(11) not NULL
) engine=InnoDB default charset=utf8mb4;

create table sim.treatments (
  _i int(11) not NULL,

  _0 int(11) not NULL,
  _1 int(11) not NULL,
  _2 int(11) not NULL,
  _3 int(11) not NULL
) engine=InnoDB default charset=utf8mb4;

insert into sim.treatments
(_i, _0, _1, _2, _3) values
(0, 8, 24, 96, 384),
(1, 8, 24, 384, 96),
(2, 8, 96, 24, 384),
(3, 8, 96, 384, 24),
(4, 8, 384, 24, 96),
(5, 8, 384, 96, 24),
(6, 24, 8, 96, 384),
(7, 24, 8, 384, 96),
(8, 24, 96, 8, 384),
(9, 24, 96, 384, 8),
(10, 24, 384, 8, 96),
(11, 24, 384, 96, 8),
(12, 96, 8, 24, 384),
(13, 96, 8, 384, 24),
(14, 96, 24, 8, 384),
(15, 96, 24, 384, 8),
(16, 96, 384, 8, 24),
(17, 96, 384, 24, 8),
(18, 384, 8, 24, 96),
(19, 384, 8, 96, 24),
(20, 384, 24, 8, 96),
(21, 384, 24, 96, 8),
(22, 384, 96, 8, 24),
(23, 384, 96, 24, 8);

-- Views v 1
create view sim.serie_1_extended as select
  _t,
  _i,

  (_V1-_W1)-(_V0-_W0) as _d1_d0,
  (_V2-_W2)-(_V1-_W1) as _d2_d1,

  _V0-_W0 as _d0,
  _V1-_W1 as _d1,
  _V2-_W2 as _d2,

  _V0,
  _W0,
  _R0,
  _V1,
  _W1,
  _R1,
  _V2,
  _W2,
  _R2,
  
  _x,
  _y,
  _z,
  _xSup,
  _ySup,
  _zSup,

  _lastV0,
  _lastW0,
  _lastC0,
  _lastV1,
  _lastW1,
  _lastC1,
  _lastV2,
  _lastW2,
  _lastC2,

  _trid,
  _obid
from sim.serie_1
order by _trid, _obid, _i, _t;

create view sim.serie_1_circularity as select
  _trid as _i,
  _obid as _j,
  count(_i) as _y
from sim.serie_1_extended
where _t = 255 and _d1_d0 > 0 and _d2_d1 > 0
group by _trid, _obid;

-- Views v 2
create view sim.serie_1_circularity as select
  _trid as _i,
  _obid as _j,
  sum(_lastC2)/sum(_lastC0 + _lastC1 + _lastC2) as _y
from sim.serie_1
where _t = 255 -- and _lastC0 <> 0 and _lastC1 <> 0 and _lastC2 <> 0
group by _trid, _obid;

create view sim.serie_1_mean_circularity as select
  _t as _t,
  sum(_lastC2)/sum(_lastC0 + _lastC1 + _lastC2) as _m
from sim.serie_1
group by _t;

-- Views v 3
create view sim.serie_1_circularity as select
  _trid as _i,
  _obid as _j,
  sum(_lastC2)/sum(_lastC0 + _lastC1 + _lastC2) as _y,
  _t as _t
from sim.serie_1
-- where _t = 255 -- and _lastC0 <> 0 and _lastC1 <> 0 and _lastC2 <> 0
group by _trid, _obid, _t;

create view sim.serie_1_treat_0_mean_circularity as select
  c._i as _i,
  t._0 as _0,
  t._1 as _1,
  t._2 as _2,
  t._3 as _3,
  avg(c._y) as _m,
  c._t as _t
from sim.serie_1_circularity c
inner join sim.treatments t on t._i = c._i and t._i = 0
group by c._t;

create view sim.serie_1_entropy as select
  _254._i as _i,
  _254._trid as _trid,
  _254._obid as _obid,

  _254._V0 - _255._V0 as _DV0,
  _254._W0 - _255._W0 as _DW0,
  _254._V1 - _255._V1 as _DV1,
  _254._W1 - _255._W1 as _DW1,
  _254._V2 - _255._V2 as _DV2,
  _254._W2 - _255._W2 as _DW2
from (select * from sim.serie_1 where _t = 255) _255
join (select * from sim.serie_1 where _t = 254) _254 on
  _254._trid = _255._trid and
  _254._obid = _255._obid and
  _254._i = _255._i
where
  _254._V0 <> _255._V0 or
  _254._W0 <> _255._W0 or
  _254._V1 <> _255._V1 or
  _254._W1 <> _255._W1 or
  _254._V2 <> _255._V2 or
  _254._W2 <> _255._W2;

commit;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
