PORTNAME=	keepassxc-proxy-rust
PORTVERSION=	0.2.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Proxy between Native Messaging browser extension
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	varjolintu

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#make makesum
#make makesum cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.-include "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${STRIP_CMD} \
		--output-file=${STAGEDIR}${PREFIX}/bin/${PORTNAME} \
		${STAGEDIR}${PREFIX}/bin/${PORTNAME:S|-rust$||}

.include <bsd.port.mk>
