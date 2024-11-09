PORTNAME=	terraform-docs
PORTVERSION=	0.19.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules
WWW=		https://${GO_MODULE}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules,1.23

PLIST_FILES=	bin/${PORTNAME}

GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

.include <bsd.port.mk>
