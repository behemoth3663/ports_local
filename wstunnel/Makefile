PORTNAME=	wstunnel
DISTVERSIONPREFIX=	v
DISTVERSION=	8.4.2
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Tunnel traffic over websocket protocol
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	NONE
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
#rm -f Makefile.deps && make makesum
#make cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
