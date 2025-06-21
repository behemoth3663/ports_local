PORTNAME=	pet-cli
DISTVERSIONPREFIX=	v
DISTVERSION=	2025.8
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Python environment tools for Visual Studio Code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/../../LICENSE

USES=		cargo # ssl
USE_GITHUB=	yes
GH_ACCOUNT=	microsoft
GH_PROJECT=	python-environment-tools
GH_TAGNAME=	release/${DISTVERSION}

CONFLICTS=	deskutils/pet

WRKSRC_SUBDIR=	crates/pet

PLIST_FILES=	bin/pet

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/pet

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
