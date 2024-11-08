PORTNAME=	rage
DISTVERSIONPREFIX=	v
DISTVERSION=	0.11.0
PORTREVISION=	1
CATEGORIES=	local security
PKGNAMESUFFIX=	-encryption

MAINTAINER=	ehaupt@FreeBSD.org
COMMENT=	Simple, modern, and secure file encryption tool, using the age format
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20 MIT
LICENSE_COMB=	dual
LICENSE_FILE_APACHE20=	${WRKSRC}/LICENSE-APACHE
LICENSE_FILE_MIT=	${WRKSRC}/LICENSE-MIT

LIB_DEPENDS=	libfuse.so:filesystems/fusefs-libs \
		libzstd.so:archivers/zstd

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	str4d

CARGO_FEATURES=	mount

CONFLICTS=	rage

PLIST_FILES=	bin/${PORTNAME} \
		bin/${PORTNAME}-keygen \
		bin/${PORTNAME}-mount

#error: found a virtual manifest at `/usr/obj/usr/ports/local/rage-encryption/work/rage-0.11.0/Cargo.toml` instead of a package manifest
CARGO_CARGOTOML=	${WRKSRC}/${PORTNAME}/Cargo.toml
CARGO_INSTALL_PATH=	${PORTNAME}

post-install: .PHONY
.for _f in ${PLIST_FILES}
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_f}
.endfor

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	${MAKE} -C ${.CURDIR} makesum
	set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.crates
	${MAKE} -C ${.CURDIR} makesum

.include <bsd.port.mk>
