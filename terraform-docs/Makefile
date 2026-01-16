PORTNAME=	terraform-docs
PORTVERSION=	0.21.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules
WWW=		https://${GO_MODULE}

LICENSE=	MIT

USES=		go:modules

PLIST_FILES=	bin/${PORTNAME}

GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"

.include <bsd.port.mk>
