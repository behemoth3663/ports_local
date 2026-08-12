PORTNAME=	syswatch
DISTVERSIONPREFIX=	v
DISTVERSION=	0.8.0
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@FreeBSD.org
COMMENT=	Single-host system diagnostics TUI
WWW=		https://netwatchlabs.com/ https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		cargo

USE_GITHUB=	yes
GH_ACCOUNT=	matthart1983

PLIST_FILES=	bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
