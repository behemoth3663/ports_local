# $FreeBSD$

PORTNAME=	sqlite3_analyzer
#DISTVERSION=    3.32.3
DISTVERSION=    3.33.0
CATEGORIES=	databases linux
MASTER_SITES=   https://www.sqlite.org/2020/ http://www2.sqlite.org/2020/ http://www3.sqlite.org/2020/
PKGNAMEPREFIX=	linux-
DISTNAME=       sqlite-tools-linux-x86-${PORTVERSION:C/\.([[:digit:]])[[:>:]]/0\1/g:S/.//g}00

MAINTAINER=	none@freebsd.org
COMMENT=        SQLite3 database analyzer

LICENSE=        PD

USES=		linux zip

ONLY_FOR_ARCHS= i386 amd64

NO_BUILD=	yes
STRIP=

PLIST_FILES=	bin/sqldiff bin/sqlite3_analyzer

do-install:
.for _i in ${PLIST_FILES}
	${INSTALL_PROGRAM} ${WRKSRC}/${_i:T} ${STAGEDIR}${PREFIX}/${_i:H:T}
.endfor

.include <bsd.port.mk>
