PORTNAME=	markdown2html
DISTVERSIONPREFIX=	v
DISTVERSION=	1.1.11
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Converting Markdown file to single HTML file with built-in CSS and JS
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	magiclen
GH_PROJECT=	${PORTNAME}-converter

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} \
		--output-file=${STAGEDIR}${PREFIX}/bin/${PORTNAME} \
		${STAGEDIR}${PREFIX}/bin/${PORTNAME}-converter

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f Makefile.deps && make makesum cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
