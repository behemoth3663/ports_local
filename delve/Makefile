PORTNAME=	delve
DISTVERSIONPREFIX=	v
DISTVERSION=	1.25.2
PORTREVISION=	2
CATEGORIES=	local devel

MAINTAINER=	dmitry.wagin@ya.ru
COMMENT=	Debugger for the Go programming language
WWW=		https://${GO_PKGNAME}

LICENSE=	MIT

ONLY_FOR_ARCHS=	amd64

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	go-delve

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${PORTNAME}
GO_TARGET=	./cmd/dlv

GO_BUILDFLAGS+=	-ldflags="-s -w"

CONFLICTS_INSTALL=	dlv # bin/dlv

PLIST_FILES=	bin/dlv

.include <bsd.port.mk>
