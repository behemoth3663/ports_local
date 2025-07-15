PORTNAME=	mdbook
DISTVERSIONPREFIX=	v
DISTVERSION=	0.4.52
#PORTREVISION=	2
CATEGORIES=	local textproc

MAINTAINER=	eduardo@FreeBSD.org
COMMENT=	Create book from markdown files
WWW=		https://rust-lang.github.io/mdBook/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo # ssl
USE_GITHUB=	yes
GH_ACCOUNT=	rust-lang
GH_PROJECT=	mdBook

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.deps
	${MAKE} -C ${.CURDIR} makesum
	set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.deps
	${MAKE} -C ${.CURDIR} makesum

.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
