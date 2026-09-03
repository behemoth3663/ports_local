PORTNAME=	rsync
DISTVERSIONPREFIX=	v
DISTVERSION=	0.6.4
CATEGORIES=	local devel
PKGNAMEPREFIX=	oc-

MAINTAINER=	none@FreeBSD.org
COMMENT=	Classic rsync re-implementation in pure Rust
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3

USES=		cargo pkgconfig ssl llvm:23

USE_GITHUB=	yes
GH_ACCOUNT=	oferchen

PLIST_FILES=	bin/${PKGNAMEPREFIX}${PORTNAME}

XARCH=		${ARCH:S/amd64/x86_64/:tu}

MAKE_ENV+=	${XARCH}_UNKNOWN_FREEBSD_OPENSSL_INCLUDE_DIR=${OPENSSLINC}
MAKE_ENV+=	${XARCH}_UNKNOWN_FREEBSD_OPENSSL_LIB_DIR=${OPENSSLLIB}
MAKE_ENV+=	CARGO_WORKSPACE_DIR=${WRKSRC}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
