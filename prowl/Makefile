PORTNAME=	prowl
DISTVERSIONPREFIX=	v
DISTVERSION=	1.9.0
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@FreeBSD.org
COMMENT=	Tiny terminal radar for GitHub pull requests
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		cargo

USE_GITHUB=	yes
GH_ACCOUNT=	caarlos0

PLIST_FILES=	bin/${PORTNAME}

post-extract: .PHONY .SILENT
	${REINPLACE_CMD} -r 's/^(version[[:blank:]]*=[[:blank:]]*).*/\1"${DISTVERSION}"/' ${WRKSRC}/Cargo.toml

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
