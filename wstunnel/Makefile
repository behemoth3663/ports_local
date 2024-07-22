PORTNAME=	wstunnel
DISTVERSIONPREFIX=	v
DISTVERSION=	9.7.4
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
deps: .PHONY deps-rm-makefile-deps makesum deps-make-makefile-deps makesum

deps-rm-makefile-deps: .PHONY
	 ${RM} Makefile.deps

deps-make-makefile-deps: .PHONY
	set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' >Makefile.deps

.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
