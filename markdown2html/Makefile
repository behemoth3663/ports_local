PORTNAME=	markdown2html
DISTVERSIONPREFIX=	v
DISTVERSION=	2.0.0
#PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Converting Markdown file to single HTML file with built-in CSS and JS
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	magiclen
GH_PROJECT=	${PORTNAME}-converter

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-converter ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
