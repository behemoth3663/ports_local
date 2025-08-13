PORTNAME=	terraform-docs
PORTVERSION=	0.20.0
DISTVERSIONPREFIX=	v
PORTREVISION=	4
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules
WWW=		https://${GO_MODULE}

LICENSE=	MIT

USES=		go:modules

PLIST_FILES=	bin/${PORTNAME}

GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

.include <bsd.port.mk>
