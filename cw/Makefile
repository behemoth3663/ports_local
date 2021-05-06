#

PORTNAME=	cw
PORTVERSION=	3.3.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Best way to tail AWS CloudWatch Logs from your terminal

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BROKEN=		PORTNAME conflicts with sysutils/cw

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	lucagrulla

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
