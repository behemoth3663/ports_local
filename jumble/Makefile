#

PORTNAME=	jumble
PORTVERSION=	0.1.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local graphics

MAINTAINER=	none@freebsd.org
COMMENT=	Create diagrams on a grid using HCL syntax

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	lucasepe

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
