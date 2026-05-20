PORTNAME=	mdbook
DISTVERSIONPREFIX=	v
DISTVERSION=	0.5.3
#PORTREVISION=	1
CATEGORIES=	local textproc
PKGNAMESUFFIX=	-local

MAINTAINER=	eduardo@FreeBSD.org
COMMENT=	Create book from markdown files
WWW=		https://rust-lang.github.io/mdBook/

LICENSE=	MPL20

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	rust-lang
GH_PROJECT=	mdBook

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${MASTERDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${MASTERDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${MASTERDIR}/Makefile.crates )
	( ${MAKE} -C ${MASTERDIR} makesum )

.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
