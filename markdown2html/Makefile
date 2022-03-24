#

PORTNAME=	markdown2html
DISTVERSIONPREFIX=	v
DISTVERSION=	1.1.8
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Converting a Markdown file to a single HTML file with built-in CSS and JS

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	magiclen
GH_PROJECT=	${PORTNAME}-converter

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#make makesum cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.-include "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} \
		--output-file=${STAGEDIR}${PREFIX}/bin/${PORTNAME} \
		${STAGEDIR}${PREFIX}/bin/${PORTNAME}-converter

.include <bsd.port.mk>
