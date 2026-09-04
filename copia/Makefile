PORTNAME=	copia
DISTVERSIONPREFIX=	v
DISTVERSION=	0.3.0
CATEGORIES=	local net

MAINTAINER=	none@FreeBSD.org
COMMENT=	Pure Rust rsync-style delta synchronization library
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3

USES=		cargo
# pkgconfig ssl llvm:23

USE_GITHUB=	yes
GH_ACCOUNT=	paiml

PLIST_FILES=	bin/${PORTNAME}

CARGO_FEATURES=	cli

post-install: .PHONY
.for _f in ${PLIST_FILES}
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_f}
.endfor

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
