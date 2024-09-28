PORTNAME=	wstunnel
DISTVERSIONPREFIX=	v
DISTVERSION=	10.1.2
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Tunnel traffic over websocket protocol
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo # ssl
USE_GITHUB=	yes
GH_ACCOUNT=	erebe

#CARGO_CARGOTOML=	${WRKSRC}/${PORTNAME}/Cargo.toml
#CARGO_INSTALL_PATH=	${PORTNAME}

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
