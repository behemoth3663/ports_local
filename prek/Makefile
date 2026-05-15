PORTNAME=	prek
DISTVERSIONPREFIX=	v
DISTVERSION=	0.4.0
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@FreeBSD.org
COMMENT=	Better pre-commit, re-engineered in Rust
WWW=		https://prek.j178.dev

LICENSE=	MIT

USES=		cargo

USE_GITHUB=	yes
GH_ACCOUNT=	j178

PLIST_FILES=	bin/${PORTNAME}

#error: found a virtual manifest at `/usr/obj/usr/ports/local/prek/work/prek-0.3.2/Cargo.toml` instead of a package manifest
CARGO_INSTALL_PATH=	crates/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${MASTERDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${MASTERDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${MASTERDIR}/Makefile.crates )
	( ${MAKE} -C ${MASTERDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
