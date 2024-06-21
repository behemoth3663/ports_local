PORTNAME=	rage
DISTVERSIONPREFIX=	v
DISTVERSION=	0.10.0
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
#rm -f Makefile.deps && make makesum
#make cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
