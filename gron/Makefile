# Created by: Ryan Steinmetz <zi@FreeBSD.org>

PORTNAME=	gron
PORTVERSION=	0.7.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	zi@FreeBSD.org
COMMENT=	Transforms JSON into discrete assignments that are easier to grep

LICENSE=	MIT

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	tomnomnom

GO_BUILDFLAGS+=	-ldflags="-s -w -X main.gronVersion=${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
