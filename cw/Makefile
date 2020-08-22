# $FreeBSD$

PORTNAME=	cw
PORTVERSION=	3.3.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Best way to tail AWS CloudWatch Logs from your terminal

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	lucagrulla

#GO_TARGET=	./cmd/${PORTNAME}
BROKEN=		PORTNAME conflicts with /usr/ports/sysutils/cw

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.sinclude "${.CURDIR}/Makefile.deps"
.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
