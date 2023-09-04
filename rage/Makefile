PORTNAME=	rage
DISTVERSIONPREFIX=	v
DISTVERSION=	0.7.1
#DISTVERSION=	0.9.0
#error: found a virtual manifest at `/usr/obj/usr/ports/local/rage/work/rage-0.9.0/Cargo.toml` instead of a package manifest
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Simple, secure, modern encryption tool

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE-MIT

USES=		cargo ssl
USE_GITHUB=	yes
GH_ACCOUNT=	str4d

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f Makefile.deps && make makesum
#make makesum cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.-include "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
