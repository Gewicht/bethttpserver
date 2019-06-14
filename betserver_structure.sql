
CREATE TABLE IF NOT EXISTS `liga` (
  `id` int(10) unsigned NOT NULL,
  `naziv` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `liga`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `liga`
  MODIFY `id` int(10) unsigned NOT NULL AUTO_INCREMENT;

CREATE TABLE IF NOT EXISTS `liga_ponude` (
  `id` int(11) NOT NULL,
  `liga_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `liga_ponude`
  ADD PRIMARY KEY (`id`);

CREATE TABLE IF NOT EXISTS `liga_tipovi` (
  `id` int(11) NOT NULL,
  `liga_id` int(11) NOT NULL,
  `naziv` varchar(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `liga_tipovi`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `liga_tipovi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

CREATE TABLE IF NOT EXISTS `ponude` (
  `id` int(11) NOT NULL,
  `naziv` varchar(50) NOT NULL,
  `broj` varchar(15) NOT NULL,
  `vrijeme` datetime NOT NULL,
  `tvkanal` varchar(15) NOT NULL,
  `ima_statistiku` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `ponude`
  ADD PRIMARY KEY (`id`);

CREATE TABLE IF NOT EXISTS `ponude_tecaj` (
  `id` int(11) NOT NULL,
  `ponuda_id` int(11) NOT NULL,
  `tecaj` float NOT NULL,
  `naziv` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `ponude_tecaj`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `ponude_tecaj`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

CREATE TABLE IF NOT EXISTS `igrac` (
  `id` int(11) NOT NULL,
  `korisnickoime` varchar(50) NOT NULL,
  `saldo` float(10,2) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
ALTER TABLE `igrac`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `igrac`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=6;


CREATE TABLE IF NOT EXISTS `listici` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `ulog` double NOT NULL,
  `koeficijent` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
ALTER TABLE `listici`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `listici`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

CREATE TABLE IF NOT EXISTS `listici_parovi` (
  `id` int(11) NOT NULL,
  `ponuda_id` int(11) NOT NULL,
  `listici_id` int(11) NOT NULL,
  `tecaj` double NOT NULL,
  `naziv` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
ALTER TABLE `listici_parovi`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `listici_parovi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;