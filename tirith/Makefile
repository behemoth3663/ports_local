PORTNAME=	tirith
PORTVERSION=	0.5.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@FreeBSD.org
COMMENT=	Terminal security for developers and AI agents
WWW=		https://tirith.sh https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	AGPLv3

USES=		cargo

USE_GITHUB=	yes
GH_ACCOUNT=	sheeki03

PLIST_FILES=	bin/${PORTNAME}

#error: found a virtual manifest at `/usr/obj/usr/ports/local/tirith/work/tirith-0.4.2/Cargo.toml` instead of a package manifest
CARGO_INSTALL_PATH=	crates/${PORTNAME}

FETCH_ENV=	HTTP_USER_AGENT='Mozilla/5.0 (X11; Linux x86_64; rv:${FIREFOX_VERSION_CMD:sh}) Gecko/20100101 Firefox/${FIREFOX_VERSION_CMD:sh}'

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} FETCH_ENV="${FETCH_ENV}" makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.pre.mk>

FIREFOX_VERSION_CMD=	${PKG_QUERY} --evaluate %n=firefox %v | ${SED} -r 's/^([0-9]+\.[0-9]+).*/\1/'

.include <bsd.port.post.mk>
