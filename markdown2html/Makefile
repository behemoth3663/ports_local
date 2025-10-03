PORTNAME=	markdown2html
DISTVERSIONPREFIX=	v
DISTVERSION=	1.1.12
PORTREVISION=	1
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
#rm -f Makefile.deps && make makesum cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
