#

PORTNAME=	jumble
PORTVERSION=	0.1.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local graphics

MAINTAINER=	none@freebsd.org
COMMENT=	Create diagrams on a grid using HCL syntax

LICENSE=	MIT

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	lucasepe

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
