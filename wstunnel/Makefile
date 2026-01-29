PORTNAME=	wstunnel
PORTVERSION=	10.5.1
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Traffic tunnel over Websocket or HTTP2 to bypass firewalls/DPI
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		cargo llvm:21,build

USE_GITHUB=	yes
GH_ACCOUNT=	erebe

JQ_GETVER_EXPR=	[.[].name | select(test("^v10\\."; "n")) | sub("^v"; "")][0]

CARGO_CARGOTOML=	${WRKSRC}/wstunnel-cli/Cargo.toml
CARGO_INSTALL_PATH=	wstunnel-cli
CARGO_FEATURES=		--no-default-features ring jemalloc

#Attempt to link statically
#CARGO_BUILD_TARGET=	x86_64-unknown-freebsd
#CARGO_BUILD_ARGS+=	-Z build-std=std,panic_abort -Z build-std-features=panic_immediate_abort
#CARGO_ENV+=		RUSTFLAGS="-C target-feature=+crt-static"

PLIST_FILES=	bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
