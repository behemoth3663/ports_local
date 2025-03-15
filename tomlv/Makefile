PORTNAME=	tomlv
PORTVERSION=	1.4.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	TOML parser for Golang with reflection
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/COPYING

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	BurntSushi
GH_PROJECT=	toml

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/tomlv

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

.include <bsd.port.mk>
