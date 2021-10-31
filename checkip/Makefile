PORTNAME=	checkip
PORTVERSION=	0.7.0
DISTVERSIONPREFIX=	v
CATEGORIES=	net

MAINTAINER=	none@FreeBSD.org
COMMENT=	Check an IP address using various public services 

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jreisinger

GO_TARGET=	./cmd

PLIST_FILES=	bin/${PORTNAME}

post-stage: .PHONY
	mv ${STAGEDIR}${PREFIX}/bin/cmd ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
