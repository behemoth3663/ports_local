PORTNAME=	markdown2html
DISTVERSIONPREFIX=	v
DISTVERSION=	2.1.1
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
