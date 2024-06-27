PORTNAME=	json-diff
DISTVERSIONPREFIX=	v
DISTVERSION=	20200926
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Compare two JSON files
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	UNLICENSE
LICENSE_FILE=	${WRKSRC}/UNLICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	ksceriath
GH_TAGNAME=	a7ed330

#CARGO_CARGOTOML=	${WRKSRC}/${PORTNAME}/Cargo.toml
#CARGO_INSTALL_PATH=	${PORTNAME}

PLIST_FILES=	bin/json_diff

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${PLIST_FILES:Mbin/*}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f Makefile.deps && make makesum && make cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
