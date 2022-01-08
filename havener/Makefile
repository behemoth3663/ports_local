#

PORTNAME=	havener
PORTVERSION=	2.0.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Swiss army knife for Kubernetes tasks

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
