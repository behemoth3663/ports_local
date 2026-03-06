PORTNAME=	goconst
PORTVERSION=	1.9.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Find in Go repeated strings that could be replaced by a constant
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jgautheron

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

.include <bsd.port.mk>
