PORTNAME=	rage
DISTVERSIONPREFIX=	v
DISTVERSION=	0.11.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Simple, secure, modern encryption tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE-MIT

USES=		cargo ssl
USE_GITHUB=	yes
GH_ACCOUNT=	str4d

#error: found a virtual manifest at `/usr/obj/usr/ports/local/rage/work/rage-0.9.2/Cargo.toml` instead of a package manifest
CARGO_CARGOTOML=	${WRKSRC}/${PORTNAME}/Cargo.toml
CARGO_INSTALL_PATH=	${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	${MAKE} -C ${.CURDIR} makesum
	set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.crates
	${MAKE} -C ${.CURDIR} makesum

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
